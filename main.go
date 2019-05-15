package main

import (
	"bytes"
	"image/png"
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/lafriks/go-tiled"
	"github.com/lafriks/go-tiled/render"
)

const (
	numImages  = 200
	gameWidth  = 512
	gameHeight = 512
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
	loadMap()

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

	var buff []byte
	buffer := bytes.NewBuffer(buff)

	mapRenderer.SaveAsPng(buffer)

	im, err := png.Decode(buffer)

	mapImage, _ = ebiten.NewImageFromImage(im, ebiten.FilterDefault)
}

func toRads(deg int) float64 {
	return math.Pi / 180 * float64(deg)
}
