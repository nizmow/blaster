package main

import (
	"github.com/hajimehoshi/ebiten"
)

type ComponentType int

const (
	RenderableType ComponentType = 0
	PlayerType     ComponentType = 1
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

func (Renderable) ComponentType() ComponentType {
	return RenderableType
}

func NewRenderable(image *ebiten.Image, x int, y int) *Renderable {
	return &Renderable{image, x, y}
}
