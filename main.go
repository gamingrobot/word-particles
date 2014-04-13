package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

func main() {
	if len(os.Args) < 1 {
		fmt.Println("Input file is missing.")
		os.Exit(1)
	}
	img, _, err := decode(os.Args[1])
	if err != nil {
		return
	}
	count := 0
	b := img.Bounds()
	fmt.Print("var word_pixels = [")
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			if matchColor(img.At(x, y), 0) {
				fmt.Printf("[%d,%d],", x, y)
				count++
			}
		}
	}
	fmt.Println("];")
	fmt.Printf("var pixel_count = %d;\n", count)
	fmt.Printf("var image_width = %d;\n", b.Max.X)
	fmt.Printf("var image_height = %d;\n", b.Max.Y)
}

func decode(filename string) (image.Image, string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, "", err
	}
	defer f.Close()
	return image.Decode(bufio.NewReader(f))
}

func matchColor(cl color.Color, match uint32) bool {
	r, g, b, _ := cl.RGBA()
	return (r == match && g == match && b == match)
}
