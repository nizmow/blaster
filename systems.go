package main

import "github.com/hajimehoshi/ebiten"

type Renderer struct{}

func (Renderer) Update(world World, screen *ebiten.Image) error {
	renderables := Components(world, RenderableType)
	for _, renderCandidate := range renderables {
		render := renderCandidate.(*Renderable)
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(float64(render.X), float64(render.Y))
		return screen.DrawImage(render.Image, opts)
	}
	return nil
}
