package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func main() {
	width, height := 128, 128
	canvas := NewCanvas(image.Rect(0, 0, width, height))
	canvas.DrawGradient()

	// Draw a series of lines from the top left corner to the bottom of the image
	for x := 0; x < width; x += 8 {
		canvas.DrawLine(color.RGBA{0, 0, 0, 255},
			Vector{0.0, 0.0},
			Vector{float64(x), float64(height)})
	}

	dirName := "img"
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		os.Mkdir(dirName, 0755)
		return
	}
	outFilename := dirName + "/lines.png"
	outFile, err := os.Create(outFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()
	log.Print("Saving image to: ", outFilename)
	png.Encode(outFile, canvas)
}
