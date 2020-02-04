package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var square *ebiten.Image

type World struct {
	Entities []Entity
}

var world World

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	screen.Fill(color.NRGBA{0xff, 0x00, 0x00, 0xff})

	ebitenutil.DebugPrint(screen, "Hello world")

	if square == nil {
		square, _ = ebiten.NewImage(16, 16, ebiten.FilterNearest)
	}

	square.Fill(color.White)

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(64, 64)
	screen.DrawImage(square, opts)

	return nil
}

func main() {
	// setup

	// player entity
	player, _ := ebiten.NewImage(16, 16, ebiten.FilterNearest)
	player.Fill(color.White)
	world.Entities = append(world.Entities, Entity{
		1,
		"Player",
		[]Component{Renderable{player, 100, 100}},
	})

	if err := ebiten.Run(update, 640, 480, 2, "Hello World"); err != nil {
		log.Fatal(err)
	}
}
