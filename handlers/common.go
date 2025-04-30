package handlers

import (
	"brb-midsvc-platform/handlers/booking"
	"brb-midsvc-platform/handlers/service"
	"brb-midsvc-platform/handlers/vendor"
)

type Store  struct{
	BookingHandler booking.Handler
	ServiceHandler service.Handler
	VendorHandler  vendor.Handler
}