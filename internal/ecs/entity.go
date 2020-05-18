package ecs

var entityIdSequence int

func init() {
	entityIdSequence = 0
}

type Entity struct {
	ID         int
	EntityName string
	components []Component
}

func NewEntity(entityName string) *Entity {
	entityIdSequence++
	return &Entity{ID: entityIdSequence, EntityName: entityName}
}

func (e *Entity) AddComponent(c Component) *Entity {
	// todo: components need to be mutable! I wish we could check the interface was hiding a reference
	e.components = append(e.components, c)
	return e
}

func (e *Entity) GetComponents() []Component {
	return e.components
}
