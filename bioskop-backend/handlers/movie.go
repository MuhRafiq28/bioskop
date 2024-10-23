package handlers

import (
    "database/sql"
    "errors"
    "fmt"
    "github.com/dgrijalva/jwt-go"
    "github.com/labstack/echo/v4"
    "io"
    "net/http"
    "os"
    "path/filepath"
    "strconv"
    "time"
)

// Struct Claims untuk JWT
type Claims struct {
    UserID uint `json:"user_id"`
    jwt.StandardClaims
}

// ValidateToken memverifikasi token dan mengembalikan user_id
func ValidateToken(token string) (uint, error) {
    claims := &Claims{}
    parsedToken, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
        return []byte(os.Getenv("mySuperSecretKey12345")), nil
    })

    if err != nil || !parsedToken.Valid {
        return 0, errors.New("invalid token")
    }
    return claims.UserID, nil
}

// Struct Movie untuk menyimpan informasi film
type Movie struct {
    ID          int           `json:"id"`
    Title       string        `json:"title"`
    Genre       string        `json:"genre"`
    Description string        `json:"description"`
    ReleaseDate time.Time     `json:"release_date"`
    UserID      sql.NullInt64 `json:"user_id"`
    ImageURL    string        `json:"image_url"`
    Harga       int           `json:"harga"`
    TrailerURL  sql.NullString `json:"trailer_url"`
    JumlahTiket int           `json:"jumlah_tiket"`
}

// AddMovie menambahkan film ke database
func AddMovie(db *sql.DB) echo.HandlerFunc {
  return func(c echo.Context) error {
      // Ambil nilai dari form
      title := c.FormValue("title")
      genre := c.FormValue("genre")
      description := c.FormValue("description")
      releaseDate := c.FormValue("release_date")
      harga := c.FormValue("harga")
      trailerURL := c.FormValue("trailer_url")
      jumlahTiket := c.FormValue("jumlah_tiket")

      // Ambil user_id dari konteks (misal dari JWT)
      userID, ok := c.Get("user_id").(uint)
      if !ok {
          return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})
      }

      // Validasi input
      if title == "" || genre == "" || description == "" || releaseDate == "" || harga == "" || trailerURL == "" || jumlahTiket == "" {
          return c.JSON(http.StatusBadRequest, map[string]string{"message": "Semua field diperlukan"})
      }

      // Log nilai yang sudah terambil
      fmt.Println("Title:", title)
      fmt.Println("Genre:", genre)
      fmt.Println("Description:", description)
      fmt.Println("Release Date:", releaseDate)
      fmt.Println("Harga:", harga)
      fmt.Println("Trailer URL:", trailerURL)
      fmt.Println("Jumlah Tiket:", jumlahTiket)
      fmt.Println("User ID:", userID)

      // Parse tanggal
      parsedReleaseDate, err := time.Parse("2006-01-02", releaseDate)
      if err != nil {
          fmt.Println("Error parsing release date:", err) // Tambahkan log untuk error parsing
          return c.JSON(http.StatusBadRequest, map[string]string{"message": "Format tanggal tidak valid, gunakan YYYY-MM-DD"})
      }

      // Konversi harga dan jumlah tiket ke int
      hargaInt, err := strconv.Atoi(harga)
      if err != nil {
          fmt.Println("Error converting harga to int:", err) // Tambahkan log untuk error konversi harga
          return c.JSON(http.StatusBadRequest, map[string]string{"message": "Harga harus berupa angka"})
      }

      jumlahTiketInt, err := strconv.Atoi(jumlahTiket)
      if err != nil {
          fmt.Println("Error converting jumlah tiket to int:", err) // Tambahkan log untuk error konversi jumlah tiket
          return c.JSON(http.StatusBadRequest, map[string]string{"message": "Jumlah tiket harus berupa angka"})
      }

      // Mengupload gambar
      file, err := c.FormFile("image")
      if err != nil {
          fmt.Println("Error getting file:", err) // Tambahkan log untuk error mendapatkan file gambar
          return c.JSON(http.StatusBadRequest, map[string]string{"message": "Error mendapatkan file gambar"})
      }

      // Tentukan direktori dan path file
      uploadPath := "uploads/"
      if err := os.MkdirAll(uploadPath, os.ModePerm); err != nil {
          fmt.Println("Error creating upload directory:", err) // Tambahkan log untuk error membuat direktori
          return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error membuat direktori"})
      }

      filePath := filepath.Join(uploadPath, file.Filename)
      dst, err := os.Create(filePath)
      if err != nil {
          fmt.Println("Error creating file:", err) // Tambahkan log untuk error saat membuat file
          return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error saat membuat file"})
      }
      defer dst.Close()

      // Salin file ke server
      src, err := file.Open()
      if err != nil {
          fmt.Println("Error opening file:", err) // Tambahkan log untuk error membuka file gambar
          return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error membuka file gambar"})
      }
      defer src.Close()

      if _, err = io.Copy(dst, src); err != nil {
          fmt.Println("Error copying file:", err) // Tambahkan log untuk error menyalin file gambar
          return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error menyalin file gambar"})
      }

      // Menyimpan data film ke database
      _, err = db.Exec(
          "INSERT INTO movies (title, genre, description, release_date, user_id, image_url, harga, trailer_url, jumlah_tiket) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)",
          title, genre, description, parsedReleaseDate, userID, filePath, hargaInt, trailerURL, jumlahTiketInt,
      )
      if err != nil {
          fmt.Println("Error saving movie data:", err) // Tambahkan log untuk error saat menyimpan data
          return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error menyimpan data film"})
      }

      return c.JSON(http.StatusCreated, map[string]string{"message": "Film berhasil ditambahkan"})
  }
}


