// SnepAPI is a simple REST API that retrieves images of snow leopards.
// It is written in Go and will use a router such as Gorilla Mux.
// Right now, the focus is on the local operation of the API.
// In the future, it will be deployed to a cloud service.

// Note that this is my first time using Go, so I'm still learning the basics.

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	fmt.Println("Hello, world!")
	fmt.Println("This will be a REST API that retrieves images of snow leopards.")
	fmt.Println("Right now, it will simply fetch images from a local directory.")

	time.Sleep(1)
	fmt.Println("Reticulating splines... (grabbing an image from the local directory)")

	// Get the ./images directory
	files, err := ioutil.ReadDir("./images")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	} else if len(files) == 0 {
		fmt.Println("Error: No images were found in the directory.\nPlease ensure the path is correct, or add some images to the ./images directory and try again.")
		os.Exit(1)
	}

	// Get the count of files in the directory
	fileCount := len(files)
	fmt.Println(fileCount, "images were found in the directory.")

	fmt.Println("Here are the images:")
	// List the file names in the directory
	for i := 0; i < fileCount; i++ {
		fmt.Println(files[i].Name())
	}

	time.Sleep(1)
	fmt.Println("Goodbye!")
}
