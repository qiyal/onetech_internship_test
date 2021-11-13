package acmp_concurrent

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

func Difficulties(urls []string) map[string]float64 {
	mp := make(map[string]float64)

	wg := sync.WaitGroup{}
	mutex := sync.Mutex{}

	wg.Add(len(urls))
	for _, val := range urls {
		go func(url string) {
			defer wg.Done()
			req, err := http.Get(url)
			defer req.Body.Close()

			if err != nil {
				fmt.Println("url:", url, "not found")
			}

			body, err := ioutil.ReadAll(req.Body)
			var bodyS = string(body)

			index := strings.Index(bodyS, "Сложность:")



			if index != -1 {
				index++;
				var sub = bodyS[index:]
				indexEnd := strings.Index(sub, "%")

				if indexEnd != -1 {
					indexEnd++
					parse, err := strconv.ParseFloat(sub[index:indexEnd], 64)

					if err != nil {
						fmt.Println("error parse float64")
					}
					mutex.Lock()
					mp[url] = parse
					mutex.Unlock()
				}
			}

		}(val)
	}
	wg.Wait()

	return mp
}
