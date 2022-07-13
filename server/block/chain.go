package block

import (
	"github.com/go-gl/mathgl/mgl64"
	"github.com/xJustJqy/MockingJay/server/block/cube"
	"github.com/xJustJqy/MockingJay/server/block/model"
	"github.com/xJustJqy/MockingJay/server/item"
	"github.com/xJustJqy/MockingJay/server/world"
)

// Chain is a metallic decoration block.
type Chain struct {
	transparent

	// Axis is the axis which the chain faces.
	Axis cube.Axis
}

// CanDisplace ...
func (Chain) CanDisplace(b world.Liquid) bool {
	_, water := b.(Water)
	return water
}

// SideClosed ...
func (Chain) SideClosed(cube.Pos, cube.Pos, *world.World) bool {
	return false
}

// UseOnBlock ...
func (c Chain) UseOnBlock(pos cube.Pos, face cube.Face, _ mgl64.Vec3, w *world.World, user item.User, ctx *item.UseContext) (used bool) {
	pos, face, used = firstReplaceable(w, pos, face, c)
	if !used {
		return
	}
	c.Axis = face.Axis()

	place(w, pos, c, user, ctx)
	return placed(ctx)
}

// BreakInfo ...
func (c Chain) BreakInfo() BreakInfo {
	return newBreakInfo(5, pickaxeHarvestable, pickaxeEffective, oneOf(c))
}

// EncodeItem ...
func (Chain) EncodeItem() (name string, meta int16) {
	return "minecraft:chain", 0
}

// EncodeBlock ...
func (c Chain) EncodeBlock() (string, map[string]any) {
	return "minecraft:chain", map[string]any{"pillar_axis": c.Axis.String()}
}

// Model ...
func (c Chain) Model() world.BlockModel {
	return model.Chain{Axis: c.Axis}
}

// allChains ...
func allChains() (chains []world.Block) {
	for _, axis := range cube.Axes() {
		chains = append(chains, Chain{Axis: axis})
	}
	return
}
