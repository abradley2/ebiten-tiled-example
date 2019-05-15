package main

import (
	"image/png"
	"math"
	"os"

	"github.com/hajimehoshi/ebiten"
	"github.com/lafriks/go-tiled"
	"github.com/lafriks/go-tiled/render"
)

const (
	numImages  = 200
	gameWidth  = 600
	gameHeight = 600
	gameScale  = 1
)

var op = &ebiten.DrawImageOptions{}

type imageEntity struct {
	xPos     float64
	yPos     float64
	rotation float64
	image    *ebiten.Image
}

var imageEntities [numImages]*imageEntity
var mapImage *ebiten.Image

func update(screen *ebiten.Image) error {
	if ebiten.IsRunningSlowly() {
		return nil
	}

	op.GeoM.Reset()
	op.GeoM.Translate(0, 0)
	screen.DrawImage(mapImage, op)

	for i := 0; i < numImages; i++ {
		ie := imageEntities[i]

		if ie != nil && ie.image != nil {
			op.GeoM.Reset()
			op.GeoM.Translate(ie.xPos, ie.yPos)
			screen.DrawImage(ie.image, op)
		}
	}

	return nil
}

func main() {
	gameMap, err := tiled.LoadFromFile("./area_1.tmx")

	if err != nil {
		panic(err)
	}

	mapRenderer, err := render.NewRenderer(gameMap)

	if err != nil {
		panic(err)
	}

	err = mapRenderer.RenderVisibleLayers()

	if err != nil {
		panic(err)
	}

	img, err := os.Create("area_1.png")

	if err != nil {
		panic(err)
	}

	mapRenderer.SaveAsPng(img)

	if err != nil {
		panic(err)
	}

	imFile, _ := os.Open("area_1.png")
	im, err := png.Decode(imFile)

	if err != nil {
		panic(err)
	}

	mapImage, _ = ebiten.NewImageFromImage(im, ebiten.FilterDefault)

	ebiten.Run(update, gameWidth, gameHeight, gameScale, "Kewl game")
}

func loadMap() {
	// load up the map as a tiled object
	gameMap, err := tiled.LoadFromFile("./area_1.tmx")

	if err != nil {
		panic(err)
	}

	// create a renderer
	mapRenderer, err := render.NewRenderer(gameMap)

	if err != nil {
		panic(err)
	}

	// render it to an in memory image
	err = mapRenderer.RenderVisibleLayers()

	if err != nil {
		panic(err)
	}

	// create a tmp file to write to then read from later
	img, err := os.Create("area_1.png")

	if err != nil {
		panic(err)
	}

	// save to the tmp file
	mapRenderer.SaveAsPng(img)
	img.Close()

	if err != nil {
		panic(err)
	}

	imFile, _ := os.Open("area_1.png")
	im, err := png.Decode(imFile)

	if err != nil {
		panic(err)
	}

	mapImage, _ = ebiten.NewImageFromImage(im, ebiten.FilterDefault)
}

func toRads(deg int) float64 {
	return math.Pi / 180 * float64(deg)
}
