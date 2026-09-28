package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/Karthikeya-Akhandam/fleetsignal/backend/go-fleet-api/models"
	"github.com/Karthikeya-Akhandam/fleetsignal/backend/go-fleet-api/service"
)

type FleetHandler struct {
	Service *service.FleetService
}

func NewFleetHandler(svc *service.FleetService) *FleetHandler {
	return &FleetHandler{Service: svc}
}

func (h *FleetHandler) GetIncidents(c *gin.Context) {
	incidents, err := h.Service.GetActiveIncidents(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	if incidents == nil {
		incidents = []models.Incident{} // Return empty array instead of null
	}
	c.JSON(http.StatusOK, incidents)
}

func (h *FleetHandler) ResolveIncident(c *gin.Context) {
	incidentID := c.Param("id")
	
	var body struct {
		Notes string `json:"notes"`
	}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := h.Service.ResolveIncident(c.Request.Context(), incidentID, body.Notes); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Incident resolved"})
}
