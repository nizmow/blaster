package ecs

// ComponentType should be created for each component, so it can be used for querying for
// components in place of reflection.
type ComponentType int

// Component represents a component in the world, which will be attached to an entity.
// Component must return a particular ComponentType, which will be used to query the
// world for components of a particular type (instead of doing something like using
// reflection).
type Component interface {
	ComponentType() ComponentType
}
