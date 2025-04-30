package repository

import (
	"github.com/PeymanSohi/Movie-Reservation-System/internal/models"
	"gorm.io/gorm"
)

type MovieRepository struct {
	db *gorm.DB
}

func NewMovieRepository(db *gorm.DB) *MovieRepository {
	return &MovieRepository{db: db}
}

func (r *MovieRepository) GetShowtimeByID(id uint) (*models.Showtime, error) {
	var showtime models.Showtime
	err := r.db.First(&showtime, id).Error
	if err != nil {
		return nil, err
	}
	return &showtime, nil
}

func (r *MovieRepository) CreateMovie(movie *models.Movie) error {
	return r.db.Create(movie).Error
}

func (r *MovieRepository) GetMovieByID(id uint) (*models.Movie, error) {
	var movie models.Movie
	err := r.db.First(&movie, id).Error
	if err != nil {
		return nil, err
	}
	return &movie, nil
}

func (r *MovieRepository) ListMovies() ([]models.Movie, error) {
	var movies []models.Movie
	err := r.db.Find(&movies).Error
	return movies, err
}

func (r *MovieRepository) UpdateMovie(movie *models.Movie) error {
	return r.db.Save(movie).Error
}

func (r *MovieRepository) DeleteMovie(id uint) error {
	return r.db.Delete(&models.Movie{}, id).Error
}
