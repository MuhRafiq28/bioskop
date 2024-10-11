package main

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    "github.com/joho/godotenv"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "bioskop-backend/handlers"

    _ "github.com/lib/pq" // PostgreSQL driver
)

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
    e.Use(middleware.CORS())

    // Routes for login and registration
    e.POST("/register", handlers.Register(db))
    e.POST("/login", handlers.Login(db))

    // CRUD routes for users
    e.GET("/users", handlers.GetAllUsers(db), handlers.VerifyToken)          // Get all users
    e.GET("/users/:id", handlers.GetUserByID(db), handlers.VerifyToken)      // Get user by ID
    e.PUT("/users/:id", handlers.UpdateUser(db), handlers.VerifyToken)       // Update user by ID
    e.DELETE("/users/:id", handlers.DeleteUser(db), handlers.VerifyToken)    // Delete user by ID

    // CRUD routes for movies
    e.GET("/movies", handlers.GetAllMovies(db))                              // Get all movies
    e.GET("/movies/:id", handlers.GetMovieByID(db))                          // Get movie by ID
    e.POST("/movies", handlers.AddMovie(db), handlers.VerifyToken)           // Add new movie
    e.PUT("/movies/:id", handlers.UpdateMovie(db), handlers.VerifyToken)     // Update movie by ID
    e.DELETE("/movies/:id", handlers.DeleteMovie(db), handlers.VerifyToken)  // Delete movie by ID

    // Start server on port 8080
    e.Logger.Fatal(e.Start(":8080"))
}
