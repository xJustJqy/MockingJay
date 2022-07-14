package item

import (
	"time"

	"github.com/xJustJqy/MockingJay/server/entity/effect"
	"github.com/xJustJqy/MockingJay/server/world"
)

// Pufferfish is a poisonous type of fish that is used to brew water breathing potions.
type Pufferfish struct {
	defaultFood
}

// Consume ...
func (p Pufferfish) Consume(_ *world.World, c Consumer) Stack {
	c.Saturate(1, 0.2)
	id, _ := effect.ID(effect.Hunger{})
	c.AddEffect(effect.New(id, effect.Hunger{}, 3, 15*time.Second))
	id, _ = effect.ID(effect.Poison{})
	c.AddEffect(effect.New(id, effect.Poison{}, 2, time.Minute))
	id, _ = effect.ID(effect.Nausea{})
	c.AddEffect(effect.New(id, effect.Nausea{}, 2, 15*time.Second))
	return Stack{}
}

// EncodeItem ...
func (p Pufferfish) EncodeItem() (name string, meta int16) {
	return "minecraft:pufferfish", 0
}
