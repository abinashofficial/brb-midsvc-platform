package service

import (
	"brb-midsvc-platform/model"
	"brb-midsvc-platform/services/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ServiceHandler struct {
	service service.ServiceService
}

func NewServiceHandler(s service.ServiceService) *ServiceHandler {
	return &ServiceHandler{service: s}
}

// @Summary      Create a service
// @Description  Admin-only endpoint to create a new service
// @Tags         services
// @Accept       json
// @Produce      json
// @Param        service  body  model.Service  true  "Service payload"
// @Success      201      {object}  model.Service
// @Failure      400      {object}  ErrorResponse
// @Router       /api/services/ [post]
func (h *ServiceHandler) CreateService(c *gin.Context) {
	var service model.Service
	if err := c.ShouldBindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.CreateService(&service); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create service"})
		return
	}
	c.JSON(http.StatusCreated, service)
}

type ServiceQuery struct {
	ServiceID uint `form:"service_id"`
	VendorID  uint `form:"vendor_id"`
}
// @Summary      Update a service
// @Description  Admin-only endpoint to update a service by ID
// @Tags         services
// @Accept       json
// @Produce      json
// @Param        id       path      int           true  "Service ID"
// @Param        service  body      model.Service  true  "Service payload"
// @Success      200      {object}  model.Service
// @Failure      400      {object}  ErrorResponse
// @Router       /api/services/{id} [put]
func (h *ServiceHandler) UpdateService(c *gin.Context) {

	
	var service model.Service
	if err := c.ShouldBindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := c.Param("id")	

	if err := h.service.UpdateService(&service, id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Service not found"})
		return
	}
	c.JSON(http.StatusOK, service)
}


func (h *ServiceHandler) AssignVendor(c *gin.Context) {

	
	var service model.LinkServiceVendor
	if err := c.ShouldBindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.AssignVendor(&service); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, service)
}


// @Summary      Toggle service status
// @Description  Admin-only endpoint to enable or disable a service
// @Tags         services
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Service ID"
// @Success      200  {object}  model.Service
// @Failure      400  {object}  ErrorResponse
// @Router       /api/services/{id}/toggle [patch]
func (h *ServiceHandler) ToggleService(c *gin.Context) {
	var service model.ServiceVendor
	if err := c.ShouldBindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updated, err := h.service.ToggleService(service)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Service not found"})
		return
	}
	c.JSON(http.StatusOK, updated)
}
