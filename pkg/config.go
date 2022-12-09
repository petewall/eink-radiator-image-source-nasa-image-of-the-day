package pkg

import (
	"fmt"
	"image"
	"os"

	"gopkg.in/yaml.v2"

	"github.com/petewall/eink-radiator-image-source-nasa-image-of-the-day/internal"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate
//counterfeiter:generate . ImageGenerator
type ImageGenerator interface {
	GenerateImage(width, height int) (image.Image, error)
}

type Config struct {
	APIKey string `json:"apiKey" yaml:"apiKey"`
	Date   string `json:"date,omitempty" yaml:"date,omitempty"`
}

func (c *Config) GenerateImage(width, height int) (image.Image, error) {
	url, err := internal.GetImageOfTheDay(c.APIKey, c.Date)
	if err != nil {
		return nil, err
	}

	return internal.ProcessImage(url, width, height)
}

func (c *Config) Validate() error {
	if c.APIKey == "" {
		return fmt.Errorf("missing api key")
	}

	return nil
}

func ParseConfig(path string) (*Config, error) {
	configData, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read image config file: %w", err)
	}

	var config *Config
	err = yaml.Unmarshal(configData, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to parse image config file: %w", err)
	}

	err = config.Validate()
	if err != nil {
		return nil, fmt.Errorf("config file is not valid: %w", err)
	}

	return config, nil
}
