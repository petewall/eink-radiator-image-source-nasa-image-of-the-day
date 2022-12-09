package pkg_test

import (
	"encoding/json"
	"errors"
	"image"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gopkg.in/yaml.v2"

	"github.com/petewall/eink-radiator-image-source-nasa-image-of-the-day/internal"
	"github.com/petewall/eink-radiator-image-source-nasa-image-of-the-day/internal/internalfakes"
	"github.com/petewall/eink-radiator-image-source-nasa-image-of-the-day/pkg"
)

var _ = Describe("Config", func() {
	Describe("GenerateImage", func() {
		var (
			imageGetter    *internalfakes.FakeImageOfTheDayGetter
			imageProcessor *internalfakes.FakeImageProcessor
			processedImage *image.RGBA
		)

		BeforeEach(func() {
			imageGetter = &internalfakes.FakeImageOfTheDayGetter{}
			imageGetter.Returns("http://nasa.example.gov/imageoftheday.jpg", nil)
			internal.GetImageOfTheDay = imageGetter.Spy

			processedImage = image.NewRGBA(image.Rect(0, 0, 1024, 768))
			imageProcessor = &internalfakes.FakeImageProcessor{}
			imageProcessor.Returns(processedImage, nil)
			internal.ProcessImage = imageProcessor.Spy
		})

		It("fetches the image of the day and returns it processed", func() {
			config := &pkg.Config{
				APIKey: "my-api-key",
			}

			img, err := config.GenerateImage(1024, 768)
			Expect(err).ToNot(HaveOccurred())

			By("fetching the image", func() {
				Expect(imageGetter.CallCount()).To(Equal(1))
				apiKey, date := imageGetter.ArgsForCall(0)
				Expect(apiKey).To(Equal("my-api-key"))
				Expect(date).To(Equal(""))
			})

			By("processing the image", func() {
				Expect(imageProcessor.CallCount()).To(Equal(1))
				url, width, height := imageProcessor.ArgsForCall(0)
				Expect(url).To(Equal("http://nasa.example.gov/imageoftheday.jpg"))
				Expect(width).To(Equal(1024))
				Expect(height).To(Equal(768))
			})

			By("returning the processed image", func() {
				Expect(img).To(Equal(processedImage))
			})
		})

		When("getting the image fails", func() {
			BeforeEach(func() {
				imageGetter.Returns("", errors.New("image get failed"))
			})

			It("returns an error", func() {
				config := &pkg.Config{
					APIKey: "my-api-key",
				}

				_, err := config.GenerateImage(200, 300)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("image get failed"))
			})
		})

		When("processing the image fails", func() {
			BeforeEach(func() {
				imageProcessor.Returns(nil, errors.New("image processing failed"))
			})

			It("returns an error", func() {
				config := &pkg.Config{
					APIKey: "my-api-key",
				}

				_, err := config.GenerateImage(200, 300)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("image processing failed"))
			})
		})
	})
})

var _ = Describe("ParseConfig", func() {
	var (
		configFile         *os.File
		configFileContents []byte
	)

	JustBeforeEach(func() {
		var err error
		configFile, err = os.CreateTemp("", "blank-config.yaml")
		Expect(err).ToNot(HaveOccurred())
		_, err = configFile.Write(configFileContents)
		Expect(err).ToNot(HaveOccurred())
	})

	BeforeEach(func() {
		config := pkg.Config{
			APIKey: "my-api-key",
		}
		var err error
		configFileContents, err = yaml.Marshal(config)
		Expect(err).ToNot(HaveOccurred())
	})

	AfterEach(func() {
		Expect(os.Remove(configFile.Name())).To(Succeed())
	})

	It("parses the image config file", func() {
		config, err := pkg.ParseConfig(configFile.Name())
		Expect(err).ToNot(HaveOccurred())
		Expect(config.APIKey).To(Equal("my-api-key"))
		Expect(config.Date).To(BeEmpty())
	})

	Context("config file is json formatted", func() {
		BeforeEach(func() {
			config := pkg.Config{
				APIKey: "my-api-key",
				Date:   "yesterday",
			}
			var err error
			configFileContents, err = json.Marshal(config)
			Expect(err).ToNot(HaveOccurred())
		})

		It("parses just fine", func() {
			config, err := pkg.ParseConfig(configFile.Name())
			Expect(err).ToNot(HaveOccurred())
			Expect(config.APIKey).To(Equal("my-api-key"))
			Expect(config.Date).To(Equal("yesterday"))
		})
	})

	When("reading the config file fails", func() {
		It("returns an error", func() {
			_, err := pkg.ParseConfig("this file does not exist")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("failed to read image config file: open this file does not exist: no such file or directory"))
		})
	})

	When("parsing the config file fails", func() {
		BeforeEach(func() {
			configFileContents = []byte("this is invalid yaml!")
		})

		It("returns an error", func() {
			_, err := pkg.ParseConfig(configFile.Name())
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("failed to parse image config file: yaml: unmarshal errors:\n  line 1: cannot unmarshal !!str `this is...` into pkg.Config"))
		})
	})

	When("the config file has missing api key", func() {
		BeforeEach(func() {
			config := pkg.Config{}
			var err error
			configFileContents, err = json.Marshal(config)
			Expect(err).ToNot(HaveOccurred())
		})

		It("returns an error", func() {
			_, err := pkg.ParseConfig(configFile.Name())
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("config file is not valid: missing api key"))
		})
	})
})
