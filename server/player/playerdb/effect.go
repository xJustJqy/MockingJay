package playerdb

import "github.com/xJustJqy/MockingJay/server/entity/effect"

func effectsToData(effects []effect.Effect) []jsonEffect {
	data := make([]jsonEffect, len(effects))
	for key, eff := range effects {
		id, ok := effect.ID(eff.Type())
		if !ok {
			continue
		}
		data[key] = jsonEffect{
			ID:       id,
			Duration: eff.Duration(),
			Level:    eff.Level(),
			Ambient:  eff.Ambient(),
		}
	}
	return data
}

func dataToEffects(data []jsonEffect) []effect.Effect {
	effects := make([]effect.Effect, len(data))
	for i, d := range data {
		e, ok := effect.ByID(d.ID)
		if !ok {
			continue
		}
		switch eff := e.(type) {
		case effect.LastingType:
			if d.Ambient {
				effects[i] = effect.NewAmbient(d.ID, eff, d.Level, d.Duration)
				continue
			}
			effects[i] = effect.New(d.ID, eff, d.Level, d.Duration)
		default:
			effects[i] = effect.NewInstant(d.ID, eff, d.Level)
		}
	}
	return effects
}
