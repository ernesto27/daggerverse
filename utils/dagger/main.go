// A generated module for Utils functions

package main

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
)

type Utils struct{}

// Do http request to the provided url

func (m *Utils) DoRequest(
	// URL to request
	url string,
	// Number of times to request the URL
	times int,
	//+optional
	secuential bool,
) string {
	allResp := []string{}

	if secuential {
		for i := 0; i < times; i++ {
			resp, err := http.Get(url)
			if err != nil {
				fmt.Println("Error: ", err)
				allResp = append(allResp, err.Error())
				continue
			}
			defer resp.Body.Close()
			result := getResponse(i, resp.Status, url)
			fmt.Println(result)
			allResp = append(allResp, result)
		}

	} else {
		wg := sync.WaitGroup{}
		wg.Add(times)

		for i := 0; i < times; i++ {
			go func(url string, i int) {
				defer wg.Done()
				resp, err := http.Get(url)
				if err != nil {
					fmt.Println("Error: ", err)
					allResp = append(allResp, err.Error())
					return
				}
				defer resp.Body.Close()
				result := getResponse(i, resp.Status, url)
				fmt.Println(result)
				allResp = append(allResp, result)
			}(url, i)
		}

		wg.Wait()

	}

	return strings.Join(allResp, "\n")
}

func getResponse(i int, status string, url string) string {
	return fmt.Sprintf("request %d - status %s - %s\n", i, status, url)
}
