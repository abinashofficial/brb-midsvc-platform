package store

import(
	"brb-midsvc-platform/store/booking"
	"brb-midsvc-platform/store/service"
	"brb-midsvc-platform/store/vendor"
)

	type Store struct {
		BookingRepo booking.BookingRepository
		ServiceRepo service.ServiceRepository
		VendorRepo  vendor.VendorRepository
}