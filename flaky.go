package main

import (
	"homeVision/application"
	"homeVision/infrastructure"
	"log"
	"os"
)

func main() {
	// Create the images folder if does not exist
	if _, err := os.Stat(application.Folder); os.IsNotExist(err) {
		err := os.Mkdir(application.Folder, 0755)
		if err != nil {
			log.Fatalln(err)
		}
	}
	infrastructure.HomeVision()
}
