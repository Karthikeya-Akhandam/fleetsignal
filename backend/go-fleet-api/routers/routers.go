package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/Karthikeya-Akhandam/fleetsignal/backend/go-fleet-api/handlers"
	"github.com/Karthikeya-Akhandam/fleetsignal/backend/go-fleet-api/service"
)

func RegisterRoutes(r *gin.Engine) {
	svc := service.NewFleetService()
	h := handlers.NewFleetHandler(svc)

	api := r.Group("/api/v1")
	{
		api.GET("/incidents", h.GetIncidents)
		api.POST("/incidents/:id/resolve", h.ResolveIncident)
	}
}
