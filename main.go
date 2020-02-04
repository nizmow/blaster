package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
)

type World struct {
	entities []Entity
}

func (w *World) AddEntity(e Entity) *World {
	w.entities = append(w.entities, e)
	return w
}
func (w *World) GetEntities() []Entity {
	return w.entities
}

var world World

var renderer Renderer

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	var err error

	err = screen.Fill(color.NRGBA{0xff, 0x00, 0x00, 0xff})
	if err != nil {
		return err
	}
	err = renderer.Update(world, screen)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	// setup

	// player entity
	player, _ := ebiten.NewImage(16, 16, ebiten.FilterNearest)
	player.Fill(color.White)
	world.AddEntity(*NewEntity(1, "Player").AddComponent(NewRenderable(player, 100, 100)))

	// systems
	renderer = Renderer{}

	if err := ebiten.Run(update, 320, 240, 2, "Hello World"); err != nil {
		log.Fatal(err)
	}
}
