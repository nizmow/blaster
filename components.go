package main

import (
	"github.com/hajimehoshi/ebiten"
)

// Base
type Component interface {
	ComponentName() string
}

// Renderable
type Renderable struct {
	Image *ebiten.Image
	X     int
	Y     int
}

func (Renderable) ComponentName() string {
	return "Renderable"
}
