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
	playerEntity := NewEntity("Player")
	playerImage, _ := ebiten.NewImage(16, 16, ebiten.FilterNearest)
	playerImage.Fill(color.White)
	playerEntity.AddComponent(NewRenderable(playerImage, (ScreenWidth-16)/2, 200))
	playerEntity.AddComponent(NewPlayer())
	world.AddEntity(*playerEntity)

	baddieEntity := NewEntity("Baddie")
	baddieImage, _ := ebiten.NewImage(16, 16, ebiten.FilterNearest)
	baddieImage.Fill(color.RGBA{
		R: 255,
		G: 0,
		B: 0,
		A: 255,
	})
	baddieEntity.AddComponent(NewRenderable(baddieImage, (ScreenWidth-16)/2, 20))
	baddieEntity.AddComponent(NewBaddie())
	world.AddEntity(*baddieEntity)

	// systems
	renderer = Renderer{}
	playerInput = PlayerInput{}
	playerBulletMover = PlayerBulletMover{}

	if err := ebiten.Run(update, ScreenWidth, ScreenHeight, 2, "Hello World"); err != nil {
		log.Fatal(err)
	}
}
