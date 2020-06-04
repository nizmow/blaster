package ecs

import "errors"

type Point struct {
	X int
	Y int
}

type Rect struct {
	topLeft     Point
	bottomRight Point
}

func NewRect(topLeft Point, bottomRight Point) (Rect, error) {
	if topLeft.X > bottomRight.X {
		return Rect{}, errors.New("unable to create rect with right on the left")
	}

	if topLeft.Y > bottomRight.Y {
		return Rect{}, errors.New("unable to create rect with top on the bottom")
	}

	return Rect{topLeft, bottomRight}, nil
}

// check if a rect overlaps another rect
func (r Rect) Intersects(target Rect) bool {
	return r.topLeft.X < target.bottomRight.X && r.bottomRight.X > target.topLeft.X && r.topLeft.Y < target.bottomRight.Y && r.bottomRight.Y > target.topLeft.Y
}
