package main

type Entity struct {
	ID         int
	EntityName string
	Components []Component
}

// todo: work out whether I should pass variable params as pointers? eg: c ...*Component
func ComponentsJoin(world World, requestedComponents ...string) []map[string]Component {

	var results []map[string]Component

	// for all entities in the world
	for _, entity := range world.Entities {
		// for each component in this entity
		resultsForEntity := make(map[string]Component)
		for _, entityComponent := range entity.Components {

			// for each component we requested
			for _, requestedComponent := range requestedComponents {
				if _, present := resultsForEntity[requestedComponent]; !present {
					resultsForEntity[requestedComponent] = nil
				}

				if entityComponent.ComponentName() == requestedComponent {
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
