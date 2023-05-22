package v1

import (
	"AllowWebsite/internal/controllers/dto"
	"AllowWebsite/internal/domain/service"
	"AllowWebsite/pkg/logger"
	"AllowWebsite/pkg/memorycache"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetPing(c *gin.Context, m *memorycache.Cache) {

	var data dto.DtoSiteName

	err := c.ShouldBindJSON(&data)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}

	_, ok := m.Get(data.SiteName)
	if ok == false {
		text := fmt.Sprintf("website %s not avalible ", data.SiteName)
		c.JSON(http.StatusOK, gin.H{"response": text})
		return
	}

	serviceWebsiteInfo := service.NewWebsiteInfoService()
	SitePing, err := serviceWebsiteInfo.GetPingFromSite(data.SiteName, m)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}

	c.JSON(http.StatusOK, SitePing)

}
