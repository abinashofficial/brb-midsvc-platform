package booking

import (
	"brb-midsvc-platform/model"

	"gorm.io/gorm"
)

type bookingRepository struct {
	db *gorm.DB
}

func NewBookingRepository(db *gorm.DB) BookingRepository {
	return &bookingRepository{db}
}

func (r *bookingRepository) CreateBooking(booking *model.Booking) error {
	return r.db.Create(booking).Error
}

func (r *bookingRepository) FindOverlappingBooking(booking *model.Booking)  (bool, error){
	var count int64
	err := r.db.Model(&model.Booking{}).
		Where("service_id = ? AND vendor_id = ? AND booking_time = ?", booking.ServiceID, booking.VendorID, booking.BookingTime).
		Count(&count).Error
	return  count != 0, err
}

func (r *bookingRepository) ListBookings(offset, limit int) ([]model.Booking, error) {
	var bookings []model.Booking
	err := r.db.Offset(offset).Limit(limit).Find(&bookings).Error
	return bookings, err
}

func (r *bookingRepository) GetBookingSummary(vendorID string) (int64, map[string]int64, error) {
	var total int64
	err := r.db.Model(&model.Booking{}).Where("vendor_id = ?", vendorID).Count(&total).Error
	if err != nil {
		return 0, nil, err
	}

	statusCount := map[string]int64{}
	statuses := []string{"pending", "confirmed", "completed"}
	for _, status := range statuses {
		var count int64
		err := r.db.Model(&model.Booking{}).Where("vendor_id = ? AND status = ?", vendorID, status).Count(&count).Error
		if err != nil {
			return 0, nil, err
		}
		statusCount[status] = count
	}
	return total, statusCount, nil
}
