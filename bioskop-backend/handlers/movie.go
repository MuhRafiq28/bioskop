package handlers

import (
    "database/sql"
    "github.com/labstack/echo/v4"
    "io"
    "log"
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
    UserID      int       `json:"user_id"`
    ImageURL    string    `json:"image_url"`
}

// Fungsi AddMovie untuk menambahkan film dengan upload gambar
func AddMovie(db *sql.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        // Mendapatkan informasi pengguna dari konteks
        user, ok := c.Get("user").(*User)
        if !ok || user == nil {
            return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})
        }

        // Memeriksa apakah pengguna adalah admin
        if user.Role != "admin" {
            return c.JSON(http.StatusForbidden, map[string]string{"message": "Hanya admin yang bisa menambahkan film"})
        }

        // Mengikat data film (tanpa gambar)
        movie := new(Movie)
        if err := c.Bind(movie); err != nil {
            log.Printf("Error binding movie data: %v\n", err)
            return c.JSON(http.StatusBadRequest, map[string]string{"message": "Permintaan tidak valid"})
        }

        // Logging untuk melihat data yang diterima
        log.Printf("Data film yang diterima: %+v\n", movie)

        // Validasi data film
        if movie.Title == "" || movie.Genre == "" || movie.Description == "" || movie.ReleaseDate.IsZero() {
            return c.JSON(http.StatusBadRequest, map[string]string{"message": "Title, Genre, Description, dan ReleaseDate diperlukan"})
        }

        // Mengelola upload gambar
        file, err := c.FormFile("image")
        if err != nil {
            return c.JSON(http.StatusBadRequest, map[string]string{"message": "Gambar tidak ditemukan"})
        }

        // Membuka file
        src, err := file.Open()
        if err != nil {
            return err
        }
        defer src.Close()

        // Menentukan jalur tujuan file
        imagePath := filepath.Join("uploads", file.Filename)

        // Memeriksa apakah file sudah ada (opsional)
        if _, err := os.Stat(imagePath); err == nil {
            return c.JSON(http.StatusConflict, map[string]string{"message": "Gambar sudah ada"})
        }

        dst, err := os.Create(imagePath)
        if err != nil {
            return err
        }
        defer dst.Close()

        // Menyalin file ke server
        if _, err = io.Copy(dst, src); err != nil {
            return err
        }

        // Menyimpan URL gambar dan data film ke dalam database
        imageUrl := "/" + imagePath
        _, err = db.Exec("INSERT INTO movies (title, genre, description, release_date, user_id, image_url) VALUES ($1, $2, $3, $4, $5, $6)",
            movie.Title, movie.Genre, movie.Description, movie.ReleaseDate, user.ID, imageUrl)
        if err != nil {
            log.Printf("Error adding movie: %v\n", err)
            return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal menambahkan film"})
        }

        return c.JSON(http.StatusOK, map[string]string{"message": "Film berhasil ditambahkan", "image_url": imageUrl})
    }
}

// Fungsi GetAllMovies untuk mengambil semua film
func GetAllMovies(db *sql.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        rows, err := db.Query("SELECT id, title, genre, description, release_date, image_url FROM movies")
        if err != nil {
            return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error fetching movies"})
        }
        defer rows.Close()

        movies := []Movie{}
        for rows.Next() {
            var movie Movie
            if err := rows.Scan(&movie.ID, &movie.Title, &movie.Genre, &movie.Description, &movie.ReleaseDate, &movie.ImageURL); err != nil {
                return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error scanning movie"})
            }
            movies = append(movies, movie)
        }
        return c.JSON(http.StatusOK, movies)
    }
}

// Fungsi GetMovieByID untuk mengambil film berdasarkan ID
func GetMovieByID(db *sql.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        id := c.Param("id")
        var movie Movie

        err := db.QueryRow("SELECT id, title, genre, description, release_date, image_url FROM movies WHERE id = $1", id).Scan(
            &movie.ID, &movie.Title, &movie.Genre, &movie.Description, &movie.ReleaseDate, &movie.ImageURL,
        )
        if err != nil {
            if err == sql.ErrNoRows {
                return c.JSON(http.StatusNotFound, map[string]string{"message": "Film tidak ditemukan"})
            }
            return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal mengambil data film"})
        }

        return c.JSON(http.StatusOK, movie)
    }
}

// Fungsi UpdateMovie untuk memperbarui film berdasarkan ID
func UpdateMovie(db *sql.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        id := c.Param("id")
        movie := new(Movie)
        if err := c.Bind(movie); err != nil {
            return c.JSON(http.StatusBadRequest, map[string]string{"message": "Permintaan tidak valid"})
        }

        // Logging untuk melihat data yang diterima
        log.Printf("Data film yang diterima untuk diperbarui: %+v\n", movie)

        // Validasi data film
        if movie.Title == "" || movie.Genre == "" || movie.Description == "" || movie.ReleaseDate.IsZero() {
            return c.JSON(http.StatusBadRequest, map[string]string{"message": "Title, Genre, Description, dan ReleaseDate diperlukan"})
        }

        // Melakukan pembaruan film berdasarkan ID yang diterima
        _, err := db.Exec("UPDATE movies SET title=$1, genre=$2, description=$3, release_date=$4, image_url=$5 WHERE id=$6",
            movie.Title, movie.Genre, movie.Description, movie.ReleaseDate, movie.ImageURL, id)
        if err != nil {
            return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error updating movie"})
        }
        return c.JSON(http.StatusOK, map[string]string{"message": "Film berhasil diperbarui"})
    }
}

// Fungsi DeleteMovie untuk menghapus film berdasarkan ID
func DeleteMovie(db *sql.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        id := c.Param("id")
        _, err := db.Exec("DELETE FROM movies WHERE id=$1", id)
        if err != nil {
            return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error deleting movie"})
        }
        return c.JSON(http.StatusOK, map[string]string{"message": "Film berhasil dihapus"})
    }
}
