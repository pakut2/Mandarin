package http_client

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/pakut2/mandarin/pkg/logger"
)

func Get(url string) (map[string]interface{}, error) {
	res, err := http.Get(url)
	if err != nil {
		logger.Logger.Errorf("error fetching %s, err: %v", url, err)
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		logger.Logger.Errorf("request falied with status code %d", res.StatusCode)
		return nil, errors.New("request failed")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		logger.Logger.Errorf("error reading response, err: %v", err)
		return nil, err
	}

	var jsonBody map[string]interface{}
	json.Unmarshal(body, &jsonBody)

	return jsonBody, nil
}
