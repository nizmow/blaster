package ecs

const (
	EntityRemovedEventType EventType = iota + 100000
	EntityAddedEventType
)

type EntityRemovedEvent struct {
	Entity Entity
}

func (EntityRemovedEvent) EventType() EventType {
	return EntityRemovedEventType
}

type EntityAddedEvent struct {
	Entity Entity
}

func (EntityAddedEvent) EventType() EventType {
	return EntityAddedEventType
}
