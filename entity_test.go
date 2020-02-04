package main

import "testing"

const (
	TestRenderableType ComponentType = iota
	TestPlayerType     ComponentType = iota
	TestEnemyType      ComponentType = iota
)

type TestRenderable struct{}

func (TestRenderable) ComponentType() ComponentType {
	return TestRenderableType
}

type TestPlayer struct{}

func (TestPlayer) ComponentType() ComponentType {
	return TestPlayerType
}

type TestEnemy struct{}

func (TestEnemy) ComponentType() ComponentType {
	return TestEnemyType
}

func Test_ComponentsJoin(t *testing.T) {
	// big boring setup -- this might be a good candidate for creating a builder pattern later

	world := World{}
	entityPlayer := Entity{}
	entityEnemy := Entity{}
	componentRenderable1 := TestRenderable{}
	componentRenderable2 := TestRenderable{}
	componentPlayer := TestPlayer{}
	componentEnemy := TestEnemy{}

	entityPlayer.Components = []Component{componentPlayer, componentRenderable1}
	entityEnemy.Components = []Component{componentEnemy, componentRenderable2}
	world.Entities = []Entity{entityPlayer, entityEnemy}

	results := ComponentsJoin(world, TestPlayerType, TestRenderableType)

	if len(results) != 1 {
		t.Errorf("Expected 1 result and got %v", len(results))
	}

	results2 := ComponentsJoin(world, TestPlayerType, TestEnemyType)
	if len(results2) != 0 {
		t.Errorf("Expected 0 results and got %v", len(results2))
	}
}
