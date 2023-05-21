package v1

import (
	"AllowWebsite/internal/controllers/dto"
	"AllowWebsite/internal/domain/service"
	"AllowWebsite/internal/methods"
	"AllowWebsite/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetMinPing(c *gin.Context, s methods.KeyValue) {

	var data dto.DtoMinPing

	err := c.ShouldBindJSON(&data)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}

	serviceWebsiteInfo := service.NewWebsiteInfoService()

	if data.Website != "min_ping" {
		logger.ErrorLogger.Println("wrong value of data[website]")
		c.JSON(http.StatusBadRequest, gin.H{"website": "value must be min_ping"})
		return
	}

	SiteWithMinPing, err := serviceWebsiteInfo.GetSiteNameWithMinPing(s)
	if err != nil {
		logger.ErrorLogger.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error"})
	}

	c.JSON(http.StatusOK, SiteWithMinPing)

}
