package update

import (
	"io/ioutil"
	"net/http"
)

func GetLatestVersion() (string, error) {
	url := "https://raw.githubusercontent.com/wouldblease/roblox-tools/main/update/version.txt"
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	latestVersion := string(body)
	return latestVersion, nil
}
