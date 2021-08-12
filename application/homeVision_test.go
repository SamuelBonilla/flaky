package application

import (
	"errors"
	"fmt"
	"path"
	"testing"
)

func TestRetry(t *testing.T) {
	var retryCounter uint8
	var expectedRetries uint8 = 3
	retry(func () error {
		retryCounter += 1
		return errors.New("error :( ")
	})
	if retryCounter != expectedRetries {
		t.Errorf("expected '%d' but got '%d'", expectedRetries, retryCounter)
	}
}

func TestImageFormat(t *testing.T) {
	t.Run("Validating format: id-[NNN]-[address].[ext]", func(t *testing.T) {
		var imagesUrl []string = []string{
			"https://image.shutterstock.com/image-photo/big-custom-made-luxury-house-260nw-374099713.jpg",
			"https://media-cdn.tripadvisor.com/media/photo-s/09/7c/a2/1f/patagonia-hostel.jpg",
			"https://image.shutterstock.com/image-photo/traditional-english-semidetached-house-260nw-231369511.jpg",
		}
		for i, imageUrl := range imagesUrl {
			var pageNumber, imageId uint8 = uint8(i), 0
			var imageName string = path.Base(imageUrl)
			var expectedFormat string = fmt.Sprintf("images/id-%d-%d-%s", imageId, pageNumber, imageName)

			var format string = imageFormat(int32(imageId), pageNumber, imageUrl)
			if format != expectedFormat {
				t.Errorf("expected '%s' but got '%s'", expectedFormat, format)
			}
		}
	})
}