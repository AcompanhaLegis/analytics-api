package dadosabertos

import (
	"io"
	"net/http"
)

var BASE_URL = "https://dadosabertos.camara.leg.br/api/v2"

type Link struct {
	Rel  string `json:"rel"`
	Href string `json:"href"`
}

func GetRequest(endpoint string) ([]byte, error) {
	// Set request headers
	logger.Debugf("Getting %s", endpoint)

	client := &http.Client{}
	url := BASE_URL + endpoint
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var body []byte
	body, err = io.ReadAll(res.Body)

	if err != nil {
		logger.Debugf("Error reading response body: %v", err)
		return nil, err
	}
	return body, nil
}
