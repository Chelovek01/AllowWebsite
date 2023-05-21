package v1

import (
	"AllowWebsite/internal/controllers/dto"
	"AllowWebsite/internal/domain/service"
	"AllowWebsite/internal/methods"
	"AllowWebsite/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetMaxPing(c *gin.Context, s methods.KeyValue) {

	var data dto.DtoMaxPing

	err := c.ShouldBindJSON(&data)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}

	serviceWebsiteInfo := service.NewWebsiteInfoService()

	if data.Website != "max_ping" {
		logger.ErrorLogger.Println("wrong value of data[website]")
		c.JSON(http.StatusBadRequest, gin.H{"website": "value must be max_ping"})
		return
	}

	SiteWithMaxPing, err := serviceWebsiteInfo.GetSiteNameWithMaxPing(s)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}

	c.JSON(http.StatusOK, SiteWithMaxPing)

}
