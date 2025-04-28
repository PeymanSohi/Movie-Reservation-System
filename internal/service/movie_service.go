package service

import (
	"errors"
	"time"

	"github.com/peymansohi/movie-res/internal/models"
	"github.com/peymansohi/movie-res/internal/repository"
)

type MovieService struct {
	movieRepo *repository.MovieRepository
}

func NewMovieService(movieRepo *repository.MovieRepository) *MovieService {
	return &MovieService{
		movieRepo: movieRepo,
	}
}

func (s *MovieService) CreateMovie(movie *models.Movie) error {
	return s.movieRepo.Create(movie)
}

func (s *MovieService) GetMovieByID(id uint) (*models.Movie, error) {
	return s.movieRepo.FindByID(id)
}

func (s *MovieService) UpdateMovie(movie *models.Movie) error {
	return s.movieRepo.Update(movie)
}

func (s *MovieService) DeleteMovie(id uint) error {
	return s.movieRepo.Delete(id)
}

func (s *MovieService) ListMovies() ([]models.Movie, error) {
	return s.movieRepo.List()
}

func (s *MovieService) GetMoviesByDate(date time.Time) ([]models.Movie, error) {
	return s.movieRepo.GetMoviesByDate(date)
}

func (s *MovieService) AddShowtime(movieID uint, showtime *models.Showtime) error {
	// Validate showtime
	if showtime.StartTime.Before(time.Now()) {
		return errors.New("showtime cannot be in the past")
	}

	if showtime.EndTime.Before(showtime.StartTime) {
		return errors.New("end time must be after start time")
	}

	return s.movieRepo.AddShowtime(movieID, showtime)
}

func (s *MovieService) GetShowtimeByID(id uint) (*models.Showtime, error) {
	return s.movieRepo.GetShowtimeByID(id)
}

func (s *MovieService) UpdateShowtime(showtime *models.Showtime) error {
	return s.movieRepo.UpdateShowtime(showtime)
}

func (s *MovieService) DeleteShowtime(id uint) error {
	return s.movieRepo.DeleteShowtime(id)
}

func (s *MovieService) AddSeats(showtimeID uint, seatCount int) error {
	if seatCount <= 0 {
		return errors.New("seat count must be positive")
	}

	seats := make([]models.Seat, seatCount)
	for i := 0; i < seatCount; i++ {
		seats[i] = models.Seat{
			SeatNumber: string(rune('A'+i/10)) + string(rune('0'+i%10)),
			IsReserved: false,
		}
	}

	return s.movieRepo.AddSeats(showtimeID, seats)
}

func (s *MovieService) GetAvailableSeats(showtimeID uint) ([]models.Seat, error) {
	return s.movieRepo.GetAvailableSeats(showtimeID)
}
