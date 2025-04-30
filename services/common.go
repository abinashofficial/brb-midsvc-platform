package services

import(
	"brb-midsvc-platform/services/booking"
	"brb-midsvc-platform/services/service"
	"brb-midsvc-platform/services/vendor"
)

type Store struct {
	BookingService booking.BookingService
	ServiceService service.ServiceService	
	VendorService  vendor.VendorService
}