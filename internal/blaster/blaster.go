package blaster

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/nizmow/blaster/internal/ecs"
)

var renderer rendererSystem
var playerInput playerInputSystem
var playerBulletMover playerBulletMoverSystem
var bulletBaddieCollision bulletBaddieCollisionSystem
var baddieMover baddieMoverSystem

type blaster struct {
	world  ecs.World
	buffer *ebiten.Image
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

	g.buffer.Fill(color.Black)

	renderer.update(g.world, g.buffer)

	return nil
}

func (g *blaster) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.buffer, nil)
}

func (g *blaster) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.world.ScreenWidth, g.world.ScreenHeight
}

func (g *blaster) Init() {
	g.world = ecs.World{ScreenWidth: 320, ScreenHeight: 240}

	g.buffer, _ = ebiten.NewImage(g.world.ScreenWidth, g.world.ScreenHeight, ebiten.FilterDefault)

	// player entity
	playerEntity := ecs.NewEntity("Player")
	playerImage, _ := ebiten.NewImage(16, 16, ebiten.FilterNearest)
	playerImage.Fill(color.White)
	playerEntity.AddComponent(NewRenderable(playerImage, (g.world.ScreenWidth-16)/2, 200))
	playerEntity.AddComponent(NewPlayer())

	g.world.AddEntity(*playerEntity)

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

		g.world.AddEntity(*baddieEntity)
	}
}

// Run begins the game loop.
func Run() {
	// setup
	game := &blaster{}
	game.Init()

	// systems
	renderer = rendererSystem{}
	playerInput = playerInputSystem{}
	playerBulletMover = playerBulletMoverSystem{}
	bulletBaddieCollision = bulletBaddieCollisionSystem{}
	baddieMover = baddieMoverSystem{}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
