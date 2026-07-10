package routes

import (
	"ticket-system/controllers"
	"ticket-system/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	router.GET("/health", controllers.Health)

	auth := router.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
	}

	tickets := router.Group("/tickets")
	tickets.Use(middleware.AuthMiddleware())
	{
		tickets.POST("", controllers.CreateTicket)
		tickets.GET("", controllers.GetTickets)
		tickets.GET("/:id", controllers.GetTicket)
		tickets.PATCH("/:id/status", controllers.UpdateTicketStatus)
	}
}