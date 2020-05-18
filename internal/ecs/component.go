package ecs

type ComponentType int

type Component interface {
	ComponentType() ComponentType
}
