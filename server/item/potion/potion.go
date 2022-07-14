package potion

import (
	"time"

	"github.com/xJustJqy/MockingJay/server/entity/effect"
)

// Potion holds the effects given by the potion type.
type Potion struct {
	potion
}

// Water ...
func Water() Potion {
	return Potion{}
}

// Mundane ...
func Mundane() Potion {
	return Potion{1}
}

// LongMundane ...
func LongMundane() Potion {
	return Potion{2}
}

// Thick ...
func Thick() Potion {
	return Potion{3}
}

// Awkward ...
func Awkward() Potion {
	return Potion{4}
}

// NightVision ...
func NightVision() Potion {
	return Potion{5}
}

// LongNightVision ...
func LongNightVision() Potion {
	return Potion{6}
}

// Invisibility ...
func Invisibility() Potion {
	return Potion{7}
}

// LongInvisibility ...
func LongInvisibility() Potion {
	return Potion{8}
}

// Leaping ...
func Leaping() Potion {
	return Potion{9}
}

// LongLeaping ...
func LongLeaping() Potion {
	return Potion{10}
}

// StrongLeaping ...
func StrongLeaping() Potion {
	return Potion{11}
}

// FireResistance ...
func FireResistance() Potion {
	return Potion{12}
}

// LongFireResistance ...
func LongFireResistance() Potion {
	return Potion{13}
}

// Swiftness ...
func Swiftness() Potion {
	return Potion{14}
}

// LongSwiftness ...
func LongSwiftness() Potion {
	return Potion{15}
}

// StrongSwiftness ...
func StrongSwiftness() Potion {
	return Potion{16}
}

// Slowness ...
func Slowness() Potion {
	return Potion{17}
}

// LongSlowness ...
func LongSlowness() Potion {
	return Potion{18}
}

// WaterBreathing ...
func WaterBreathing() Potion {
	return Potion{19}
}

// LongWaterBreathing ...
func LongWaterBreathing() Potion {
	return Potion{20}
}

// Healing ...
func Healing() Potion {
	return Potion{21}
}

// StrongHealing ...
func StrongHealing() Potion {
	return Potion{22}
}

// Harming ...
func Harming() Potion {
	return Potion{23}
}

// StrongHarming ...
func StrongHarming() Potion {
	return Potion{24}
}

// Poison ...
func Poison() Potion {
	return Potion{25}
}

// LongPoison ...
func LongPoison() Potion {
	return Potion{26}
}

// StrongPoison ...
func StrongPoison() Potion {
	return Potion{27}
}

// Regeneration ...
func Regeneration() Potion {
	return Potion{28}
}

// LongRegeneration ...
func LongRegeneration() Potion {
	return Potion{29}
}

// StrongRegeneration ...
func StrongRegeneration() Potion {
	return Potion{30}
}

// Strength ...
func Strength() Potion {
	return Potion{31}
}

// LongStrength ...
func LongStrength() Potion {
	return Potion{32}
}

// StrongStrength ...
func StrongStrength() Potion {
	return Potion{33}
}

// Weakness ...
func Weakness() Potion {
	return Potion{34}
}

// LongWeakness ...
func LongWeakness() Potion {
	return Potion{35}
}

// Wither ...
func Wither() Potion {
	return Potion{36}
}

// TurtleMaster ...
func TurtleMaster() Potion {
	return Potion{37}
}

// LongTurtleMaster ...
func LongTurtleMaster() Potion {
	return Potion{38}
}

// StrongTurtleMaster ...
func StrongTurtleMaster() Potion {
	return Potion{39}
}

// SlowFalling ...
func SlowFalling() Potion {
	return Potion{40}
}

// LongSlowFalling ...
func LongSlowFalling() Potion {
	return Potion{41}
}

// StrongSlowness ...
func StrongSlowness() Potion {
	return Potion{42}
}

// From returns a Potion by the ID given.
func From(id int32) Potion {
	return Potion{potion(id)}
}

