package block

import (
	"github.com/xJustJqy/MockingJay/server/block/cube"
	"github.com/xJustJqy/MockingJay/server/item"
	"github.com/xJustJqy/MockingJay/server/world"
)

// Placer represents an entity that is able to place a block at a specific position in the world.
type Placer interface {
	item.User
	PlaceBlock(pos cube.Pos, b world.Block, ctx *item.UseContext)
}
