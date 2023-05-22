package methods

import (
	"AllowWebsite/pkg/memorycache"
	"bufio"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

// UpdateSiteInfo fills the cache with the necessary data
func UpdateSiteInfo(cache *memorycache.Cache) {

	var wg sync.WaitGroup

	for {

		f, err := os.Open("websiteList.txt")
		if err != nil {
			panic(err)
		}

		var websiteList []string

		sc := bufio.NewScanner(f)
		for sc.Scan() {
			websiteList = append(websiteList, sc.Text())
		}

		f.Close()

		for _, val := range websiteList {

			siteName := val

			wg.Add(1)

			go func() {
				defer wg.Done()

				url := fmt.Sprintf("https://%s", siteName)

				req, err := http.NewRequestWithContext(context.Background(), "tcp", url, nil)
				if err != nil {
					log.Fatal(err)

				}
				client := http.Client{
					Timeout: 5 * time.Second,
				}

				if err != nil {
					log.Fatal(err)
				}

				start := time.Now()
				_, err = client.Do(req)
				reqTime := time.Since(start).Seconds()

				if err != nil {
					fmt.Sprintf("'%s' website not allow", siteName)
					return
				}

				cache.Set(siteName, float32(reqTime))

				_, existMaxPing := cache.Get("max_ping")
				if existMaxPing == false {
					cache.Set("max_ping", siteName)
				} else {
					key, _ := cache.Get("max_ping")
					val, _ := cache.Get(key.(string))
					if float32(reqTime) > val.(float32) {
						cache.Set("max_ping", siteName)
					}
				}

				_, existMinPing := cache.Get("min_ping")
				if existMinPing == false {
					cache.Set("min_ping", siteName)
				} else {
					key, _ := cache.Get("min_ping")
					val, _ := cache.Get(key.(string))
					if float32(reqTime) < val.(float32) {

						cache.Set("min_ping", siteName)
					}
				}

			}()

		}
		wg.Wait()

		time.Sleep(time.Second * 60)
	}

}
