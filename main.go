package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
)

const ScreenWidth = 320
const ScreenHeight = 240

type World struct {
	entities []Entity
}

func (world *World) AddEntity(e Entity) *World {
	world.entities = append(world.entities, e)
	return world
}

func (world *World) RemoveEntity(idToRemove int) {
	for i, e := range world.entities {
		if e.ID == idToRemove {
			world.entities = append(world.entities[:i], world.entities[i+1:]...)
			break
		}
	}
}

func (world *World) GetEntities() []Entity {
	return world.entities
}

var world World

var renderer Renderer
var playerInput PlayerInput
var playerBulletMover PlayerBulletMover

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	var err error

	err = screen.Fill(color.Black)
	if err != nil {
		return err
	}
	err = renderer.Update(world, screen)
	if err != nil {
		return err
	}
	err = playerInput.Update(&world)
	if err != nil {
		return err
	}
	err = playerBulletMover.Update(&world)
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
	world.AddEntity(*NewEntity("Player").AddComponent(NewRenderable(player, (ScreenWidth-16)/2, 200)).AddComponent(NewPlayer()))

	// systems
	renderer = Renderer{}
	playerInput = PlayerInput{}
	playerBulletMover = PlayerBulletMover{}

	if err := ebiten.Run(update, ScreenWidth, ScreenHeight, 2, "Hello World"); err != nil {
		log.Fatal(err)
	}
}
