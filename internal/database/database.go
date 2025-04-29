package database

import (
	"fmt"
	"log"

	"github.com/PeymanSohi/Movie-Reservation-System/internal/config"
	"github.com/PeymanSohi/Movie-Reservation-System/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	// Auto Migrate the schema
	err = db.AutoMigrate(
		&models.User{},
		&models.Movie{},
		&models.Showtime{},
		&models.Seat{},
		&models.Reservation{},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %v", err)
	}

	return db, nil
}

func SeedAdmin(db *gorm.DB, email, password string) error {
	var count int64
	db.Model(&models.User{}).Where("role = ?", models.RoleAdmin).Count(&count)
	if count > 0 {
		return nil
	}

	admin := &models.User{
		Email: email,
		Role:  models.RoleAdmin,
	}
	if err := admin.HashPassword(password); err != nil {
		return err
	}

	if err := db.Create(admin).Error; err != nil {
		return err
	}

	log.Println("Admin user created successfully")
	return nil
}
