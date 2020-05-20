package ecs

var entityIDSequence int

func init() {
	entityIDSequence = 0
}

// Entity is a game entity.
type Entity struct {
	ID         int
	EntityName string
	components []Component
}

// NewEntity creates a new entity of a particular name. Names are purely informational and no
// uniqueness is enforced.
func NewEntity(entityName string) *Entity {
	entityIDSequence++
	return &Entity{ID: entityIDSequence, EntityName: entityName}
}

// AddComponent adds a component to an entity.
func (e *Entity) AddComponent(c Component) *Entity {
	// todo: components need to be mutable! I wish we could check the interface was hiding a reference
	e.components = append(e.components, c)
	return e
}

// GetComponents returns all components belonging to the entity.
func (e *Entity) GetComponents() []Component {
	return e.components
}

// GetComponent returns the first component of a requested ComponentType belonging
// to the entity, or nil if no such components were found.
func (e *Entity) GetComponent(componentType ComponentType) Component {
	for _, component := range e.components {
		if component.ComponentType() == componentType {
			return component
		}
	}

	return nil
}