func GetMovieByID(db *sql.DB) echo.HandlerFunc {
  return func(c echo.Context) error {
      id := c.Param("id")
      var movie Movie
      err := db.QueryRow("SELECT id, title, genre, description, release_date, image_url, harga, trailer_url, jumlah_tiket, user_id FROM movies WHERE id = $1", id).
          Scan(&movie.ID, &movie.Title, &movie.Genre, &movie.Description, &movie.ReleaseDate, &movie.ImageURL, &movie.Harga, &movie.TrailerURL, &movie.JumlahTiket, &movie.UserID)
      if err == sql.ErrNoRows {
          return c.JSON(http.StatusNotFound, map[string]string{"message": "Film tidak ditemukan"})
      } else if err != nil {
          return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal mendapatkan film"})
      }

      // Mengembalikan path relatif untuk image_url
      movie.ImageURL = "/uploads/" + filepath.Base(movie.ImageURL)

      return c.JSON(http.StatusOK, movie)
  }
}

func GetAllMovies(db *sql.DB) echo.HandlerFunc {
  return func(c echo.Context) error {
      rows, err := db.Query("SELECT id, title, genre, description, release_date, image_url, harga, trailer_url, jumlah_tiket, user_id FROM movies")
      if err != nil {
          return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal mendapatkan film"})
      }
      defer rows.Close()

      var movies []Movie
      for rows.Next() {
          var movie Movie
          if err := rows.Scan(&movie.ID, &movie.Title, &movie.Genre, &movie.Description, &movie.ReleaseDate, &movie.ImageURL, &movie.Harga, &movie.TrailerURL, &movie.JumlahTiket, &movie.UserID); err != nil {
              return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal memindahkan data film"})
          }

          // Mengembalikan path relatif untuk image_url
          movie.ImageURL = "/uploads/" + filepath.Base(movie.ImageURL)

          movies = append(movies, movie)
      }

      return c.JSON(http.StatusOK, movies)
  }
}

// Fungsi UpdateMovie untuk memperbarui informasi film
func UpdateMovie(db *sql.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        id := c.Param("id")
        title := c.FormValue("title")
        genre := c.FormValue("genre")
        description := c.FormValue("description")
        releaseDate := c.FormValue("release_date")
        harga := c.FormValue("harga")
        trailerURL := c.FormValue("trailer_url")
        jumlahTiket := c.FormValue("jumlah_tiket")

        // Log data yang diterima
        fmt.Printf("Update film ID %s dengan data: Title: %s, Genre: %s, Description: %s, Release Date: %s, Harga: %s, Trailer URL: %s, Jumlah Tiket: %s\n",
            id, title, genre, description, releaseDate, harga, trailerURL, jumlahTiket)

        // Validasi input
        if title == "" || genre == "" || description == "" || releaseDate == "" || harga == "" || trailerURL == "" || jumlahTiket == "" {
            return c.JSON(http.StatusBadRequest, map[string]string{"message": "Semua field diperlukan"})
        }

        parsedReleaseDate, err := time.Parse("2006-01-02", releaseDate)
        if err != nil {
            return c.JSON(http.StatusBadRequest, map[string]string{"message": "Format tanggal tidak valid, gunakan YYYY-MM-DD"})
        }

        hargaInt, err := strconv.Atoi(harga)
        if err != nil {
            return c.JSON(http.StatusBadRequest, map[string]string{"message": "Harga harus berupa angka"})
        }

        jumlahTiketInt, err := strconv.Atoi(jumlahTiket)
        if err != nil {
            return c.JSON(http.StatusBadRequest, map[string]string{"message": "Jumlah tiket harus berupa angka"})
        }

        _, err = db.Exec("UPDATE movies SET title=$1, genre=$2, description=$3, release_date=$4, harga=$5, trailer_url=$6, jumlah_tiket=$7 WHERE id=$8",
            title, genre, description, parsedReleaseDate, hargaInt, trailerURL, jumlahTiketInt, id)
        if err != nil {
            return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal memperbarui film"})
        }

        return c.JSON(http.StatusOK, map[string]string{"message": "Film berhasil diperbarui"})
    }
}

// Fungsi DeleteMovie untuk menghapus film berdasarkan ID
func DeleteMovie(db *sql.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        id := c.Param("id")
        _, err := db.Exec("DELETE FROM movies WHERE id = $1", id)
        if err != nil {
            return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal menghapus film"})
        }
        return c.JSON(http.StatusOK, map[string]string{"message": "Film berhasil dihapus"})
    }
}
