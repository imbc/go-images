package main

import (
	"image"
	"image/png"
	"log"
	"os"
)

func main() {
	width, height := 128, 128
	m := image.NewRGBA(image.Rect(0, 0, width, height))
	dirName := "img"
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		os.Mkdir(dirName, 0755)
		return
	}
	outFilename := dirName + "/canvas.png"
	outFile, err := os.Create(outFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()
	log.Print("Saving image to: ", outFilename)
	png.Encode(outFile, m)
}
