package main

import (
	"image/png"
	"math"
	"os"
	"path"

	"github.com/hajimehoshi/ebiten"
)

const (
	gameHeight = 800
	gameWidth  = 800
)

type body struct {
	xPos     float64
	yPos     float64
	active   bool
	rotation int
	dir      string
	yMove    int
	face     bool
}

const (
	movSpeed = 4
	down     = "down"
	left     = "left"
	right    = "right"
)

func (b *body) update() {
	if b.dir == right {
		b.xPos = b.xPos + movSpeed
		if b.xPos > gameWidth-64 {
			b.dir = down
		}
	}
	if b.dir == left {
		b.xPos = b.xPos - movSpeed
		if b.xPos < 1 {
			b.dir = down
		}
	}
	if b.dir == down {
		b.yPos = b.yPos + movSpeed
		b.yMove = b.yMove + movSpeed
		if b.yMove == 64 {
			b.yMove = 0
			if b.xPos > 100 {
				b.dir = left
			} else {
				b.dir = right
			}
		}
	}
}

type centipede struct {
	sections [9]*body
}

var pede *centipede

var op = &ebiten.DrawImageOptions{}

var faceImg *ebiten.Image
var bodyImg *ebiten.Image

var x float64
var y float64
var rot int

func update(screen *ebiten.Image) error {
	if ebiten.IsRunningSlowly() {
		return nil
	}

	for i := 0; i < len(pede.sections); i++ {
		s := pede.sections[i]
		op.GeoM.Reset()
		s.update()
		op.GeoM.Translate(s.xPos, s.yPos)
		if s.face == true {
			screen.DrawImage(faceImg, op)
		} else {
			screen.DrawImage(faceImg, op)
		}
	}

	return nil
}

func main() {

	loadSpacey()

	loadPede()

	ebiten.Run(update, gameWidth, gameHeight, 1, "Ebiten Starter")
}

func loadPede() {
	pede = new(centipede)
	for i := 0; i < len(pede.sections); i++ {
		pede.sections[i] = &body{
			xPos:     float64(0 + (-64 * i)),
			yPos:     0,
			rotation: 0,
			dir:      "right",
			face:     i == 0,
		}
	}
}

func toRads(deg int) float64 {
	return math.Pi / 180 * float64(deg)
}

func loadBody() {
	wd, _ := os.Getwd()
	f, _ := os.Open(path.Join(wd, "./assets/body.png"))
	i, _ := png.Decode(f)

	bi, _ := ebiten.NewImageFromImage(i, ebiten.FilterDefault)

	bodyImg = bi
}

func loadSpacey() {
	wd, _ := os.Getwd()
	f, _ := os.Open(path.Join(wd, "./assets/space_face.png"))
	i, _ := png.Decode(f)

	fi, _ := ebiten.NewImageFromImage(i, ebiten.FilterDefault)

	faceImg = fi
}
