package blaster

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/nizmow/blaster/internal/ecs"
)

const playerSpeed = 2
const baddieSpeed = 1
const maxPlayerBullets = 3

var currentPlayerBullets int = 0

type rendererSystem struct{}

func (rendererSystem) update(world ecs.World, screen *ebiten.Image) error {
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

type playerInputSystem struct {
	gunHeat     int
	playerSpeed int
}

func (playerInputSystem *playerInputSystem) update(world *ecs.World) error {
	// We'd better not have multiple players, but if so, ignore them.
	playerTypeResult := world.FindComponents(PlayerType)[0]

	playerRenderable := playerTypeResult.Entity.GetComponent(RenderableType).(*Renderable)

	switch {
	case ebiten.IsKeyPressed(ebiten.KeyLeft):
		if playerRenderable.Location.X > 2 {
			playerRenderable.Location.X -= playerSpeed
		}
	case ebiten.IsKeyPressed(ebiten.KeyRight):
		if playerRenderable.Location.X < ScreenWidth-18 {
			playerRenderable.Location.X += playerSpeed
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		// fire a bullet, but only if our gun isn't too hot!
		if currentPlayerBullets < maxPlayerBullets {
			bulletEntity := ecs.NewEntity(fmt.Sprintf("player-bullet-%d", currentPlayerBullets))

			bulletEntity.AddComponent(NewPlayerBullet())

			image, _ := ebiten.NewImage(2, 2, ebiten.FilterNearest)
			image.Fill(color.White)
			bulletEntity.AddComponent(NewRenderable(image, playerRenderable.Location.X+7, playerRenderable.Location.Y))

			world.AddEntity(*bulletEntity)
			currentPlayerBullets++
		}
	}

	return nil
}

type playerBulletMoverSystem struct{}

func (playerBulletMoverSystem) update(world *ecs.World) error {
	allPlayerBulletComponents := world.FindComponents(PlayerBulletType)

	for _, playerBulletComponents := range allPlayerBulletComponents {
		playerBulletRenderable := playerBulletComponents.Entity.GetComponent(RenderableType).(*Renderable)
		if playerBulletRenderable.Location.Y > 0 {
			playerBulletRenderable.Location.Y -= 5
		} else {
			world.RemoveEntity(playerBulletComponents.Entity.ID)
			currentPlayerBullets--
		}
	}

	return nil
}

// BulletBaddieCollision contains logic to test for collisions between bullets and baddies,
// and perform appropriate actions (no more baddies).
type bulletBaddieCollisionSystem struct{}

// Update runs the collision detection system.
func (bulletBaddieCollisionSystem) update(world *ecs.World) error {
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
				currentPlayerBullets--
			}
		}
	}

	return nil
}

type baddieMoverSystem struct{}

func (baddieMoverSystem) update(world *ecs.World) error {
	baddieComponentsResult := world.FindComponents(BaddieType)
	// First we have to check if any baddies in any group will become out of bounds.
	// We can only movie baddies if we're sure we won't change the group movement
	// direction half way through -- otherwise they break formation!
	for _, baddie := range baddieComponentsResult {
		baddieRenderable := baddie.Entity.GetComponent(RenderableType).(*Renderable)
		baddieGroup := baddie.Entity.GetComponent(BaddieGroupType).(*baddieGroup)

		if baddieGroup.direction == 1 && baddieRenderable.Location.X >= ScreenWidth-18 {
			// Out of bounds on the right hand side, moving right.
			baddieGroup.direction = -1
		}

		if baddieGroup.direction == -1 && baddieRenderable.Location.X < 2 {
			// Out of bounds on the left hand side, moving left.
			baddieGroup.direction = 1
		}
	}

	// Now we just move our baddies in accodance with the direction of the group.
	for _, baddie := range baddieComponentsResult {
		baddieRenderable := baddie.Entity.GetComponent(RenderableType).(*Renderable)
		baddieGroup := baddie.Entity.GetComponent(BaddieGroupType).(*baddieGroup)

		if baddieGroup.direction == 1 {
			baddieRenderable.Location.X += baddieSpeed
		}

		if baddieGroup.direction == -1 {
			baddieRenderable.Location.X -= baddieSpeed
		}
	}

	return nil
}
