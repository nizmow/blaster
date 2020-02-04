package main

import (
	"github.com/hajimehoshi/ebiten"
)

type ComponentType int

const (
	RenderableType   ComponentType = iota
	PlayerType       ComponentType = iota
	PlayerBulletType ComponentType = iota
)

// Base
type Component interface {
	ComponentType() ComponentType
}

// Renderable
type Renderable struct {
	Image *ebiten.Image
	X     int
	Y     int
}

func (*Renderable) ComponentType() ComponentType {
	return RenderableType
}

func NewRenderable(image *ebiten.Image, x int, y int) *Renderable {
	return &Renderable{image, x, y}
}

type Player struct{}

func (*Player) ComponentType() ComponentType {
	return PlayerType
}

func NewPlayer() *Player {
	return &Player{}
}

type PlayerBullet struct{}

func (*PlayerBullet) ComponentType() ComponentType {
	return PlayerBulletType
}

func NewPlayerBullet() *PlayerBullet {
	return &PlayerBullet{}
}
