package handlers

import (
    "database/sql"
    "log"
    "net/http"
    "time"
    "strconv"
    "github.com/dgrijalva/jwt-go"
    "github.com/labstack/echo/v4"
    "golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("JWT_SECRET") 

// User struct
type User struct {
    ID       int    `json:"id"`
    Name     string `json:"name"`
    Password string `json:"password"`
    Role     string `json:"role"`
    Email    string `json:"email"`
}

// GenerateJWT function
func GenerateJWT(user User) (string, error) {
    claims := jwt.MapClaims{
        "id":   user.ID,
        "role": user.Role,
        "exp":  time.Now().Add(time.Hour * 72).Unix(), // Token expired dalam 72 jam
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtSecret)
}

// Register function
func Register(db *sql.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        user := new(User)
        if err := c.Bind(user); err != nil {
            return c.JSON(http.StatusBadRequest, map[string]string{"message": "Permintaan tidak valid"})
        }

        // Cek apakah email sudah terdaftar
        var existingUser User
        err := db.QueryRow("SELECT id FROM users WHERE email=$1", user.Email).Scan(&existingUser.ID)
        if err == nil {
            return c.JSON(http.StatusConflict, map[string]string{"message": "Email sudah digunakan"})
        } else if err != sql.ErrNoRows {
            log.Printf("Error saat mengecek email: %v\n", err)
            return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal memeriksa email"})
        }

        // Hash password sebelum menyimpan
        hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
        if err != nil {
            return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal meng-hash password"})
        }

        // Simpan user ke database
        _, err = db.Exec("INSERT INTO users (name, password, role, email) VALUES ($1, $2, $3, $4)",
            user.Name, string(hashedPassword), user.Role, user.Email)
        if err != nil {
            log.Printf("Error saat registrasi user: %v\n", err)
            return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal registrasi"})
        }

        return c.JSON(http.StatusOK, map[string]string{"message": "Registrasi berhasil"})
    }
}

// Login function
func Login(db *sql.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        user := new(User)
        if err := c.Bind(user); err != nil {
            return c.JSON(http.StatusBadRequest, map[string]string{"message": "Permintaan tidak valid"})
        }

        // Ambil user dari database berdasarkan email
        var storedUser User
        err := db.QueryRow("SELECT id, name, password, role FROM users WHERE email=$1", user.Email).Scan(&storedUser.ID, &storedUser.Name, &storedUser.Password, &storedUser.Role)
        if err != nil {
            log.Printf("Error saat login: %v\n", err)
            return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Email atau password salah"})
        }

        // Verifikasi password
        if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
            return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Email atau password salah"})
        }

        // Generate JWT
        token, err := GenerateJWT(storedUser)
        if err != nil {
            log.Printf("Error saat generate token: %v\n", err)
            return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal generate token"})
        }

        // Return token bersama dengan data user (name, role, dll.)
        return c.JSON(http.StatusOK, map[string]interface{}{
            "token": token,
            "id":    storedUser.ID,
            "name":  storedUser.Name,
            "role":  storedUser.Role, // Kirim role ke frontend
        })
    }
}

// GetAllUsers function
func GetAllUsers(db *sql.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        // Cek token dan ambil informasi pengguna
        user, ok := c.Get("user").(*User)
        if !ok || user == nil {
            return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})
        }

        // Ambil semua pengguna dari database
        rows, err := db.Query("SELECT id, name, role, email FROM users")
        if err != nil {
            return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error fetching users"})
        }
        defer rows.Close()

        users := []User{}
        for rows.Next() {
            var u User
            if err := rows.Scan(&u.ID, &u.Name, &u.Role, &u.Email); err != nil {
                return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error scanning user"})
            }
            users = append(users, u)
        }
        return c.JSON(http.StatusOK, users)
    }
}

// GetUserByID function
func GetUserByID(db *sql.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        id := c.Param("id") // Mendapatkan ID dari parameter URL
        var u User

        // Menjalankan query untuk mencari user berdasarkan ID
        err := db.QueryRow("SELECT id, name, role, email FROM users WHERE id = $1", id).Scan(
            &u.ID, &u.Name, &u.Role, &u.Email,
        )
        if err != nil {
            if err == sql.ErrNoRows {
                return c.JSON(http.StatusNotFound, map[string]string{"message": "User tidak ditemukan"})
            }
            return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal mengambil data user"})
        }

        return c.JSON(http.StatusOK, u) // Mengembalikan data user
    }
}

// UpdateUser function
func UpdateUser(db *sql.DB) echo.HandlerFunc {
  return func(c echo.Context) error {
      // Ambil data pengguna dari token (yang disimpan di context oleh middleware)
      loggedInUser, ok := c.Get("user").(*User)
      if !ok || loggedInUser == nil {
          return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})
      }

      id := c.Param("id") // Mendapatkan ID dari parameter URL
      if id != strconv.Itoa(loggedInUser.ID) && loggedInUser.Role != "admin" {
          // Hanya pengguna itu sendiri atau admin yang bisa mengupdate
          return c.JSON(http.StatusForbidden, map[string]string{"message": "Tidak memiliki izin untuk memperbarui data pengguna ini"})
      }

      updatedUser := new(User)
      if err := c.Bind(updatedUser); err != nil {
          return c.JSON(http.StatusBadRequest, map[string]string{"message": "Permintaan tidak valid"})
      }

      // Jika password di-update, hash password baru
      var hashedPassword string
      if updatedUser.Password != "" {
          hashed, err := bcrypt.GenerateFromPassword([]byte(updatedUser.Password), bcrypt.DefaultCost)
          if err != nil {
              return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal meng-hash password"})
          }
          hashedPassword = string(hashed)
      } else {
          // Jika password tidak diubah, ambil password yang lama
          var storedPassword string
          err := db.QueryRow("SELECT password FROM users WHERE id=$1", id).Scan(&storedPassword)
          if err != nil {
              return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal mengambil password lama"})
          }
          hashedPassword = storedPassword
      }

      // Update pengguna
      _, err := db.Exec("UPDATE users SET name=$1, password=$2, role=$3, email=$4 WHERE id=$5",
          updatedUser.Name, hashedPassword, updatedUser.Role, updatedUser.Email, id)
      if err != nil {
          return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error updating user"})
      }
      return c.JSON(http.StatusOK, map[string]string{"message": "User berhasil diperbarui"})
  }
}


// DeleteUser function
func DeleteUser(db *sql.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        id := c.Param("id")
        _, err := db.Exec("DELETE FROM users WHERE id=$1", id)
        if err != nil {
            return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error deleting user"})
        }
        return c.JSON(http.StatusOK, map[string]string{"message": "User berhasil dihapus"})
    }
}

// Middleware untuk memverifikasi token
func VerifyToken(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        authHeader := c.Request().Header.Get("Authorization")
        if authHeader == "" || len(authHeader) <= 7 || authHeader[:7] != "Bearer " {
            return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Token tidak ditemukan atau tidak valid"})
        }

        tokenStr := authHeader[7:]
        claims := &jwt.MapClaims{}
        token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
            return jwtSecret, nil
        })

        if err != nil || !token.Valid {
            return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Token tidak valid"})
        }

        // Set user information ke context
        userID := int((*claims)["id"].(float64)) // Memastikan konversi ID dari float64 ke int
        userRole := (*claims)["role"].(string)

        c.Set("user", &User{
            ID:   userID,
            Role: userRole,
        })

        return next(c)
    }
}
