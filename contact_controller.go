package controllers

import (
	"bitespeed/models"
	"bitespeed/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type ContactController struct {
    service *services.ContactService
}

func NewContactController(db *gorm.DB) *ContactController {
    return &ContactController{
        service: services.NewContactService(db),
    }
}

func (ctrl *ContactController) Identify(c *gin.Context) {
    var input models.Contact
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    contact, err := ctrl.service.Identify(&input)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"contact": contact})
}
