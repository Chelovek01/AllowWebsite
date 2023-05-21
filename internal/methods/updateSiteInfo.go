package methods

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sort"
	"sync"
	"time"
)

func UpdateSiteInfo(websiteList []string, resultMap map[string]float32, maxValue *KeyValue, minValue *KeyValue) {

	var wg sync.WaitGroup

	var sortedStruct []KeyValue

	for {

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
				resp, err := client.Do(req)

				if err != nil {
					fmt.Sprintf("'%s' website not allow", siteName)
					return
				}

				reqTime := time.Since(start).Seconds()

				res := fmt.Sprintf("Status: %d | time: %2f", resp.StatusCode, reqTime)
				fmt.Println(res)

				resultMap[siteName] = float32(reqTime)

			}()

		}
		wg.Wait()

		for key, val := range resultMap {
			sortedStruct = append(sortedStruct, KeyValue{key, val})
		}

		sort.Slice(sortedStruct, func(i, j int) bool {
			return sortedStruct[i].Value > sortedStruct[j].Value
		})

		maxValue.Value = sortedStruct[0].Value
		maxValue.Key = sortedStruct[0].Key

		minValue.Value = sortedStruct[len(sortedStruct)-1].Value
		minValue.Key = sortedStruct[len(sortedStruct)-1].Key

		time.Sleep(time.Second * 60)
	}

}
