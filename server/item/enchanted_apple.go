package item

import (
	"time"

	"github.com/xJustJqy/MockingJay/server/entity/effect"
	"github.com/xJustJqy/MockingJay/server/world"
)

// EnchantedApple is a rare variant of the golden apple that has stronger effects.
type EnchantedApple struct{}

// AlwaysConsumable ...
func (EnchantedApple) AlwaysConsumable() bool {
	return true
}

// ConsumeDuration ...
func (EnchantedApple) ConsumeDuration() time.Duration {
	return DefaultConsumeDuration
}

// Consume ...
func (EnchantedApple) Consume(_ *world.World, c Consumer) Stack {
	c.Saturate(4, 9.6)
	id, _ := effect.ID(effect.Absorption{})
	c.AddEffect(effect.New(id, effect.Absorption{}, 4, 2*time.Minute))
	id, _ = effect.ID(effect.Regeneration{})
	c.AddEffect(effect.New(id, effect.Regeneration{}, 5, 30*time.Second))
	id, _ = effect.ID(effect.FireResistance{})
	c.AddEffect(effect.New(id, effect.FireResistance{}, 1, 5*time.Minute))
	id, _ = effect.ID(effect.Resistance{})
	c.AddEffect(effect.New(id, effect.Resistance{}, 1, 5*time.Minute))
	return Stack{}
}

// EncodeItem ...
func (EnchantedApple) EncodeItem() (name string, meta int16) {
	return "minecraft:enchanted_golden_apple", 0
}
