package blaster

import (
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

type PlayerInput struct{}

func (PlayerInput) Update(world *ecs.World) error {
	player := world.FindComponentsJoin(RenderableType, PlayerType)[0]

	playerRenderable := player.RequestedComponents[RenderableType].(*Renderable)

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
		playerBulletComponents := world.FindComponents(PlayerBulletType)
		if len(playerBulletComponents) <= MaxPlayerBullets {
			bulletEntity := ecs.NewEntity("Player Bullet")

			bulletEntity.AddComponent(NewPlayerBullet())

			image, _ := ebiten.NewImage(2, 2, ebiten.FilterNearest)
			image.Fill(color.White)
			bulletEntity.AddComponent(NewRenderable(image, playerRenderable.Location.X+7, playerRenderable.Location.Y))

			world.AddEntity(*bulletEntity)
		}
	}

	return nil
}

type PlayerBulletMover struct{}

func (PlayerBulletMover) Update(world *ecs.World) error {
	allPlayerBulletComponents := world.FindComponentsJoin(RenderableType, PlayerBulletType)

	for _, playerBulletComponents := range allPlayerBulletComponents {
		playerBulletRenderable := playerBulletComponents.RequestedComponents[RenderableType].(*Renderable)
		if playerBulletRenderable.Location.Y > 0 {
			playerBulletRenderable.Location.Y -= 5
		} else {
			world.RemoveEntity(playerBulletComponents.Entity.ID)
		}
	}

	return nil
}

//type PlayerBulletBaddieCollider struct{}
//
//func (PlayerBulletBaddieCollider) Update(world *World) error {
//	allPlayerBullets := world.FindComponentsJoin(RenderableType, PlayerBulletType)
//	allBaddies := world.FindComponentsJoin(RenderableType, BaddieType)
//
//	for _, playerBullet := range allPlayerBullets {
//		playerBulletRenderable := playerBullet.RequestedComponents[RenderableType].(*Renderable)
//
//		for _, baddie := range allBaddies {
//
//		}
//	}
//}
