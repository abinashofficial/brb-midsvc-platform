package app

import(
	"brb-midsvc-platform/config"
	"github.com/gin-gonic/gin"
	book_handlers "brb-midsvc-platform/handlers/booking"
	service_handlers "brb-midsvc-platform/handlers/service"
	vendor_handlers "brb-midsvc-platform/handlers/vendor"
	 "brb-midsvc-platform/services"
	 "brb-midsvc-platform/handlers"
	 book_services "brb-midsvc-platform/services/booking"
	 service_service "brb-midsvc-platform/services/service"
	 vendor_service "brb-midsvc-platform/services/vendor"
	 repo "brb-midsvc-platform/store"
	 repo_booking "brb-midsvc-platform/store/booking"
	 repo_service "brb-midsvc-platform/store/service"
	 repo_vendor "brb-midsvc-platform/store/vendor"
	 "gorm.io/gorm"

)

var service services.Store
var h handlers.Store
var r repo.Store


func setupHandlers() {
	h = handlers.Store{
		BookingHandler: book_handlers.NewBookingHandler(service.BookingService),
		ServiceHandler: service_handlers.NewServiceHandler(service.ServiceService),	
		VendorHandler:  vendor_handlers.NewVendorHandler(service.VendorService),
	}
}

func setupService(){
	service = services.Store{
		BookingService: book_services.NewBookingService(r.BookingRepo),
		ServiceService: service_service.NewServiceService(r.ServiceRepo),
		VendorService:  vendor_service.NewVendorService(r.VendorRepo),
	}
}

func setupRepo(db *gorm.DB) {
	r = repo.Store{
		BookingRepo: repo_booking.NewBookingRepository(db),
		ServiceRepo: repo_service.NewServiceRepository(db),
		VendorRepo:  repo_vendor.NewVendorRepository(db),
	}
}


func Start() {
	db := config.SetupDatabase()
	setupRepo(db)
	setupService()
	setupHandlers()


	// Setup Gin
	g := gin.Default()
	runServer(g, db, h)
}