package service

import (
	"AllowWebsite/internal/domain/entity"
	"AllowWebsite/internal/methods"
	"github.com/patrickmn/go-cache"
)

type userService struct {
}

func NewWebsiteInfoService() *userService {
	return &userService{}
}

func (u *userService) GetPingFromSite(nameOfSite string, m map[string]float32) (*entity.WebsiteInfo, error) {
	var result entity.WebsiteInfo

	result.Website = nameOfSite
	result.Ping = m[nameOfSite]

	return &result, nil
}

func (u *userService) GetSiteNameWithMaxPing(s methods.KeyValue) (*entity.WebsiteInfo, error) {

	var result entity.WebsiteInfo

	result.Website = s.Key
	result.Ping = s.Value

	return &result, nil
}

func (u *userService) GetSiteNameWithMinPing(s methods.KeyValue) (*entity.WebsiteInfo, error) {
	var result entity.WebsiteInfo

	result.Website = s.Key
	result.Ping = s.Value

	return &result, nil
}

func (u *userService) GetStatistic(cache *cache.Cache) (*entity.RequestStat, error) {
	var result entity.RequestStat

	ping, found := cache.Get("got_ping")
	if found {
		result.GotPing = ping.(int)
	}

	maxPing, found := cache.Get("got_max_ping")
	if found {

		result.GotMaxPing = maxPing.(int)
	}

	minPing, found := cache.Get("got_min_ping")
	if found {

		result.GotMinPing = minPing.(int)
	}

	return &result, nil
}
