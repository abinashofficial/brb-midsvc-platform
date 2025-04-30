package vendor

import (
	"brb-midsvc-platform/model"
	"brb-midsvc-platform/services/vendor"
	"net/http"

	"github.com/gin-gonic/gin"
)

type VendorHandler struct {
	service vendor.VendorService
}

func NewVendorHandler(s vendor.VendorService) *VendorHandler {
	return &VendorHandler{service: s}
}

func (h *VendorHandler) RegisterRoutes(r *gin.RouterGroup) {
	vendor := r.Group("/vendors")
	{
		vendor.POST("/", h.CreateVendor)
	}
}


// @Summary      Create a vendor
// @Description  Admin-only endpoint to create a new vendor
// @Tags         vendors
// @Accept       json
// @Produce      json
// @Param        vendor  body      model.Vendor  true  "Vendor payload"
// @Success      201     {object}  model.Vendor
// @Failure      400     {object}  ErrorResponse
// @Router       /api/vendors/ [post]
func (h *VendorHandler) CreateVendor(c *gin.Context) {
	var vendor model.Vendor
	if err := c.ShouldBindJSON(&vendor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.CreateVendor(&vendor); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create vendor"})
		return
	}
	c.JSON(http.StatusCreated, vendor)
}
