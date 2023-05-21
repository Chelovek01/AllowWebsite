package app

import (
	v1 "AllowWebsite/internal/controllers/http/handlers/v1"
	"AllowWebsite/internal/methods"
	"AllowWebsite/pkg/logger"
	"bufio"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"os"
)

func RunApp() {

	logger.Init()
	logger.InfoLogger.Println("Starting the application...")

	r := gin.Default()

	f, err := os.Open("websiteList.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var websiteList []string

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		websiteList = append(websiteList, sc.Text())
	}

	resultMap := make(map[string]float32)

	var objectMaxValue methods.KeyValue

	var objectMinValue methods.KeyValue

	c := cache.New(-1, -1)
	c.Set("got_ping", 0, cache.NoExpiration)
	c.Set("got_max_ping", 0, cache.NoExpiration)
	c.Set("got_min_ping", 0, cache.NoExpiration)

	go methods.UpdateSiteInfo(websiteList, resultMap, &objectMaxValue, &objectMinValue)

	public := r.Group("/v1")

	public.GET("get_ping_from_site", func(context *gin.Context) {
		v1.GetPing(context, resultMap)

		ping, found := c.Get("got_ping")
		if found {
			c.Set("got_ping", ping.(int)+1, cache.NoExpiration)
			fmt.Println(ping)
		}

	})

	public.GET("get_site_with_max_ping", func(context *gin.Context) {
		v1.GetMaxPing(context, objectMaxValue)

		ping, found := c.Get("got_max_ping")
		if found {
			c.Set("got_max_ping", ping.(int)+1, cache.NoExpiration)
		}

	})

	public.GET("get_site_with_min_ping", func(context *gin.Context) {
		v1.GetMinPing(context, objectMinValue)
		ping, found := c.Get("got_min_ping")
		if found {
			c.Set("got_min_ping", ping.(int)+1, cache.NoExpiration)
			fmt.Println(ping)
		}

	})

	public.GET("get_stat", func(context *gin.Context) {

		v1.GetStatInfo(context, c)

	})

	err = r.Run()
	if err != nil {
		logger.ErrorLogger.Println(err)
		return
	}

}
