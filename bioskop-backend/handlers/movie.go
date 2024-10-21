package handlers

import (
    "database/sql"
    "github.com/labstack/echo/v4"
    "io"
    "fmt"
    "net/http"
    "os"
    "path/filepath"
    "time"
)

// Struct Movie untuk menyimpan informasi film
type Movie struct {
    ID          int       `json:"id"`
    Title       string    `json:"title"`
    Genre       string    `json:"genre"`
    Description string    `json:"description"`
    ReleaseDate time.Time `json:"release_date"`
    UserID      uint      `json:"user_id"`
    ImageURL    string    `json:"image_url"`
}

// Fungsi AddMovie untuk menambahkan film dengan upload gambar
func AddMovie(db *sql.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        // Mengambil nilai dari form
        title := c.FormValue("title")
        genre := c.FormValue("genre")
        description := c.FormValue("description")
        releaseDate := c.FormValue("release_date")

        if title == "" || genre == "" || description == "" || releaseDate == "" {
            return c.JSON(http.StatusBadRequest, map[string]string{"message": "Title, Genre, Description, dan ReleaseDate diperlukan"})
        }

        parsedReleaseDate, err := time.Parse("2006-01-02", releaseDate)
        if err != nil {
            return c.JSON(http.StatusBadRequest, map[string]string{"message": "Format tanggal tidak valid, gunakan YYYY-MM-DD"})
        }

        file, err := c.FormFile("image")
        if err != nil {
            return c.JSON(http.StatusBadRequest, map[string]string{"message": "Gambar tidak ditemukan"})
        }

        src, err := file.Open()
        if err != nil {
            return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal membuka gambar"})
        }
        defer src.Close()

        // Buat folder uploads jika belum ada
        uploadsDir := "uploads"
        if _, err := os.Stat(uploadsDir); os.IsNotExist(err) {
            os.Mkdir(uploadsDir, os.ModePerm)
        }

        // Simpan gambar dengan nama file
        filePath := filepath.Join(uploadsDir, file.Filename)
        dst, err := os.Create(filePath)
        if err != nil {
            return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal menyimpan gambar"})
        }
        defer dst.Close()

        // Salin konten gambar ke file yang baru dibuat
        if _, err := io.Copy(dst, src); err != nil {
            return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal menyalin gambar"})
        }

        // Simpan informasi film ke database
        imageURL := fmt.Sprintf("/uploads/%s", file.Filename)
        _, err = db.Exec("INSERT INTO movies (title, genre, description, release_date, image_url) VALUES ($1, $2, $3, $4, $5)",
            title, genre, description, parsedReleaseDate, imageURL)
        if err != nil {
            return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal menambahkan film ke database"})
        }

        return c.JSON(http.StatusCreated, map[string]string{"message": "Film berhasil ditambahkan"})
    }
}

func GetMovieByID(db *sql.DB) echo.HandlerFunc {
  return func(c echo.Context) error {
      id := c.Param("id")
      var movie Movie
      err := db.QueryRow("SELECT id, title, genre, description, release_date, image_url FROM movies WHERE id = $1", id).Scan(&movie.ID, &movie.Title, &movie.Genre, &movie.Description, &movie.ReleaseDate, &movie.ImageURL)
      if err == sql.ErrNoRows {
          return c.JSON(http.StatusNotFound, map[string]string{"message": "Film tidak ditemukan"})
      } else if err != nil {
          return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal mendapatkan film"})
      }

      return c.JSON(http.StatusOK, movie)
  }
}

// Fungsi GetAllMovies untuk mendapatkan semua film
func GetAllMovies(db *sql.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        rows, err := db.Query("SELECT id, title, genre, description, release_date, image_url FROM movies")
        if err != nil {
            return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal mendapatkan film"})
        }
        defer rows.Close()

        var movies []Movie
        for rows.Next() {
            var movie Movie
            if err := rows.Scan(&movie.ID, &movie.Title, &movie.Genre, &movie.Description, &movie.ReleaseDate, &movie.ImageURL); err != nil {
                return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal memindahkan data film"})
            }
            movies = append(movies, movie)
        }

        return c.JSON(http.StatusOK, movies)
    }
}

// Fungsi UpdateMovie untuk memperbarui informasi film
func UpdateMovie(db *sql.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        // Ambil ID film dari parameter
        id := c.Param("id")

        // Ambil nilai dari form
        title := c.FormValue("title")
        genre := c.FormValue("genre")
        description := c.FormValue("description")
        releaseDate := c.FormValue("release_date")

        // Cek apakah film ada di database
        var existingMovie Movie
        err := db.QueryRow("SELECT id FROM movies WHERE id = $1", id).Scan(&existingMovie.ID)
        if err == sql.ErrNoRows {
            return c.JSON(http.StatusNotFound, map[string]string{"message": "Film tidak ditemukan"})
        } else if err != nil {
            return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal mencari film"})
        }

        // Memperbarui film jika ada
        _, err = db.Exec("UPDATE movies SET title = $1, genre = $2, description = $3, release_date = $4 WHERE id = $5",
            title, genre, description, releaseDate, existingMovie.ID)
        if err != nil {
            return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal memperbarui film"})
        }

        return c.JSON(http.StatusOK, map[string]string{"message": "Film berhasil diperbarui"})
    }
}

// Fungsi DeleteMovie untuk menghapus film
func DeleteMovie(db *sql.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        // Ambil ID film dari parameter
        id := c.Param("id")

        // Cek apakah film ada di database
        var existingMovie Movie
        err := db.QueryRow("SELECT id FROM movies WHERE id = $1", id).Scan(&existingMovie.ID)
        if err == sql.ErrNoRows {
            return c.JSON(http.StatusNotFound, map[string]string{"message": "Film tidak ditemukan"})
        } else if err != nil {
            return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal mencari film"})
        }

        // Hapus film dari database
        _, err = db.Exec("DELETE FROM movies WHERE id = $1", id)
        if err != nil {
            return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal menghapus film"})
        }

        return c.JSON(http.StatusOK, map[string]string{"message": "Film berhasil dihapus"})
    }
}
