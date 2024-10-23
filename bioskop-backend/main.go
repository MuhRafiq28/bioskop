package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"
    "os"

    "bioskop-backend/handlers"
    "github.com/dgrijalva/jwt-go"
    "github.com/joho/godotenv"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    _ "github.com/lib/pq" // Driver PostgreSQL
)

// Definisikan struktur User
type User struct {
    ID   uint   `json:"id"`
    Role string `json:"role"`
}

// Middleware JWT
func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        token := c.Request().Header.Get("Authorization")
        if token == "" {
            return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Token tidak ada"})
        }

        // Validasi token
        user, err := ValidateToken(token)
        if err != nil {
            return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Token tidak valid"})
        }

        c.Set("user_id", user.ID) // Set user ID ke konteks
        return next(c)
    }
}

// Fungsi untuk memvalidasi token
func ValidateToken(tokenString string) (*User, error) {
    secretKey := []byte(os.Getenv("JmySuperSecretKey12345")) // Ambil secret dari environment

    // Log token yang diterima untuk debugging
    log.Println("Token diterima: ", tokenString)

    // Menghapus awalan 'Bearer ' jika ada
    if len(tokenString) > 6 && tokenString[:7] == "Bearer " {
        tokenString = tokenString[7:]
    }

    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        log.Println("Memvalidasi token dengan algoritma: ", token.Header["alg"]) // Log algoritma
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return secretKey, nil
    })

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        userID := uint(claims["id"].(float64))
        role := claims["role"].(string)

        log.Println("Token valid untuk user ID: ", userID, " dengan role: ", role) // Log jika valid
        return &User{ID: userID, Role: role}, nil
    } else {
        log.Println("Error saat memvalidasi token: ", err) // Log error saat validasi
        return nil, err
    }
}

func main() {
    // Load configuration from .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    // Create connection string using parameters from .env
    connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
    )

    // Connect to PostgreSQL database
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatalf("Error connecting to the database: %v", err)
    }

    // Ensure the connection is successful
    err = db.Ping()
    if err != nil {
        log.Fatalf("Failed to ping database: %v", err)
    }
    defer db.Close()

    // Initialize Echo framework
    e := echo.New()

    // Middleware for logging and CORS
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())
    e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
        AllowOrigins: []string{"*"}, // Ganti '*' dengan domain spesifik jika perlu
        AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
        AllowHeaders: []string{"Content-Type", "Authorization"}, // Menambahkan header yang diizinkan
    }))

    // Serve static files from the uploads directory
    e.Static("/uploads", "uploads")

    // Routes for login and registration
    e.POST("/register", handlers.Register(db))
    e.POST("/login", handlers.Login(db))

    // CRUD routes for users
    e.GET("/users", handlers.GetAllUsers(db), JWTMiddleware) // Get all users
    e.GET("/users/:id", handlers.GetUserByID(db), JWTMiddleware) // Get user by ID
    e.PUT("/users/:id", handlers.UpdateUser(db), JWTMiddleware) // Update user by ID
    e.DELETE("/users/:id", handlers.DeleteUser(db), JWTMiddleware) // Delete user by ID

    // CRUD routes for movies
    e.GET("/movies", handlers.GetAllMovies(db)) // Get all movies
    e.POST("/movies", handlers.AddMovie(db), JWTMiddleware) // Tambahkan JWTMiddleware untuk autentikasi
    e.GET("/movies/:id", handlers.GetMovieByID(db)) // Dapatkan film berdasarkan ID
    e.PUT("/movies/:id", handlers.UpdateMovie(db), JWTMiddleware) // Update movie by ID
    e.DELETE("/movies/:id", handlers.DeleteMovie(db), JWTMiddleware) // Delete movie by ID

    // Start server on port 8080
    e.Logger.Fatal(e.Start(":8080"))
}
