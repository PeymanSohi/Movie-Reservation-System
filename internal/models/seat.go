package models

import "time"

type Seat struct {
	ID          uint         `json:"id" gorm:"primaryKey"`
	ShowtimeID  uint         `json:"showtime_id" gorm:"not null"`
	Showtime    Showtime     `json:"showtime,omitempty" gorm:"foreignKey:ShowtimeID"`
	SeatNumber  string       `json:"seat_number" gorm:"not null"`
	IsReserved  bool         `json:"is_reserved" gorm:"default:false"`
	Reservation *Reservation `json:"reservation,omitempty" gorm:"foreignKey:SeatID"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}
