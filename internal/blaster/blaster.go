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

var renderer rendererSystem
var playerInput playerInputSystem
var playerBulletMover playerBulletMoverSystem
var bulletBaddieCollision bulletBaddieCollisionSystem
var baddieMover baddieMoverSystem

type blaster struct {
	world ecs.World
}

// Update performs system logic every tick.
func (g *blaster) Update(screen *ebiten.Image) error {
	var err error

	err = playerInput.update(&g.world)

	err = playerBulletMover.update(&g.world)

	err = bulletBaddieCollision.update(&g.world)

	err = baddieMover.update(&g.world)

	if err != nil {
		return err
	}

	return nil
}

func (g *blaster) Draw(screen *ebiten.Image) {

	screen.Fill(color.Black)

	renderer.update(g.world, screen)
}

func (g *blaster) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

// Run begins the game loop.
func Run() {
	// setup
	world := ecs.World{}

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
	renderer = rendererSystem{}
	playerInput = playerInputSystem{}
	playerBulletMover = playerBulletMoverSystem{}
	bulletBaddieCollision = bulletBaddieCollisionSystem{}
	baddieMover = baddieMoverSystem{}

	if err := ebiten.RunGame(&blaster{world}); err != nil {
		log.Fatal(err)
	}
}
