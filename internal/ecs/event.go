package ecs

// EventType is the type of event -- we need this to avoid generics and reflection,
// though to be honest reflection might be good enough depending how fast it is.
type EventType int

// Event is something emitted by or queried for by a system. Events have types and
// may extend with additional data that can be used by consuming systems.
type Event interface {
	EventType() EventType
}

// EventHandlers will be called to handle any fired events of the particular type they're interested in.
type EventHandler interface {
	DesiredEventType() EventType
	HandleEvent(Event, *World)
}
