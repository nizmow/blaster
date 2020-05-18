package blaster

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/nizmow/blaster/internal/ecs"
)

const (
	RenderableType   ecs.ComponentType = iota
	PlayerType       ecs.ComponentType = iota
	PlayerBulletType ecs.ComponentType = iota
	BaddieType       ecs.ComponentType = iota
	HitBoxType       ecs.ComponentType = iota
)

// Renderable
type Renderable struct {
	Image    *ebiten.Image
	Location ecs.Point
}

func (*Renderable) ComponentType() ecs.ComponentType {
	return RenderableType
}

func NewRenderable(image *ebiten.Image, x int, y int) *Renderable {
	return &Renderable{image, ecs.Point{x, y}}
}

// Player
type Player struct{}

func (*Player) ComponentType() ecs.ComponentType {
	return PlayerType
}

func NewPlayer() *Player {
	return &Player{}
}

// PlayerBullet
type PlayerBullet struct{}

func (*PlayerBullet) ComponentType() ecs.ComponentType {
	return PlayerBulletType
}

func NewPlayerBullet() *PlayerBullet {
	return &PlayerBullet{}
}

// Baddie
type Baddie struct{}

func (*Baddie) ComponentType() ecs.ComponentType {
	return BaddieType
}

func NewBaddie() *Baddie {
	return &Baddie{}
}

// HitBox
type HitBox struct {
	TopLeft     ecs.Point
	TopRight    ecs.Point
	BottomRight ecs.Point
	BottomLeft  ecs.Point
}

func (*HitBox) ComponentType() ecs.ComponentType {
	return HitBoxType
}
