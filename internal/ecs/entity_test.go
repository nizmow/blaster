package ecs

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

	entityPlayer.components = []Component{componentPlayer, componentRenderable1}
	entityEnemy.components = []Component{componentEnemy, componentRenderable2}
	world.entities = []Entity{entityPlayer, entityEnemy}

	results := world.FindComponentsJoin(TestPlayerType, TestRenderableType)

	if len(results) != 1 {
		t.Errorf("Expected 1 result and got %v", len(results))
	}

	results2 := world.FindComponentsJoin(TestPlayerType, TestEnemyType)
	if len(results2) != 0 {
		t.Errorf("Expected 0 results and got %v", len(results2))
	}
}
