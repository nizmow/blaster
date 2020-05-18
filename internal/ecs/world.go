package ecs

type World struct {
	entities []Entity
}

type FindComponentsResult struct {
	Entity             Entity
	RequestedComponent Component
}

type FindComponentsJoinResult struct {
	Entity              Entity
	RequestedComponents map[ComponentType]Component
}

func (world *World) AddEntity(e Entity) *World {
	world.entities = append(world.entities, e)
	return world
}

func (world *World) RemoveEntity(idToRemove int) {
	for i, e := range world.entities {
		if e.ID == idToRemove {
			world.entities = append(world.entities[:i], world.entities[i+1:]...)
			break
		}
	}
}

func (world *World) GetEntities() []Entity {
	return world.entities
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
