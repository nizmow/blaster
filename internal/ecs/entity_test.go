package ecs

import "testing"

const (
	SharedComponentType ComponentType = iota
	UniqueComponentType ComponentType = iota
)

type SharedComponent struct {
	Value int
}

func (*SharedComponent) ComponentType() ComponentType {
	return SharedComponentType
}

type UniqueComponent struct{}

func (*UniqueComponent) ComponentType() ComponentType {
	return UniqueComponentType
}

// Test_SharedComponent demonstrates that we can share single component between multiple entities
// (basically, I still feel the need to validate Go is doing what I want whenever I pass something to a function)
func Test_SharedComponent(t *testing.T) {
	firstEntity := NewEntity("first")
	secondEntity := NewEntity("second")
	var world World

	sharedComponent := SharedComponent{Value: 0}

	firstEntity.AddComponent(&sharedComponent)
	firstEntity.AddComponent(&UniqueComponent{})
	secondEntity.AddComponent(&sharedComponent)
	secondEntity.AddComponent(&UniqueComponent{})

	world.AddEntity(*firstEntity)
	world.AddEntity(*secondEntity)

	componentsToManipulate := world.FindEntitiesWithComponents(SharedComponentType, UniqueComponentType)
	for _, findComponentsResult := range componentsToManipulate {
		sharedComponentToUpdate := findComponentsResult.RequestedComponents[SharedComponentType].(*SharedComponent)
		sharedComponentToUpdate.Value++
	}

	if sharedComponent.Value != 2 {
		t.Error("Foo")
	}
}
