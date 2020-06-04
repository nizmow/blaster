package blaster

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/nizmow/blaster/internal/ecs"
)

type blaster struct {
	world  ecs.World
	buffer *ebiten.Image
}

var ticks = 0

// Update performs system logic every tick (60 per second)
func (g *blaster) Update(screen *ebiten.Image) error {
	err := g.world.Tick(g.buffer)
	ticks++
	return err
}

func (g *blaster) Draw(screen *ebiten.Image) {
	_ = screen.DrawImage(g.buffer, nil)
}

func (g *blaster) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.world.ScreenWidth, g.world.ScreenHeight
}

func (g *blaster) Init() {
	//g.world = ecs.World{ScreenWidth: 320, ScreenHeight: 240}

	g.world = ecs.NewWorld(320, 240)

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
		baddieGroup.numberOfEntities++

		g.world.AddEntity(*baddieEntity)
	}

	// Systems will be executed in the order that they are added
	g.world.AddLogicSystem(&playerInputSystem{})
	g.world.AddLogicSystem(&playerBulletMoverSystem{})
	g.world.AddLogicSystem(&collisionSystem{})
	g.world.AddLogicSystem(&baddieMoverSystem{})

	g.world.AddRenderSystem(&rendererSystem{})

	g.world.AddEventHandler(&BaddieCollisionEventHandler{})
	g.world.AddEventHandler(&EntityRemovedEventHandler{})
}

// Run begins the game loop.
func Run() {
	// setup
	game := &blaster{}
	game.Init()

	ebiten.SetWindowSize(1024, 768)
	ebiten.SetWindowTitle("blaster by nizmow")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
