package application

import (
	"encoding/json"
	"errors"
	"fmt"
	"homeVision/domain"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"sync"
)

const (
	Folder  string = "images"
	retries uint8  = 3
)

type ImageInfo struct {
	Url      string
	Filename string
}

func retry(callback func() error) error {
	var err error
	var i uint8 = 0
	for ; i < retries; i++ {
		err = callback()
		if err == nil {
			return nil
		}
	}
	return err
}

func imageFormat(id int32, page uint8, url string) string {
	var imageName string = path.Base(url)
	return fmt.Sprintf("%s/id-%d-%d-%s", Folder, id, page, imageName)
}

func DownloadAndWriteFile(filename string, fileUrl string, wg *sync.WaitGroup) error {
	defer wg.Done()
	var resp *http.Response

	err := retry( func() error {
		var err error
		resp, err = http.Get(fileUrl)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			log.Printf("Image: %d %s", resp.StatusCode, filename)
			return errors.New("http Error")
		}
		return err
	})
	if err != nil {
		return err
	}

	out, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	log.Printf("download and saved : %s", filename)
	return err
}

func readApiUrl(url string) (domain.HomeVision, error) {
	var target *domain.HomeVision = new(domain.HomeVision)
	resp, err := http.Get(url)

	if err != nil {
		log.Println(err)
		return *target, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		// log.Printf("API: %d", resp.StatusCode)
		return *target, errors.New("http Error")
	}
	log.Printf("StatusCode: %d, URL: %s", resp.StatusCode, url)
	json.NewDecoder(resp.Body).Decode(target)
	return *target, nil
}

func FetchHouseImage(url string, page uint8, images chan []ImageInfo) error {
	var houses domain.HomeVision
	err := retry(func() error {
		var readErr error
		houses, readErr = readApiUrl(url)
		return readErr
	})
	if err != nil {
		return err
	}

	var imagesInfo []ImageInfo = make([]ImageInfo, len(houses.Houses))
	for index, house := range houses.Houses {
		fileUrl := house.PhotoURL
		filename := imageFormat(house.Id, page, fileUrl)
		imagesInfo[index] = ImageInfo{fileUrl, filename}
	}
	images <- imagesInfo
	return nil
}
