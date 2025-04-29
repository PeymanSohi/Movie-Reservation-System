package repository

import (
	"github.com/PeymanSohi/Movie-Reservation-System/internal/models"
	"gorm.io/gorm"
)

type SeatRepository struct {
	db *gorm.DB
}

func NewSeatRepository(db *gorm.DB) *SeatRepository {
	return &SeatRepository{db: db}
}

func (r *SeatRepository) Create(seat *models.Seat) error {
	return r.db.Create(seat).Error
}

func (r *SeatRepository) FindByID(id uint) (*models.Seat, error) {
	var seat models.Seat
	err := r.db.Preload("Showtime").First(&seat, id).Error
	if err != nil {
		return nil, err
	}
	return &seat, nil
}

func (r *SeatRepository) Update(seat *models.Seat) error {
	return r.db.Save(seat).Error
}

func (r *SeatRepository) Delete(id uint) error {
	return r.db.Delete(&models.Seat{}, id).Error
}

func (r *SeatRepository) List() ([]models.Seat, error) {
	var seats []models.Seat
	err := r.db.Preload("Showtime").Find(&seats).Error
	return seats, err
}

func (r *SeatRepository) GetByShowtimeID(showtimeID uint) ([]models.Seat, error) {
	var seats []models.Seat
	err := r.db.Where("showtime_id = ?", showtimeID).Find(&seats).Error
	return seats, err
}

func (r *SeatRepository) UpdateReservationStatus(id uint, isReserved bool) error {
	return r.db.Model(&models.Seat{}).
		Where("id = ?", id).
		Update("is_reserved", isReserved).Error
}
