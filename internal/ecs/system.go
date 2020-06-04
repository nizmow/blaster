package ecs

import "github.com/hajimehoshi/ebiten"

// LogicSystem is a system which applies logic. These are run before render systems.
type LogicSystem interface {
	Update(*World) error
}

// RenderSystem is a system which performs some kind of screen rendering. These are run last.
type RenderSystem interface {
	Update(*World, *ebiten.Image) error
}
