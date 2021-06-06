package v1

import (
	"errors"
	"net/http"

	log "github.com/dhawton/log4g"
	"github.com/gin-gonic/gin"
	"github.com/nzvirtual/go-api/lib/database/models"
	"gorm.io/gorm"
)

func GetAirport(c *gin.Context) {
	icao := c.Param("icao")
	airport := models.Airport{}
	if err := models.DB.Where("icao = ?", icao).First(&airport).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		log.Category("api/airport").Fatal("Error encountered during record lookup: " + err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OK", "data": airport})
}
