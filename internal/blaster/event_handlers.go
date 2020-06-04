package blaster

import (
	"fmt"
	"github.com/nizmow/blaster/internal/ecs"
)

type BaddieCollisionEventHandler struct{}

func (BaddieCollisionEventHandler) DesiredEventType() ecs.EventType {
	return CollisionEventType
}

func (BaddieCollisionEventHandler) HandleEvent(e ecs.Event, world *ecs.World) {
	event := e.(*CollisionEvent)

	var baddieEntity *ecs.Entity
	var bulletEntity *ecs.Entity
	for _, entity := range event.InvolvedEntities {
		if entity.GetComponent(BaddieType) != nil {
			baddieEntity = &entity
		} else if entity.GetComponent(PlayerBulletType) != nil {
			bulletEntity = &entity
		}
	}

	if baddieEntity != nil && bulletEntity != nil {
		// collision detected!
		world.RemoveEntity(baddieEntity.ID)
		world.RemoveEntity(bulletEntity.ID)
	}
}

type EntityRemovedEventHandler struct{}

func (EntityRemovedEventHandler) DesiredEventType() ecs.EventType {
	return ecs.EntityRemovedEventType
}

func (EntityRemovedEventHandler) HandleEvent(e ecs.Event, world *ecs.World) {
	event := e.(ecs.EntityRemovedEvent)
	fmt.Printf("Entity '%s' removed!\n", event.Entity.EntityName)
}
