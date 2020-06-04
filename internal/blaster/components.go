package blaster

import (
	"image"

	"github.com/hajimehoshi/ebiten"
	"github.com/nizmow/blaster/internal/ecs"
)

const (
	RenderableType ecs.ComponentType = iota
	PlayerType
	PlayerBulletType
	BaddieType
	BaddieGroupType
)

// Renderable
type Renderable struct {
	Image    *ebiten.Image
	Location image.Point
	Hitbox   image.Rectangle
}

func (*Renderable) ComponentType() ecs.ComponentType {
	return RenderableType
}

func NewRenderable(renderImage *ebiten.Image, x int, y int) *Renderable {
	return &Renderable{renderImage, image.Point{x, y}, image.Rectangle{renderImage.Bounds().Min, renderImage.Bounds().Max}}
}

// TranslateHitboxToScreen translates a renderable's defined hitbox to real screen coordinates
// based on the current renderable point Location.
func (renderable *Renderable) TranslateHitboxToScreen() image.Rectangle {
	return renderable.Hitbox.Add(renderable.Location)
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

type baddieGroup struct {
	// direction of 1 is right, -1 is left, 0 is still
	direction        int
	ticksPerMove     int
	ticksUntilMove   int
	numberOfEntities int
}

func (*baddieGroup) ComponentType() ecs.ComponentType {
	return BaddieGroupType
}

func newBaddieGroup() *baddieGroup {
	return &baddieGroup{1, 30, 30, 0}
}
