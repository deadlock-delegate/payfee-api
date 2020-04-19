package handlers

import (
	"log"
	"net/http"

	"github.com/deadlock-delegate/payfee-api/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func ListLinks(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var links []models.Link
	if result := db.Find(&links); result.Error != nil {
		log.Println("Problem fetching all links:", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{"error": "There was a problem fetching links"})
		return
	}
	c.JSON(200, links)
}

func CreateLink(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var link models.Link
	if err := c.ShouldBindJSON(&link); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.Create(&link).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, link)
}

func GetLink(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	linkID := c.Param("linkID")
	var link models.Link
	if err := db.Where("id = ?", linkID).First(&link).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, link)
}
