package service

import (
	"time"

	"github.com/peymansohi/movie-res/internal/models"
	"github.com/peymansohi/movie-res/internal/repository"
)

type ReservationService struct {
	reservationRepo *repository.ReservationRepository
	movieRepo       *repository.MovieRepository
}

func NewReservationService(reservationRepo *repository.ReservationRepository, movieRepo *repository.MovieRepository) *ReservationService {
	return &ReservationService{
		reservationRepo: reservationRepo,
		movieRepo:       movieRepo,
	}
}

func (s *ReservationService) CreateReservation(userID, seatID uint) (*models.Reservation, error) {
	// Check if seat exists and is available
	showtime, err := s.movieRepo.GetShowtimeByID(seatID)
	if err != nil {
		return nil, err
	}

	// Check if showtime has passed
	if showtime.StartTime.Before(time.Now()) {
		return nil, models.ErrShowtimePassed
	}

	reservation := &models.Reservation{
		UserID: userID,
		SeatID: seatID,
		Status: "active",
	}

	if err := s.reservationRepo.Create(reservation); err != nil {
		return nil, err
	}

	return reservation, nil
}

func (s *ReservationService) GetReservationByID(id uint) (*models.Reservation, error) {
	return s.reservationRepo.FindByID(id)
}

func (s *ReservationService) GetUserReservations(userID uint) ([]models.Reservation, error) {
	return s.reservationRepo.FindByUserID(userID)
}

func (s *ReservationService) CancelReservation(id uint) error {
	return s.reservationRepo.Cancel(id)
}

func (s *ReservationService) ListReservations() ([]models.Reservation, error) {
	return s.reservationRepo.List()
}

func (s *ReservationService) GetRevenue(startDate, endDate time.Time) (float64, error) {
	return s.reservationRepo.GetRevenue(startDate, endDate)
}
