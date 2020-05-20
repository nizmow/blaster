package blaster

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/nizmow/blaster/internal/ecs"
)

const PlayerSpeed = 2
const MaxPlayerBullets = 5

type Renderer struct{}

func (Renderer) Update(world ecs.World, screen *ebiten.Image) error {
	renderables := world.FindComponents(RenderableType)
	for _, renderCandidate := range renderables {
		render := renderCandidate.RequestedComponent.(*Renderable)
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(float64(render.Location.X), float64(render.Location.Y))
		err := screen.DrawImage(render.Image, opts)
		if err != nil {
			return err
		}
	}
	return nil
}

type PlayerInput struct {
	gunHeat int
}

func (playerInput *PlayerInput) Update(world *ecs.World) error {
	// We'd better not have multiple players, but if so, ignore them.
	playerTypeResult := world.FindComponents(PlayerType)[0]

	playerRenderable := playerTypeResult.Entity.GetComponent(RenderableType).(*Renderable)

	switch {
	case ebiten.IsKeyPressed(ebiten.KeyLeft):
		if playerRenderable.Location.X > 2 {
			playerRenderable.Location.X -= PlayerSpeed
		}
	case ebiten.IsKeyPressed(ebiten.KeyRight):
		if playerRenderable.Location.X < ScreenWidth-18 {
			playerRenderable.Location.X += PlayerSpeed
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		// fire a bullet, but only if our gun isn't too hot!
		if playerInput.gunHeat < 200 {
			playerInput.gunHeat = playerInput.gunHeat + 50
			fmt.Printf("Player firing, heat is now %d\n", playerInput.gunHeat)
			bulletEntity := ecs.NewEntity("Player Bullet")

			bulletEntity.AddComponent(NewPlayerBullet())

			image, _ := ebiten.NewImage(2, 2, ebiten.FilterNearest)
			image.Fill(color.White)
			bulletEntity.AddComponent(NewRenderable(image, playerRenderable.Location.X+7, playerRenderable.Location.Y))

			world.AddEntity(*bulletEntity)
		}
	}

	// Cool down the gun.
	if playerInput.gunHeat > 0 {
		playerInput.gunHeat = playerInput.gunHeat - 1
		if playerInput.gunHeat%10 == 0 {
			fmt.Printf("Gun cooling, heat is now %d\n", playerInput.gunHeat)
		}
	}

	return nil
}

type PlayerBulletMover struct{}

func (PlayerBulletMover) Update(world *ecs.World) error {
	allPlayerBulletComponents := world.FindComponents(PlayerBulletType)

	for _, playerBulletComponents := range allPlayerBulletComponents {
		playerBulletRenderable := playerBulletComponents.Entity.GetComponent(RenderableType).(*Renderable)
		if playerBulletRenderable.Location.Y > 0 {
			playerBulletRenderable.Location.Y -= 5
		} else {
			world.RemoveEntity(playerBulletComponents.Entity.ID)
		}
	}

	return nil
}

// BulletBaddieCollision contains logic to test for collisions between bullets and baddies,
// and perform appropriate actions (no more baddies).
type BulletBaddieCollision struct{}

// Update runs the collision detection system.
func (BulletBaddieCollision) Update(world *ecs.World) error {
	baddieComponentsResult := world.FindComponents(BaddieType)
	playerBulletComponentsResult := world.FindComponents(PlayerBulletType)

	for _, baddie := range baddieComponentsResult {
		baddieRenderable := baddie.Entity.GetComponent(RenderableType).(*Renderable)
		if baddieRenderable == nil {
			continue
		}
		for _, playerBullet := range playerBulletComponentsResult {
			playerBulletRenderable := playerBullet.Entity.GetComponent(RenderableType).(*Renderable)
			if playerBulletRenderable == nil {
				break
			}

			if baddieRenderable.TranslateHitboxToScreen().Overlaps(playerBulletRenderable.TranslateHitboxToScreen()) {
				world.RemoveEntity(playerBullet.Entity.ID)
				world.RemoveEntity(baddie.Entity.ID)
			}
		}
	}

	return nil
}
