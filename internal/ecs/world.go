package ecs

import "github.com/hajimehoshi/ebiten"

type World struct {
	entities      []Entity
	ScreenWidth   int
	ScreenHeight  int
	events        map[EventType]Event
	logicSystems  []LogicSystem
	renderSystems []RenderSystem
}

type FindEntitiesWithComponentResult struct {
	Entity             Entity
	RequestedComponent Component
}

type FindEntitiesWithComponentsResult struct {
	Entity              Entity
	RequestedComponents map[ComponentType]Component
}

// RenderTick ticks over the rendering to the given image, mostly because this is the way ebiten wants to work.
func (w *World) RenderTick(screen *ebiten.Image) error {
	for _, rs := range w.renderSystems {
		err := rs.Update(w, screen)
		if err != nil {
			return err
		}
	}

	return nil
}

// LogicTick ticks over the game logic, so we can run logic independently of rendering if we need to. Mostly it's a
// bit because this is the way ebiten wants to work.
func (w *World) LogicTick() error {
	for _, ls := range w.logicSystems {
		err := ls.Update(w)
		if err != nil {
			return err
		}
	}

	return nil
}

// AddEntity adds an entity to the world.
func (world *World) AddEntity(e Entity) *World {
	world.entities = append(world.entities, e)
	return world
}

// RemoveEntity removes an entity from the world.
func (world *World) RemoveEntity(idToRemove int) {
	for i, e := range world.entities {
		if e.ID == idToRemove {
			world.entities = append(world.entities[:i], world.entities[i+1:]...)
			break
		}
	}
}

// GetEntities returns all entities in the world.
func (world *World) GetEntities() []Entity {
	return world.entities
}

// FindEntitiesWithComponent finds ALL entities that have the requested component attached.
func (world World) FindEntitiesWithComponent(requestedComponent ComponentType) []FindEntitiesWithComponentResult {
	var results []FindEntitiesWithComponentResult

	for _, entity := range world.GetEntities() {
		for _, entityComponent := range entity.GetComponents() {
			if entityComponent.ComponentType() == requestedComponent {
				results = append(results, FindEntitiesWithComponentResult{entity, entityComponent})
			}
		}
	}

	return results
}

// FindEntitiesWithComponents finds ALL entities that have ALL the requested components attached.
func (world World) FindEntitiesWithComponents(requestedComponents ...ComponentType) []FindEntitiesWithComponentsResult {
	var results []FindEntitiesWithComponentsResult

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
				r := FindEntitiesWithComponentsResult{entity, matchedComponents}
				results = append(results, r)
			}
		}
	}

	return results
}

func (w *World) FireEvent(e Event) {
	w.events[e.EventType()] = e
}

func (w *World) AddRenderSystem(s RenderSystem) {
	w.renderSystems = append(w.renderSystems, s)
}

func (w *World) AddLogicSystem(s LogicSystem) {
	w.logicSystems = append(w.logicSystems, s)
}
