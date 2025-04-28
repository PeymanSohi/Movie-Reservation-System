package repository

import (
	"time"

	"github.com/peymansohi/movie-res/internal/models"
	"gorm.io/gorm"
)

type ReservationRepository struct {
	db *gorm.DB
}

func NewReservationRepository(db *gorm.DB) *ReservationRepository {
	return &ReservationRepository{db: db}
}

func (r *ReservationRepository) Create(reservation *models.Reservation) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// Check if seat is available
		var seat models.Seat
		if err := tx.First(&seat, reservation.SeatID).Error; err != nil {
			return err
		}

		if seat.IsReserved {
			return models.ErrSeatAlreadyReserved
		}

		// Update seat status
		seat.IsReserved = true
		if err := tx.Save(&seat).Error; err != nil {
			return err
		}

		// Create reservation
		return tx.Create(reservation).Error
	})
}

func (r *ReservationRepository) FindByID(id uint) (*models.Reservation, error) {
	var reservation models.Reservation
	err := r.db.Preload("Seat.Showtime.Movie").First(&reservation, id).Error
	if err != nil {
		return nil, err
	}
	return &reservation, nil
}

func (r *ReservationRepository) FindByUserID(userID uint) ([]models.Reservation, error) {
	var reservations []models.Reservation
	err := r.db.Preload("Seat.Showtime.Movie").
		Where("user_id = ?", userID).
		Find(&reservations).Error
	return reservations, err
}

func (r *ReservationRepository) Cancel(id uint) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		var reservation models.Reservation
		if err := tx.First(&reservation, id).Error; err != nil {
			return err
		}

		// Check if showtime has passed
		var seat models.Seat
		if err := tx.Preload("Showtime").First(&seat, reservation.SeatID).Error; err != nil {
			return err
		}

		if seat.Showtime.StartTime.Before(time.Now()) {
			return models.ErrShowtimePassed
		}

		// Update seat status
		seat.IsReserved = false
		if err := tx.Save(&seat).Error; err != nil {
			return err
		}

		// Update reservation status
		reservation.Status = "cancelled"
		return tx.Save(&reservation).Error
	})
}

func (r *ReservationRepository) List() ([]models.Reservation, error) {
	var reservations []models.Reservation
	err := r.db.Preload("Seat.Showtime.Movie").
		Preload("User").
		Find(&reservations).Error
	return reservations, err
}

func (r *ReservationRepository) GetRevenue(startDate, endDate time.Time) (float64, error) {
	var total float64
	err := r.db.Model(&models.Reservation{}).
		Where("created_at BETWEEN ? AND ?", startDate, endDate).
		Select("COALESCE(SUM(price), 0)").
		Row().
		Scan(&total)
	return total, err
}
