package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"image"
	"image/png"
	"os"

	"unsafe"
)

type LevelData struct {
  Name   string  `json:"name"`                 
  Width  int     `json:"width"`                
  Height int     `json:"height"`               
	Levelmap    [][]int	`json:"map"`
                                               
  AmbientLight float64 `json:"ambientLight"`   
                                               
  PlayerCoordData                              
}                                              
                                               
type PlayerCoordData struct {                  
  PlayerX     float64 `json:"playerX"`         
  PlayerY     float64 `json:"playerY"`         
  PlayerAngle float64 `json:"playerAngle"`     
}                                              

var (
	filename *string
)

func parseFlags() {
	filename = flag.String("file", "", "Filename to generate level from.")

	flag.Parse()
}

func main() {
	parseFlags()

	if filename == nil {
		fmt.Println("Missing file flag.")
		os.Exit(1)
	}

	f, _ := os.Open(*filename)
	defer f.Close()

	img, _ := png.Decode(f)
	bounds := img.Bounds()

	width := bounds.Max.X - bounds.Min.X
	height := bounds.Max.Y - bounds.Min.Y

	levelData := LevelData{
		Name: "whatever",
		Width: width,
		Height: height,
		Levelmap: make([][]int, height),
		AmbientLight: 1.0,
	}
	pixelData := img.(*image.RGBA).Pix

	for y := 0; y < height; y++ {
		dataRow := make([]int, width)

		for x := 0; x < width; x++ {
			data := 0
			index := (x + y*width) << 2
			pixel := (*uint32)(unsafe.Pointer(&pixelData[index]))

			if checkWhite(*pixel) {
				data = 1
			}
			if checkGreen(*pixel) {
				levelData.PlayerCoordData = PlayerCoordData{
					PlayerX: float64(x),
					PlayerY: float64(y),
					PlayerAngle: 0.0,
				}
			}
			dataRow[x] = data
		}

		levelData.Levelmap[y] = dataRow
	}


	levelString, _ := json.Marshal(levelData)
	fmt.Print(string(levelString))
}

func checkWhite(pixel uint32) bool {
	if (pixel & 0xFFFFFF) == 0xFFFFFF {
		return true
	}
	return false
}

func checkBlack(pixel uint32) bool {
	if (pixel & 0xFFFFFF) == 0x000000 {
		return true
	}
	return false
}

func checkGreen(pixel uint32) bool {
	if (pixel & 0xFFFFFF) == 0x00FF00 {
		return true
	}
	return false
}
