package blaster

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/nizmow/blaster/internal/ecs"
)

// ScreenWidth is the width of the screen
const ScreenWidth = 320

// ScreenHeight is the heigh of the screen
const ScreenHeight = 240

var world ecs.World

var renderer Renderer
var playerInput PlayerInput
var playerBulletMover PlayerBulletMover
var bulletBaddieCollision BulletBaddieCollision
var baddieMoverSystem baddieMover

// Run begins the game loop.
func Run() {
	// setup

	// player entity
	playerEntity := ecs.NewEntity("Player")
	playerImage, _ := ebiten.NewImage(16, 16, ebiten.FilterNearest)
	playerImage.Fill(color.White)
	playerEntity.AddComponent(NewRenderable(playerImage, (ScreenWidth-16)/2, 200))
	playerEntity.AddComponent(NewPlayer())

	world.AddEntity(*playerEntity)

	baddieGroup := newBaddieGroup()
	baddieGroup.direction = 1
	for i := 0; i < 30; i++ {
		y := 20 + 32*(i/10)
		x := 8 + 32*(i%10)
		baddieEntity := ecs.NewEntity(fmt.Sprintf("baddie-%d", i))
		baddieImage, _ := ebiten.NewImage(16, 16, ebiten.FilterNearest)
		baddieImage.Fill(color.RGBA{
			R: 255,
			G: 0,
			B: 0,
			A: 255,
		})
		baddieEntity.AddComponent(NewRenderable(baddieImage, x, y))
		baddieEntity.AddComponent(NewBaddie())
		baddieEntity.AddComponent(baddieGroup)

		world.AddEntity(*baddieEntity)
	}

	// systems
	renderer = Renderer{}
	playerInput = PlayerInput{}
	playerBulletMover = PlayerBulletMover{}
	bulletBaddieCollision = BulletBaddieCollision{}
	baddieMoverSystem = baddieMover{}

	if err := ebiten.Run(update, ScreenWidth, ScreenHeight, 4, "Hello World"); err != nil {
		log.Fatal(err)
	}
}

func update(screen *ebiten.Image) error {
	var err error

	err = playerInput.Update(&world)

	err = playerBulletMover.Update(&world)

	err = bulletBaddieCollision.Update(&world)

	err = baddieMoverSystem.update(&world)

	// RENDER COMMANDS COME AFTER THIS!
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	err = screen.Fill(color.Black)

	err = renderer.Update(world, screen)

	if err != nil {
		return err
	}

	return nil
}
