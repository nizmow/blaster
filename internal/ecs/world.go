package ecs

import "github.com/hajimehoshi/ebiten"

type World struct {
	ScreenWidth  int
	ScreenHeight int

	entities      []Entity
	events        map[EventType][]Event
	logicSystems  []LogicSystem
	renderSystems []RenderSystem
	eventHandlers []EventHandler
}

type FindEntitiesWithComponentResult struct {
	Entity             Entity
	RequestedComponent Component
}

type FindEntitiesWithComponentsResult struct {
	Entity              Entity
	RequestedComponents map[ComponentType]Component
}

func NewWorld(screenWidth int, screenHeight int) World {
	return World{
		events:       make(map[EventType][]Event),
		ScreenWidth:  screenWidth,
		ScreenHeight: screenHeight,
	}
}

// Tick runs the game world, logic then renderer. Please pass an image that's a buffer here, so it can be swapped to
// screen during render to ensure things keep running correctly.
func (w *World) Tick(screen *ebiten.Image) error {
	// Run logic systems first to ensure the world is up to date before firing events
	for _, ls := range w.logicSystems {
		err := ls.Update(w)
		if err != nil {
			return err
		}
	}

	// Run event handlers to ensure the world is up to date before rendering
	for _, eh := range w.eventHandlers {
		for _, validEvent := range w.events[eh.DesiredEventType()] {
			eh.HandleEvent(validEvent, w)
		}
	}

	for _, rs := range w.renderSystems {
		err := rs.Update(w, screen)
		if err != nil {
			return err
		}
	}

	return nil
}

// AddEntity adds an entity to the world.
func (w *World) AddEntity(e Entity) *World {
	w.entities = append(w.entities, e)
	return w
}

// RemoveEntity removes an entity from the world.
func (w *World) RemoveEntity(idToRemove int) {
	for i, e := range w.entities {
		if e.ID == idToRemove {
			w.entities = append(w.entities[:i], w.entities[i+1:]...)
			break
		}
	}
}

// GetEntities returns all entities in the world.
func (w *World) GetEntities() []Entity {
	return w.entities
}

// FindEntitiesWithComponent finds ALL entities that have the requested component attached.
func (w World) FindEntitiesWithComponent(requestedComponent ComponentType) []FindEntitiesWithComponentResult {
	var results []FindEntitiesWithComponentResult

	for _, entity := range w.GetEntities() {
		for _, entityComponent := range entity.GetComponents() {
			if entityComponent.ComponentType() == requestedComponent {
				results = append(results, FindEntitiesWithComponentResult{entity, entityComponent})
			}
		}
	}

	return results
}

// FindEntitiesWithComponents finds ALL entities that have ALL the requested components attached.
func (w World) FindEntitiesWithComponents(requestedComponents ...ComponentType) []FindEntitiesWithComponentsResult {
	var results []FindEntitiesWithComponentsResult

	// for all entities in the world
	for _, entity := range w.GetEntities() {
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
	w.events[e.EventType()] = append(w.events[e.EventType()], e)
}

func (w *World) AddRenderSystem(s RenderSystem) {
	w.renderSystems = append(w.renderSystems, s)
}

func (w *World) AddLogicSystem(s LogicSystem) {
	w.logicSystems = append(w.logicSystems, s)
}

func (w *World) AddEventHandler(e EventHandler) {
	w.eventHandlers = append(w.eventHandlers, e)
}
