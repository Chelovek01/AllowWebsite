package v1

import (
	"AllowWebsite/internal/controllers/dto"
	"AllowWebsite/internal/domain/service"
	"AllowWebsite/pkg/logger"
	"AllowWebsite/pkg/memorycache"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetStatInfo gets statistic about requests from endpoints
func GetStatInfo(c *gin.Context, cache *memorycache.Cache) {

	var data dto.DtoStat

	err := c.ShouldBindJSON(&data)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}

	serviceWebsiteInfo := service.NewWebsiteInfoService()

	if data.Stat != "get_stat" {
		logger.ErrorLogger.Println("wrong value of data[website]")
		c.JSON(http.StatusBadRequest, gin.H{"stat": "value must be get_stat"})
		return
	}

	stat, err := serviceWebsiteInfo.GetStatistic(cache)
	if err != nil {
		logger.ErrorLogger.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error"})

	}

	c.JSON(http.StatusOK, stat)

}
