package models

import "time"

type Movie struct {
	ID          uint       `json:"id" gorm:"primaryKey"`
	Title       string     `json:"title" gorm:"not null"`
	Description string     `json:"description"`
	PosterURL   string     `json:"poster_url"`
	Genre       string     `json:"genre"`
	Duration    int        `json:"duration"` // Duration in minutes
	Showtimes   []Showtime `json:"showtimes,omitempty" gorm:"foreignKey:MovieID"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type Showtime struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	MovieID   uint      `json:"movie_id" gorm:"not null"`
	StartTime time.Time `json:"start_time" gorm:"not null"`
	EndTime   time.Time `json:"end_time" gorm:"not null"`
	Theater   string    `json:"theater" gorm:"not null"`
	Seats     []Seat    `json:"seats,omitempty" gorm:"foreignKey:ShowtimeID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
