package item

import (
	"time"

	"github.com/xJustJqy/MockingJay/server/entity/effect"
	"github.com/xJustJqy/MockingJay/server/world"
)

// GoldenApple is a special food item that bestows beneficial effects.
type GoldenApple struct{}

// AlwaysConsumable ...
func (e GoldenApple) AlwaysConsumable() bool {
	return true
}

// ConsumeDuration ...
func (e GoldenApple) ConsumeDuration() time.Duration {
	return DefaultConsumeDuration
}

// Consume ...
func (e GoldenApple) Consume(_ *world.World, c Consumer) Stack {
	c.Saturate(4, 9.6)
	id, _ := effect.ID(effect.Absorption{})
	c.AddEffect(effect.New(id, effect.Absorption{}, 1, 2*time.Minute))
	id, _ = effect.ID(effect.Regeneration{})
	c.AddEffect(effect.New(id, effect.Regeneration{}, 2, 5*time.Second))
	return Stack{}
}

// EncodeItem ...
func (e GoldenApple) EncodeItem() (name string, meta int16) {
	return "minecraft:golden_apple", 0
}
