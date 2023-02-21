package repository

import "github.com/gjaubert/bookings-project/internal/models"

type DatabaseRepo interface {
	AllUsers() bool

	InsertReservation(res models.Reservation) error
}
