package service

import (
	"AllowWebsite/internal/domain/entity"
	"AllowWebsite/internal/methods"
	"AllowWebsite/pkg/memorycache"
)

// userService service for object of action
type userService struct {
}

// NewWebsiteInfoService init new service
func NewWebsiteInfoService() *userService {
	return &userService{}
}

// GetPingFromSite returns an action object with site data
func (u *userService) GetPingFromSite(nameOfSite string, m *memorycache.Cache) (*entity.WebsiteInfo, error) {
	var result entity.WebsiteInfo

	result.Website = nameOfSite
	ping, _ := m.Get(nameOfSite)
	result.Ping = ping.(float32)

	return &result, nil
}

// GetSiteNameWithMaxPing returns an action object with data about the site with the maximum ping from the list
func (u *userService) GetSiteNameWithMaxPing(s methods.KeyValue) (*entity.WebsiteInfo, error) {

	var result entity.WebsiteInfo

	result.Website = s.Key
	result.Ping = s.Value

	return &result, nil
}

// GetSiteNameWithMinPing returns an action object with data about the site with the minimum ping from the list
func (u *userService) GetSiteNameWithMinPing(s methods.KeyValue) (*entity.WebsiteInfo, error) {
	var result entity.WebsiteInfo

	result.Website = s.Key
	result.Ping = s.Value

	return &result, nil
}

// GetStatistic returns an action object with data about requests of endpoints
func (u *userService) GetStatistic(cache *memorycache.Cache) (*entity.RequestStat, error) {
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
