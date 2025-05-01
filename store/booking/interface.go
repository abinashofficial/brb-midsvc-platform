package booking

import (
	"brb-midsvc-platform/model"
)

type BookingRepository interface {
	CreateBooking(booking *model.Booking) error
	FindOverlappingBooking(booking *model.Booking) (bool, error)
	ListBookings(offset, limit int) ([]model.Booking, error)
	GetBookingSummary(vendorID string) (int64, map[string]int64, error)
}
