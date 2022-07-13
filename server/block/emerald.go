package block

import (
	"github.com/xJustJqy/MockingJay/server/item"
	"github.com/xJustJqy/MockingJay/server/world/sound"
)

// Emerald is a precious mineral block crafted using 9 emeralds.
type Emerald struct {
	solid
}

// Instrument ...
func (e Emerald) Instrument() sound.Instrument {
	return sound.Bit()
}

// BreakInfo ...
func (e Emerald) BreakInfo() BreakInfo {
	return newBreakInfo(5, func(t item.Tool) bool {
		return t.ToolType() == item.TypePickaxe && t.HarvestLevel() >= item.ToolTierIron.HarvestLevel
	}, pickaxeEffective, oneOf(e))
}

// PowersBeacon ...
func (Emerald) PowersBeacon() bool {
	return true
}

// EncodeItem ...
func (Emerald) EncodeItem() (name string, meta int16) {
	return "minecraft:emerald_block", 0
}

// EncodeBlock ...
func (Emerald) EncodeBlock() (string, map[string]any) {
	return "minecraft:emerald_block", nil
}
