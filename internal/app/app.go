package app

import (
	v1 "AllowWebsite/internal/controllers/http/handlers/v1"
	"AllowWebsite/internal/methods"
	"AllowWebsite/pkg/logger"
	"AllowWebsite/pkg/memorycache"
	"github.com/gin-gonic/gin"
)

func RunApp() {

	logger.Init()
	logger.InfoLogger.Println("Starting the application...")

	r := gin.Default()

	memoryCache := memorycache.New()

	var objectMaxValue methods.KeyValue

	var objectMinValue methods.KeyValue

	memoryCache.Set("got_ping", 0)
	memoryCache.Set("got_max_ping", 0)
	memoryCache.Set("got_min_ping", 0)

	go methods.UpdateSiteInfo(memoryCache)

	public := r.Group("/v1")

	public.GET("get_ping_from_site", func(context *gin.Context) {

		v1.GetPing(context, memoryCache)

		ping, found := memoryCache.Get("got_ping")
		if found {
			memoryCache.Set("got_ping", ping.(int)+1)
		}

	})

	public.GET("get_site_with_max_ping", func(context *gin.Context) {

		key, _ := memoryCache.Get("max_ping")
		value, _ := memoryCache.Get(key.(string))

		objectMaxValue.Key = key.(string)
		objectMaxValue.Value = value.(float32)

		v1.GetMaxPing(context, objectMaxValue)

		ping, found := memoryCache.Get("got_max_ping")
		if found {
			memoryCache.Set("got_max_ping", ping.(int)+1)
		}

	})

	public.GET("get_site_with_min_ping", func(context *gin.Context) {

		key, _ := memoryCache.Get("min_ping")
		value, _ := memoryCache.Get(key.(string))

		objectMinValue.Key = key.(string)
		objectMinValue.Value = value.(float32)

		v1.GetMinPing(context, objectMinValue)

		ping, found := memoryCache.Get("got_min_ping")
		if found {
			memoryCache.Set("got_min_ping", ping.(int)+1)
		}

	})

	public.GET("get_stat", func(context *gin.Context) {

		v1.GetStatInfo(context, memoryCache)

	})

	err := r.Run()
	if err != nil {
		logger.ErrorLogger.Println(err)
		return
	}

}