// Effects returns the effects of the potion.
func (p Potion) Effects() []effect.Effect {
	switch p {
	case NightVision():
		id, _ := effect.ID(effect.NightVision{})
		return []effect.Effect{effect.New(id, effect.NightVision{}, 1, 3*time.Minute)}
	case LongNightVision():
		id, _ := effect.ID(effect.NightVision{})
		return []effect.Effect{effect.New(id, effect.NightVision{}, 1, 8*time.Minute)}
	case Invisibility():
		id, _ := effect.ID(effect.Invisibility{})
		return []effect.Effect{effect.New(id, effect.Invisibility{}, 1, 3*time.Minute)}
	case LongInvisibility():
		id, _ := effect.ID(effect.Invisibility{})
		return []effect.Effect{effect.New(id, effect.Invisibility{}, 1, 8*time.Minute)}
	case Leaping():
		id, _ := effect.ID(effect.JumpBoost{})
		return []effect.Effect{effect.New(id, effect.JumpBoost{}, 1, 3*time.Minute)}
	case LongLeaping():
		id, _ := effect.ID(effect.JumpBoost{})
		return []effect.Effect{effect.New(id, effect.JumpBoost{}, 1, 8*time.Minute)}
	case StrongLeaping():
		id, _ := effect.ID(effect.JumpBoost{})
		return []effect.Effect{effect.New(id,effect.JumpBoost{}, 2, 90*time.Second)}
	case FireResistance():
		id, _ := effect.ID(effect.FireResistance{})
		return []effect.Effect{effect.New(id, effect.FireResistance{}, 1, 3*time.Minute)}
	case LongFireResistance():
		id, _ := effect.ID(effect.FireResistance{})
		return []effect.Effect{effect.New(id, effect.FireResistance{}, 1, 8*time.Minute)}
	case Swiftness():
		id, _ := effect.ID(effect.Speed{})
		return []effect.Effect{effect.New(id, effect.Speed{}, 1, 3*time.Minute)}
	case LongSwiftness():
		id, _ := effect.ID(effect.Speed{})
		return []effect.Effect{effect.New(id, effect.Speed{}, 1, 8*time.Minute)}
	case StrongSwiftness():
		id, _ := effect.ID(effect.Speed{})
		return []effect.Effect{effect.New(id, effect.Speed{}, 2, 90*time.Second)}
	case Slowness():
		id, _ := effect.ID(effect.Slowness{})
		return []effect.Effect{effect.New(id, effect.Slowness{}, 1, 90*time.Second)}
	case LongSlowness():
		id, _ := effect.ID(effect.Slowness{})
		return []effect.Effect{effect.New(id, effect.Slowness{}, 1, 4*time.Minute)}
	case WaterBreathing():
		id, _ := effect.ID(effect.WaterBreathing{})
		return []effect.Effect{effect.New(id, effect.WaterBreathing{}, 1, 3*time.Minute)}
	case LongWaterBreathing():
		id, _ := effect.ID(effect.WaterBreathing{})
		return []effect.Effect{effect.New(id, effect.WaterBreathing{}, 1, 8*time.Minute)}
	case Healing():
		id, _ := effect.ID(effect.InstantHealth{})
		return []effect.Effect{effect.NewInstant(id,effect.InstantHealth{}, 1)}
	case StrongHealing():
		id, _ := effect.ID(effect.InstantHealth{})
		return []effect.Effect{effect.NewInstant(id, effect.InstantHealth{}, 2)}
	case Harming():
		id, _ := effect.ID(effect.InstantDamage{})
		return []effect.Effect{effect.NewInstant(id, effect.InstantDamage{}, 1)}
	case StrongHarming():
		id, _ := effect.ID(effect.InstantDamage{})
		return []effect.Effect{effect.NewInstant(id, effect.InstantDamage{}, 2)}
	case Poison():
		id, _ := effect.ID(effect.Poison{})
		return []effect.Effect{effect.New(id, effect.Poison{}, 1, 45*time.Second)}
	case LongPoison():
		id, _ := effect.ID(effect.Poison{})
		return []effect.Effect{effect.New(id, effect.Poison{}, 1, 2*time.Minute)}
	case StrongPoison():
		id, _ := effect.ID(effect.Poison{})
		return []effect.Effect{effect.New(id, effect.Poison{}, 2, 22500*time.Millisecond)}
	case Regeneration():
		id, _ := effect.ID(effect.Regeneration{})
		return []effect.Effect{effect.New(id, effect.Regeneration{}, 1, 45*time.Second)}
	case LongRegeneration():
		id, _ := effect.ID(effect.Regeneration{})
		return []effect.Effect{effect.New(id, effect.Regeneration{}, 1, 2*time.Minute)}
	case StrongRegeneration():
		id, _ := effect.ID(effect.Regeneration{})
		return []effect.Effect{effect.New(id, effect.Regeneration{}, 2, 22*time.Second)}
	case Strength():
		id, _ := effect.ID(effect.Strength{})
		return []effect.Effect{effect.New(id, effect.Strength{}, 1, 3*time.Minute)}
	case LongStrength():
		id, _ := effect.ID(effect.Strength{})
		return []effect.Effect{effect.New(id, effect.Strength{}, 1, 8*time.Minute)}
	case StrongStrength():
		id, _ := effect.ID(effect.Strength{})
		return []effect.Effect{effect.New(id, effect.Strength{}, 2, 90*time.Second)}
	case Weakness():
		id, _ := effect.ID(effect.Weakness{})
		return []effect.Effect{effect.New(id, effect.Weakness{}, 1, 90*time.Second)}
	case LongWeakness():
		id, _ := effect.ID(effect.Weakness{})
		return []effect.Effect{effect.New(id, effect.Weakness{}, 1, 4*time.Minute)}
	case Wither():
		id, _ := effect.ID(effect.Wither{})
		return []effect.Effect{effect.New(id, effect.Wither{}, 1, 40*time.Second)}
	case TurtleMaster():
		id1, _ := effect.ID(effect.Resistance{})
		id2, _ := effect.ID(effect.Slowness{})
		return []effect.Effect{
			effect.New(id1, effect.Resistance{}, 3, 20*time.Second),
			effect.New(id2, effect.Slowness{}, 4, 20*time.Second),
		}
	case LongTurtleMaster():
		id1, _ := effect.ID(effect.Resistance{})
		id2, _ := effect.ID(effect.Slowness{})
		return []effect.Effect{
			effect.New(id1, effect.Resistance{}, 3, 40*time.Second),
			effect.New(id2, effect.Slowness{}, 4, 40*time.Second),
		}
	case StrongTurtleMaster():
		id1, _ := effect.ID(effect.Resistance{})
		id2, _ := effect.ID(effect.Slowness{})
		return []effect.Effect{
			effect.New(id1, effect.Resistance{}, 5, 20*time.Second),
			effect.New(id2, effect.Slowness{}, 6, 20*time.Second),
		}
	case SlowFalling():
		id, _ := effect.ID(effect.SlowFalling{})
		return []effect.Effect{effect.New(id, effect.SlowFalling{}, 1, 90*time.Second)}
	case LongSlowFalling():
		id, _ := effect.ID(effect.SlowFalling{})
		return []effect.Effect{effect.New(id, effect.SlowFalling{}, 1, 4*time.Minute)}
	case StrongSlowness():
		id, _ := effect.ID(effect.Slowness{})
		return []effect.Effect{effect.New(id, effect.Slowness{}, 4, 20*time.Second)}
	}
	return []effect.Effect{}
}

type potion uint8

// Uint8 returns the potion type as a uint8.
func (p potion) Uint8() uint8 {
	return uint8(p)
}

// All ...
func All() []Potion {
	return []Potion{
		Water(), Mundane(), LongMundane(), Thick(), Awkward(), NightVision(), LongNightVision(), Invisibility(),
		LongInvisibility(), Leaping(), LongLeaping(), StrongLeaping(), FireResistance(), LongFireResistance(),
		Swiftness(), LongSwiftness(), StrongSwiftness(), Slowness(), LongSlowness(), WaterBreathing(),
		LongWaterBreathing(), Healing(), StrongHealing(), Harming(), StrongHarming(), Poison(), LongPoison(),
		StrongPoison(), Regeneration(), LongRegeneration(), StrongRegeneration(), Strength(), LongStrength(),
		StrongStrength(), Weakness(), LongWeakness(), Wither(), TurtleMaster(), LongTurtleMaster(), StrongTurtleMaster(),
		SlowFalling(), LongSlowFalling(), StrongSlowness(),
	}
}
