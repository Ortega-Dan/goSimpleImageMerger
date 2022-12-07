package main

import (
	"image/png"
	"os"

	gim "github.com/ozankasikci/go-image-merge"
)

func main() {

	// accepts *Grid instances, grid unit count x, grid unit count y
	// returns an *image.RGBA object
	grids := []*gim.Grid{
		{ImageFilePath: os.Args[1]},
		{ImageFilePath: os.Args[2]},
	}

	rgba, err := gim.New(grids, 1, 2).Merge()
	if err != nil {
		panic(err)
	}

	// save the output to jpg or png
	file, err := os.Create("Joined.png")
	if err != nil {
		panic(err)
	}

	// err = jpeg.Encode(file, rgba, &jpeg.Options{Quality: 80})
	err = png.Encode(file, rgba)
	if err != nil {
		panic(err)
	}
}
