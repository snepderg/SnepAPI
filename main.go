// SnepAPI is a simple REST API that retrieves images of snow leopards.
// It is written in Go and will use a router such as Gorilla Mux.
// Right now, the focus is on the local operation of the API.
// In the future, it will be deployed to a cloud service.

// Note that this is my first time using Go, so I'm still learning the basics.

package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
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

	// Fetch a random image from the directory
	fmt.Println("Fetching a random image...")
	rand.Seed(time.Now().Unix())
	randImage := files[rand.Intn(fileCount)]
	fmt.Println("Image retrieved:", randImage.Name())

	fmt.Println("Would you like to see the image? (y/n)")
	var input string
	fmt.Scanln(&input)
	if input == "y" {
		fmt.Println("Opening image...")
	} else {
		fmt.Println("Okay, goodbye!")
		os.Exit(0)
	}

	// Open the image (default program for the OS)
	// TODO: Find a package to handle opening images in a cross-platform way
	cmd := exec.Command("cmd", "/C", "start", "./images/"+randImage.Name())
	err = cmd.Start()
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	} else {
		fmt.Println("Alright! Your image should appear shortly.")
	}

	time.Sleep(1)
	fmt.Println("Goodbye!")
}
