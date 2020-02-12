package main

import (
	"github.com/hajimehoshi/ebiten"
)

type ComponentType int

const (
	RenderableType   ComponentType = iota
	PlayerType       ComponentType = iota
	PlayerBulletType ComponentType = iota
	BaddieType       ComponentType = iota
	HitBoxType       ComponentType = iota
)

// Base
type Component interface {
	ComponentType() ComponentType
}

// Renderable
type Renderable struct {
	Image    *ebiten.Image
	Location Point
}

func (*Renderable) ComponentType() ComponentType {
	return RenderableType
}

func NewRenderable(image *ebiten.Image, x int, y int) *Renderable {
	return &Renderable{image, Point{x, y}}
}

// Player
type Player struct{}

func (*Player) ComponentType() ComponentType {
	return PlayerType
}

func NewPlayer() *Player {
	return &Player{}
}

// PlayerBullet
type PlayerBullet struct{}

func (*PlayerBullet) ComponentType() ComponentType {
	return PlayerBulletType
}

func NewPlayerBullet() *PlayerBullet {
	return &PlayerBullet{}
}

// Baddie
type Baddie struct{}

func (*Baddie) ComponentType() ComponentType {
	return BaddieType
}

func NewBaddie() *Baddie {
	return &Baddie{}
}

// HitBox
type HitBox struct {
	TopLeft     Point
	TopRight    Point
	BottomRight Point
	BottomLeft  Point
}

func (*HitBox) ComponentType() ComponentType {
	return HitBoxType
}

//
//func NewHitBoxAbsolute(topLeft int, topRight int, bottomRight int, bottomLeft int) *HitBox {
//	return &HitBox{topLeft, topRight, bottomRight, bottomLeft}
//}
//
//func NewHitBoxRelative(width int, height int) *HitBox {
//	return &HitBox(0, width, )
//}
