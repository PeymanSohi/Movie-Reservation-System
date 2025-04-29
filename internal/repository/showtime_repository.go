package repository

import (
	"time"

	"github.com/PeymanSohi/Movie-Reservation-System/internal/models"
	"gorm.io/gorm"
)

type ShowtimeRepository struct {
	db *gorm.DB
}

func NewShowtimeRepository(db *gorm.DB) *ShowtimeRepository {
	return &ShowtimeRepository{db: db}
}

func (r *ShowtimeRepository) Create(showtime *models.Showtime) error {
	return r.db.Create(showtime).Error
}

func (r *ShowtimeRepository) FindByID(id uint) (*models.Showtime, error) {
	var showtime models.Showtime
	err := r.db.Preload("Movie").Preload("Seats").First(&showtime, id).Error
	if err != nil {
		return nil, err
	}
	return &showtime, nil
}

func (r *ShowtimeRepository) Update(showtime *models.Showtime) error {
	return r.db.Save(showtime).Error
}

func (r *ShowtimeRepository) Delete(id uint) error {
	return r.db.Delete(&models.Showtime{}, id).Error
}

func (r *ShowtimeRepository) List() ([]models.Showtime, error) {
	var showtimes []models.Showtime
	err := r.db.Preload("Movie").Preload("Seats").Find(&showtimes).Error
	return showtimes, err
}

func (r *ShowtimeRepository) GetByDateRange(startDate, endDate time.Time) ([]models.Showtime, error) {
	var showtimes []models.Showtime
	err := r.db.Preload("Movie").Preload("Seats").
		Where("start_time BETWEEN ? AND ?", startDate, endDate).
		Find(&showtimes).Error
	return showtimes, err
}

func (r *ShowtimeRepository) GetByMovieID(movieID uint) ([]models.Showtime, error) {
	var showtimes []models.Showtime
	err := r.db.Preload("Movie").Preload("Seats").
		Where("movie_id = ?", movieID).
		Find(&showtimes).Error
	return showtimes, err
}

func (r *ShowtimeRepository) GetAvailableSeats(showtimeID uint) ([]models.Seat, error) {
	var seats []models.Seat
	err := r.db.Where("showtime_id = ? AND is_reserved = ?", showtimeID, false).
		Find(&seats).Error
	return seats, err
}

func (r *ShowtimeRepository) UpdateSeatStatus(seatID uint, isReserved bool) error {
	return r.db.Model(&models.Seat{}).
		Where("id = ?", seatID).
		Update("is_reserved", isReserved).Error
}
