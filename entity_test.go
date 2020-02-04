package main

import "testing"

type TestRenderable struct{}

func (TestRenderable) ComponentName() string {
	return "TestRenderable"
}

type TestPlayer struct{}

func (TestPlayer) ComponentName() string {
	return "TestPlayer"
}

type TestEnemy struct{}

func (TestEnemy) ComponentName() string {
	return "TestEnemy"
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

	results := ComponentsJoin(world, "TestPlayer", "TestRenderable")

	if len(results) != 1 {
		t.Errorf("Expected 1 result and got %v", len(results))
	}

	results2 := ComponentsJoin(world, "TestPlayer", "TestEnemy")
	if len(results2) != 0 {
		t.Errorf("Expected 0 results and got %v", len(results2))
	}
}
