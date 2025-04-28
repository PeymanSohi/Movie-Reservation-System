package repository

import (
	"time"

	"github.com/peymansohi/movie-res/internal/models"
	"gorm.io/gorm"
)

type MovieRepository struct {
	db *gorm.DB
}

func NewMovieRepository(db *gorm.DB) *MovieRepository {
	return &MovieRepository{db: db}
}

func (r *MovieRepository) Create(movie *models.Movie) error {
	return r.db.Create(movie).Error
}

func (r *MovieRepository) FindByID(id uint) (*models.Movie, error) {
	var movie models.Movie
	err := r.db.Preload("Showtimes").First(&movie, id).Error
	if err != nil {
		return nil, err
	}
	return &movie, nil
}

func (r *MovieRepository) Update(movie *models.Movie) error {
	return r.db.Save(movie).Error
}

func (r *MovieRepository) Delete(id uint) error {
	return r.db.Delete(&models.Movie{}, id).Error
}

func (r *MovieRepository) List() ([]models.Movie, error) {
	var movies []models.Movie
	err := r.db.Preload("Showtimes").Find(&movies).Error
	return movies, err
}

func (r *MovieRepository) GetMoviesByDate(date time.Time) ([]models.Movie, error) {
	var movies []models.Movie
	startOfDay := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	endOfDay := startOfDay.Add(24 * time.Hour)

	err := r.db.Preload("Showtimes", "start_time >= ? AND start_time < ?", startOfDay, endOfDay).
		Find(&movies).Error
	return movies, err
}

func (r *MovieRepository) AddShowtime(movieID uint, showtime *models.Showtime) error {
	showtime.MovieID = movieID
	return r.db.Create(showtime).Error
}

func (r *MovieRepository) GetShowtimeByID(id uint) (*models.Showtime, error) {
	var showtime models.Showtime
	err := r.db.Preload("Seats").First(&showtime, id).Error
	if err != nil {
		return nil, err
	}
	return &showtime, nil
}

func (r *MovieRepository) UpdateShowtime(showtime *models.Showtime) error {
	return r.db.Save(showtime).Error
}

func (r *MovieRepository) DeleteShowtime(id uint) error {
	return r.db.Delete(&models.Showtime{}, id).Error
}

func (r *MovieRepository) AddSeats(showtimeID uint, seats []models.Seat) error {
	for i := range seats {
		seats[i].ShowtimeID = showtimeID
	}
	return r.db.Create(&seats).Error
}

func (r *MovieRepository) GetAvailableSeats(showtimeID uint) ([]models.Seat, error) {
	var seats []models.Seat
	err := r.db.Where("showtime_id = ? AND is_reserved = ?", showtimeID, false).Find(&seats).Error
	return seats, err
}
