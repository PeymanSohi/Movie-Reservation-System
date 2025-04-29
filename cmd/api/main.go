package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/PeymanSohi/Movie-Reservation-System/internal/config"
	"github.com/PeymanSohi/Movie-Reservation-System/internal/database"
	"github.com/PeymanSohi/Movie-Reservation-System/internal/handlers"
	"github.com/PeymanSohi/Movie-Reservation-System/internal/middleware"
	"github.com/PeymanSohi/Movie-Reservation-System/internal/repository"
	"github.com/PeymanSohi/Movie-Reservation-System/internal/service"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize database
	db, err := database.InitDB(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Seed admin user
	adminEmail := os.Getenv("ADMIN_EMAIL")
	adminPassword := os.Getenv("ADMIN_PASSWORD")
	if adminEmail == "" || adminPassword == "" {
		log.Fatal("ADMIN_EMAIL and ADMIN_PASSWORD must be set in environment variables")
	}
	if err := database.SeedAdmin(db, adminEmail, adminPassword); err != nil {
		log.Fatalf("Failed to seed admin user: %v", err)
	}

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)
	movieRepo := repository.NewMovieRepository(db)
	reservationRepo := repository.NewReservationRepository(db)

	// Initialize services
	userService := service.NewUserService(userRepo, cfg)
	movieService := service.NewMovieService(movieRepo)
	reservationService := service.NewReservationService(reservationRepo, movieRepo)

	// Initialize handlers
	userHandler := handlers.NewUserHandler(userService)
	movieHandler := handlers.NewMovieHandler(movieService)
	reservationHandler := handlers.NewReservationHandler(reservationService)

	// Initialize Gin router
	r := gin.Default()

	// Public routes
	public := r.Group("/api")
	{
		public.POST("/register", userHandler.Register)
		public.POST("/login", userHandler.Login)
		public.GET("/movies", movieHandler.ListMovies)
		public.GET("/movies/date", movieHandler.GetMoviesByDate)
		public.GET("/movies/:id", movieHandler.GetMovie)
		public.GET("/showtimes/:showtime_id/seats", movieHandler.GetAvailableSeats)
	}

	// Protected routes
	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware(cfg))
	{
		// User routes
		protected.GET("/profile", userHandler.GetProfile)
		protected.PUT("/profile", userHandler.UpdateProfile)

		// Reservation routes
		protected.POST("/reservations", reservationHandler.CreateReservation)
		protected.GET("/reservations", reservationHandler.GetUserReservations)
		protected.GET("/reservations/:id", reservationHandler.GetReservation)
		protected.DELETE("/reservations/:id", reservationHandler.CancelReservation)
	}

	// Admin routes
	admin := r.Group("/api/admin")
	admin.Use(middleware.AuthMiddleware(cfg), middleware.AdminMiddleware())
	{
		// User management
		admin.GET("/users", userHandler.ListUsers)
		admin.POST("/users/:id/promote", userHandler.PromoteToAdmin)

		// Movie management
		admin.POST("/movies", movieHandler.CreateMovie)
		admin.PUT("/movies/:id", movieHandler.UpdateMovie)
		admin.DELETE("/movies/:id", movieHandler.DeleteMovie)
		admin.POST("/movies/:id/showtimes", movieHandler.AddShowtime)

		// Reservation management
		admin.GET("/reservations", reservationHandler.ListReservations)
		admin.GET("/revenue", reservationHandler.GetRevenue)
	}

	// Start server
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
