package util

import (
	"io/ioutil"
	"net/http"
	"os"
)

func GetProxyScrapeProxies() {
	fileName := "proxies.txt"

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		resp, err := http.Get("https://api.proxyscrape.com/v2/?request=getproxies&protocol=http&timeout=10000")
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		err = ioutil.WriteFile(fileName, body, 0644)
		if err != nil {
			panic(err)
		}

		println("proxies.txt file created successfully.")
	}
}
