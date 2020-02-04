package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
)

var square *ebiten.Image

type World struct {
	Entities []Entity
}

var world World

var renderer Renderer

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	screen.Fill(color.NRGBA{0xff, 0x00, 0x00, 0xff})

	renderer := NewRenderer(screen)
	renderer.Update(world)

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
		[]Component{&Renderable{player, 100, 100}},
	})

	if err := ebiten.Run(update, 640, 480, 2, "Hello World"); err != nil {
		log.Fatal(err)
	}
}
