package main

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

func Components(world World, requestedComponent ComponentType) []Component {
	var results []Component

	for _, entity := range world.GetEntities() {
		for _, entityComponent := range entity.GetComponents() {
			if entityComponent.ComponentType() == requestedComponent {
				results = append(results, entityComponent)
			}
		}
	}

	return results
}

func ComponentsJoin(world World, requestedComponents ...ComponentType) []map[ComponentType]Component {
	var results []map[ComponentType]Component

	// for all entities in the world
	for _, entity := range world.GetEntities() {
		// for each component in this entity
		resultsForEntity := make(map[ComponentType]Component)
		for _, entityComponent := range entity.GetComponents() {

			// for each component we requested
			for _, requestedComponent := range requestedComponents {
				if _, present := resultsForEntity[requestedComponent]; !present {
					resultsForEntity[requestedComponent] = nil
				}

				if entityComponent.ComponentType() == requestedComponent {
					resultsForEntity[requestedComponent] = entityComponent
				}
			}

			// validate we got all of the components requested for each join
			complete := true
			for _, value := range resultsForEntity {
				if value == nil {
					complete = false
					break
				}
			}
			if complete {
				results = append(results, resultsForEntity)
			}
		}
	}

	return results
}
