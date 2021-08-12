package infrastructure

import (
	"fmt"
	"homeVision/application"
	"io/ioutil"
	"log"
	"sync"
)

const (
	API     string = "https://app-homevision-staging.herokuapp.com/api_project/houses?page=%d"
	pages   uint8  = 10
)

func HomeVision () {
	imagesInfo := make(chan []application.ImageInfo, int(pages))
	var imagesExpected int

	for i := uint8(1); i <= pages; i++ {
		var api string = fmt.Sprintf(API, i)
		go application.FetchHouseImage(api, i, imagesInfo)
	}

	var wg1 sync.WaitGroup
	var counter int = 1
	for images := range imagesInfo {
		var lenImages int = len(images)
		imagesExpected += lenImages
		wg1.Add(lenImages)
		if counter == int(pages) {
			close(imagesInfo)
		}

		for _, image := range images {
			go application.DownloadAndWriteFile(image.Filename, image.Url, &wg1)
		}
		counter++
	}

	log.Println("Almost Done!")
	wg1.Wait()
	log.Println("Done!")
	fmt.Println("--------------------")
	files, _ := ioutil.ReadDir(application.Folder)
	log.Printf("images expected '%d' got '%d'", imagesExpected, len(files))
}
