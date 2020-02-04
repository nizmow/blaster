package main

import "github.com/hajimehoshi/ebiten"

type System interface {
	Update(World)
}

type Renderer struct {
	screen *ebiten.Image
}

func NewRenderer(screen *ebiten.Image) Renderer {
	return Renderer{screen}
}

func (r Renderer) Update(world World) {
	renderables := Components(world, RenderableType)
	for _, renderCandidate := range renderables {
		render := renderCandidate.(*Renderable)
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(float64(render.X), float64(render.Y))
		r.screen.DrawImage(render.Image, opts)
	}
}
