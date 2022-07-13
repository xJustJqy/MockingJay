package model

import (
	"github.com/xJustJqy/MockingJay/server/block/cube"
	"github.com/xJustJqy/MockingJay/server/world"
)

// Empty is a model that is completely empty. It has no collision boxes or solid faces.
type Empty struct{}

// BBox returns an empty slice.
func (Empty) BBox(cube.Pos, *world.World) []cube.BBox {
	return nil
}

// FaceSolid always returns false.
func (Empty) FaceSolid(cube.Pos, cube.Face, *world.World) bool {
	return false
}
