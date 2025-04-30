package repository

import (
	"time"

	"github.com/PeymanSohi/Movie-Reservation-System/internal/models"
	"gorm.io/gorm"
)

type ReservationRepository struct {
	db *gorm.DB
}

func NewReservationRepository(db *gorm.DB) *ReservationRepository {
	return &ReservationRepository{db: db}
}

func (r *ReservationRepository) Create(reservation *models.Reservation) error {
	return r.db.Create(reservation).Error
}

func (r *ReservationRepository) FindByID(id uint) (*models.Reservation, error) {
	var reservation models.Reservation
	err := r.db.First(&reservation, id).Error
	if err != nil {
		return nil, err
	}
	return &reservation, nil
}

func (r *ReservationRepository) FindByUserID(userID uint) ([]models.Reservation, error) {
	var reservations []models.Reservation
	err := r.db.Where("user_id = ?", userID).Find(&reservations).Error
	return reservations, err
}

func (r *ReservationRepository) Cancel(id uint) error {
	return r.db.Model(&models.Reservation{}).Where("id = ?", id).Update("status", "cancelled").Error
}

func (r *ReservationRepository) List() ([]models.Reservation, error) {
	var reservations []models.Reservation
	err := r.db.Find(&reservations).Error
	return reservations, err
}

func (r *ReservationRepository) GetRevenue(startDate, endDate time.Time) (float64, error) {
	var total float64
	err := r.db.Model(&models.Reservation{}).
		Where("created_at BETWEEN ? AND ? AND status = ?", startDate, endDate, "active").
		Select("COALESCE(SUM(amount), 0)").
		Scan(&total).Error
	return total, err
}
