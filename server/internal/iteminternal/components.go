package iteminternal

import (
	"strings"

	"github.com/xJustJqy/MockingJay/server/item"
	"github.com/xJustJqy/MockingJay/server/world"
)

// Components returns all the components of the given custom item. If the item has no components, a nil map and false
// are returned.
func Components(it world.CustomItem) (map[string]any, bool) {
	category := it.Category()
	identifier, _ := it.EncodeItem()
	name := strings.Split(identifier, ":")[1]

	builder := NewComponentBuilder(it.Name(), identifier, category)

	if x, ok := it.(item.Armour); ok {
		builder.AddComponent("minecraft:armor", map[string]any{
			"protection": int32(x.DefencePoints()),
		})
		builder.AddComponent("minecraft:knockback_resistance", map[string]any{
			"protection": float32(x.KnockBackResistance()),
		})

		var slot int32
		switch it.(type) {
		case item.HelmetType:
			slot = 2
		case item.ChestplateType:
			slot = 3
		case item.LeggingsType:
			slot = 4
		case item.BootsType:
			slot = 5
		}
		builder.AddComponent("minecraft:wearable", map[string]any{
			"slot": slot,
		})
	}
	if x, ok := it.(item.Consumable); ok {
		builder.AddProperty("use_duration", int32(x.ConsumeDuration().Seconds()*20))
		builder.AddComponent("minecraft:food", map[string]any{
			"can_always_eat": x.AlwaysConsumable(),
		})

		if y, ok := it.(item.Drinkable); ok && y.Drinkable() {
			builder.AddProperty("use_animation", int32(2))
		} else {
			builder.AddProperty("use_animation", int32(1))
		}
	}
	if x, ok := it.(item.Cooldown); ok {
		builder.AddComponent("minecraft:cooldown", map[string]any{
			"category": name,
			"duration": float32(x.Cooldown().Seconds()),
		})
	}
	if x, ok := it.(item.Durable); ok {
		builder.AddComponent("minecraft:durability", map[string]any{
			"max_durability": int32(x.DurabilityInfo().MaxDurability),
		})
	}
	if x, ok := it.(item.MaxCounter); ok {
		builder.AddProperty("max_stack_size", int32(x.MaxCount()))
	}
	if x, ok := it.(item.OffHand); ok {
		builder.AddProperty("allow_off_hand", x.OffHand())
	}
	if x, ok := it.(item.Throwable); ok {
		// The data in minecraft:projectile is only used by vanilla server-side, but we must send at least an empty map
		// so the client will play the throwing animation.
		builder.AddComponent("minecraft:projectile", map[string]any{})
		builder.AddComponent("minecraft:throwable", map[string]any{
			"do_swing_animation": x.SwingAnimation(),
		})
	}
	if x, ok := it.(item.Glinted); ok {
		builder.AddProperty("foil", x.Glinted())
	}
	if x, ok := it.(item.HandEquipped); ok {
		builder.AddProperty("hand_equipped", x.HandEquipped())
	}

	// If an item has no new components or properties then it should not be considered a component-based item.
	if builder.Empty() {
		return nil, false
	}
	return builder.Construct(), true
}
