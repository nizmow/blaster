package blaster

import "github.com/nizmow/blaster/internal/ecs"

const (
	BaddieRemovedEventType ecs.EventType = iota
	CollisionEventType
)

// BaddieRemovedEvent is fired whenever a baddie is removed for any reason.
type BaddieRemovedEvent struct {
	Entity ecs.Entity
}

func (*BaddieRemovedEvent) EventType() ecs.EventType {
	return BaddieRemovedEventType
}

// CollisionEvent is fired when two renderables collide.
type CollisionEvent struct {
	FirstEntity ecs.Entity
	OtherEntity ecs.Entity
}

func (*CollisionEvent) EventType() ecs.EventType {
	return CollisionEventType
}
