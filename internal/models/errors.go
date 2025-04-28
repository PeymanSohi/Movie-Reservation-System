package models

import "errors"

var (
	ErrSeatAlreadyReserved = errors.New("seat is already reserved")
	ErrShowtimePassed      = errors.New("showtime has already passed")
	ErrInvalidSeat         = errors.New("invalid seat")
	ErrInvalidShowtime     = errors.New("invalid showtime")
)
