package booking


import(
	"brb-midsvc-platform/model"
	"brb-midsvc-platform/store/booking"
	"errors"
	"time"
)
type bookingService struct {
	repo booking.BookingRepository
}

func NewBookingService(r booking.BookingRepository) BookingService {
	return &bookingService{repo: r}
}

func (s *bookingService) CreateBooking(booking *model.Booking) (*model.Booking, error) {
	hour := booking.BookingTime.Hour()


	if hour < 9 || hour > 16  || booking.BookingTime.Minute() != 0 ||  booking.BookingTime.Second() != 0 {
		return nil, errors.New("booking must be between 9 AM and 5 PM in 1-hour intervals ")
	}

	if valid, _ := s.repo.FindOverlappingBooking(booking); valid {
		return nil, errors.New("slot already booked")
	}

	booking.Status = "pending"
	maxRetries := 3
	for i := 0; i < maxRetries; i++ {
		if err := s.repo.CreateBooking(booking); err == nil {
			return booking, nil
		}
		time.Sleep(500 * time.Millisecond)
	}
	return nil, errors.New("failed to create booking after retries")
}

func (s *bookingService) ListBookings(page, limit int) ([]model.Booking, error) {
	offset := (page - 1) * limit
	return s.repo.ListBookings(offset, limit)
}

func (s *bookingService) GetVendorSummary(vendorID string) (int64, map[string]int64, error) {
	return s.repo.GetBookingSummary(vendorID)
}