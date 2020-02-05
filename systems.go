package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"image/color"
)

const PlayerSpeed = 2
const MaxPlayerBullets = 5

type Renderer struct{}

func (Renderer) Update(world World, screen *ebiten.Image) error {
	renderables := world.FindComponents(RenderableType)
	for _, renderCandidate := range renderables {
		render := renderCandidate.RequestedComponent.(*Renderable)
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(float64(render.X), float64(render.Y))
		err := screen.DrawImage(render.Image, opts)
		if err != nil {
			return err
		}
	}
	return nil
}

type PlayerInput struct{}

func (PlayerInput) Update(world *World) error {
	player := world.FindComponentsJoin(RenderableType, PlayerType)[0]

	playerRenderable := player.RequestedComponents[RenderableType].(*Renderable)

	switch {
	case ebiten.IsKeyPressed(ebiten.KeyLeft):
		if playerRenderable.X > 2 {
			playerRenderable.X -= PlayerSpeed
		}
	case ebiten.IsKeyPressed(ebiten.KeyRight):
		if playerRenderable.X < ScreenWidth-18 {
			playerRenderable.X += PlayerSpeed
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		playerBulletComponents := world.FindComponents(PlayerBulletType)
		if len(playerBulletComponents) <= MaxPlayerBullets {
			bulletEntity := NewEntity("Player Bullet")

			bulletEntity.AddComponent(NewPlayerBullet())

			image, _ := ebiten.NewImage(2, 2, ebiten.FilterNearest)
			image.Fill(color.White)
			bulletEntity.AddComponent(NewRenderable(image, playerRenderable.X+7, playerRenderable.Y))

			world.AddEntity(*bulletEntity)
		}
	}

	return nil
}

type PlayerBulletMover struct{}

func (PlayerBulletMover) Update(world *World) error {
	allPlayerBulletComponents := world.FindComponentsJoin(RenderableType, PlayerBulletType)

	for _, playerBulletComponents := range allPlayerBulletComponents {
		playerBulletRenderable := playerBulletComponents.RequestedComponents[RenderableType].(*Renderable)
		if playerBulletRenderable.Y > 0 {
			playerBulletRenderable.Y -= 5
		} else {
			world.RemoveEntity(playerBulletComponents.Entity.ID)
		}
	}

	return nil
}
