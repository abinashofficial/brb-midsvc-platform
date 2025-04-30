package booking

import(
	"brb-midsvc-platform/model"
	"brb-midsvc-platform/services/booking"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)



type BookingHandler struct {
	service booking.BookingService
}

func NewBookingHandler(s booking.BookingService) *BookingHandler {
	return &BookingHandler{service: s}
}


// @Summary      Create a booking
// @Description  Create a booking by selecting a service and vendor
// @Tags         bookings
// @Accept       json
// @Produce      json
// @Param        booking  body      model.Booking  true  "Booking payload"
// @Success      201      {object}  model.Booking
// @Failure      400      {object}  ErrorResponse
// @Router       /api/bookings/ [post]
func (h *BookingHandler) CreateBooking(c *gin.Context) {
	var booking model.Booking
	if err := c.ShouldBindJSON(&booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.service.CreateBooking(&booking)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, result)
}


// @Summary      List all bookings
// @Description  Fetch a list of all bookings
// @Tags         bookings
// @Produce      json
// @Success      200  {array}   model.Booking
// @Failure      400  {object}  ErrorResponse
// @Router       /api/bookings/ [get]
func (h *BookingHandler) ListBookings(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	bookings, err := h.service.ListBookings(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch bookings"})
		return
	}

	c.JSON(http.StatusOK, bookings)
}


// @Summary      Vendor summary
// @Description  Admin-only summary of bookings for a given vendor
// @Tags         summary
// @Produce      json
// @Param        id   path      int  true  "Vendor ID"
// @Success      200  {object}  model.VendorSummary
// @Failure      400  {object}  ErrorResponse
// @Router       /api/summary/vendor/{id} [get]
func (h *BookingHandler) VendorSummary(c *gin.Context) {
	id := c.Param("id")
	total, statusCounts, err := h.service.GetVendorSummary(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch summary"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"total_bookings": total,
		"status_counts":  statusCounts,
	})
}

