package controllers

import (
	"net/http"

	"ticket-system/config"
	"ticket-system/models"

	"github.com/gin-gonic/gin"
)

type CreateTicketRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func CreateTicket(c *gin.Context) {

	var req CreateTicketRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID := c.MustGet("user_id").(uint)

	ticket := models.Ticket{
		Title:       req.Title,
		Description: req.Description,
		Status:      "open",
		UserID:      userID,
	}

	if err := config.DB.Create(&ticket).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to create ticket",
		})
		return
	}

	c.JSON(http.StatusCreated, ticket)
}

func GetTickets(c *gin.Context) {

	userID := c.MustGet("user_id").(uint)

	var tickets []models.Ticket

	if err := config.DB.Where("user_id = ?", userID).Find(&tickets).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to fetch tickets",
		})

		return
	}

	c.JSON(http.StatusOK, tickets)
}
func GetTicket(c *gin.Context) {

	userID := c.MustGet("user_id").(uint)

	id := c.Param("id")

	var ticket models.Ticket

	if err := config.DB.Where("id = ? AND user_id = ?", id, userID).First(&ticket).Error; err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"error": "Ticket not found",
		})

		return
	}

	c.JSON(http.StatusOK, ticket)
}

type UpdateStatusRequest struct {
	Status string `json:"status"`
}

func UpdateTicketStatus(c *gin.Context) {

	userID := c.MustGet("user_id").(uint)

	id := c.Param("id")

	var req UpdateStatusRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	var ticket models.Ticket

	if err := config.DB.Where("id = ? AND user_id = ?", id, userID).First(&ticket).Error; err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"error": "Ticket not found",
		})

		return
	}

	switch ticket.Status {

	case "open":
		if req.Status != "in_progress" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "open ticket can only move to in_progress",
			})
			return
		}

	case "in_progress":
		if req.Status != "closed" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "in_progress ticket can only move to closed",
			})
			return
		}

	case "closed":
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "closed ticket cannot be reopened",
		})
		return
	}

	ticket.Status = req.Status

	config.DB.Save(&ticket)

	c.JSON(http.StatusOK, ticket)
}