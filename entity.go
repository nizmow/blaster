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

type FindComponentsResult struct {
	Entity             Entity
	RequestedComponent Component
}

type FindComponentsJoinResult struct {
	Entity              Entity
	RequestedComponents map[ComponentType]Component
}

func (world World) FindComponents(requestedComponent ComponentType) []FindComponentsResult {
	var results []FindComponentsResult

	for _, entity := range world.GetEntities() {
		for _, entityComponent := range entity.GetComponents() {
			if entityComponent.ComponentType() == requestedComponent {
				results = append(results, FindComponentsResult{entity, entityComponent})
			}
		}
	}

	return results
}

func (world World) FindComponentsJoin(requestedComponents ...ComponentType) []FindComponentsJoinResult {
	var results []FindComponentsJoinResult

	// for all entities in the world
	for _, entity := range world.GetEntities() {
		// for each component in this entity
		matchedComponents := make(map[ComponentType]Component)
		for _, entityComponent := range entity.GetComponents() {

			// for each component we requested
			for _, requestedComponent := range requestedComponents {
				if _, present := matchedComponents[requestedComponent]; !present {
					matchedComponents[requestedComponent] = nil
				}

				if entityComponent.ComponentType() == requestedComponent {
					matchedComponents[requestedComponent] = entityComponent
				}
			}

			// validate we got all of the components requested for each join
			successful := true
			for _, value := range matchedComponents {
				if value == nil {
					successful = false
					break
				}
			}
			if successful {
				r := FindComponentsJoinResult{entity, matchedComponents}
				results = append(results, r)
			}
		}
	}

	return results
}
