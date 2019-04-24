package main

import (
	"fmt"
	"image/jpeg"
	"io/ioutil"
	"log"
	"os"

	"github.com/nfnt/resize"
)

func main() {
	folderPath := "img/"
	//read Files from img folder
	files, err := ioutil.ReadDir(folderPath)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {

		file, err := os.Open(fmt.Sprintf("%s%s", folderPath, f.Name())) //img/img.jpg
		if err != nil {
			log.Fatal(err)
		}
		// decode jpeg into image.Image
		img, err := jpeg.Decode(file)
		if err != nil {
			log.Fatal(err)
		}
		file.Close()

		// resize to width 1000 using Lanczos resampling
		// and preserve aspect ratio
		m := resize.Resize(1000, 0, img, resize.Lanczos3)

		out, err := os.Create(fmt.Sprintf("%s%s", "result/", f.Name()))
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()

		// write new image to file
		jpeg.Encode(out, m, nil)
	}
	// open "test.jpg"
}
