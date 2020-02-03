package main

import "github.com/hajimehoshi/ebiten"

type Component interface {
	ComponentName() string
}

type Renderable struct {
	image *ebiten.Image
	x int
	y int
}

func (Renderable) ComponentName() string {
	return "Renderable"
}
