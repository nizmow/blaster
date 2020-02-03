package main

type System interface {
	Update()
}


type Renderer struct {
}
func (Renderer) Update(entities *[]Entity) {
	//for _, entity := range *entities {
	//	renderable := sort.Search(len(entity.components), func(int i) { return entity.components[i].ComponentName() == "Renderable" })
	//}
}
