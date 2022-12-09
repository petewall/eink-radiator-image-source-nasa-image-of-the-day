package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

//counterfeiter:generate . ImageOfTheDayGetter
type ImageOfTheDayGetter func(apiKey, date string) (string, error)

type APODResponse struct {
	ServiceVersion string `json:"service_version"`
	Title          string `json:"title"`
	Date           string `json:"date"`
	Explanation    string `json:"explanation"`
	Type           string `json:"media_type"`
	Copyright      string `json:"copyright"`
	URL            string `json:"url"`
	Thumbnail      string `json:"thumbnail_url"`
}

var GetImageOfTheDay ImageOfTheDayGetter = func(apiKey, date string) (string, error) {
	url := fmt.Sprintf("https://api.nasa.gov/planetary/apod?api_key=%s&thumbs=true&date=%s", apiKey, date)
	res, err := HttpGet(url)
	if err != nil {
		return "", fmt.Errorf("failed to query image of the day (%s): %w", url, err)
	}

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to request image of the day: %s: %w", http.StatusText(res.StatusCode), err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read image of the day response: %w", err)
	}

	var response APODResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", fmt.Errorf("failed to decode image of the day response: %w", err)
	}

	if response.Type == "image" {
		return response.URL, nil
	} else if response.Type == "video" {
		return response.Thumbnail, nil
	}

	return "", fmt.Errorf("unexpected image type %s", response.Type)
}
