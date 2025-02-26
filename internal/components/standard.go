package components

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

// Position component
type Position struct {
	X, Y float64
}

// Velocity component
type Velocity struct {
	DX, DY, MaxSpeed, StopThreshold float64
}

// Acceleration component
type Acceleration struct {
	A, DA float64
}

// If an entity has a sprite
type Sprite struct {
	Image *ebiten.Image
}

func (v *Velocity) ApplyDecelleration(decelleration float64) {
	if math.Abs(v.DX) > v.StopThreshold {
		if v.DX > 0 {
			v.DX -= decelleration
		} else if v.DX < 0 {
			v.DX += decelleration
		} else {
			v.DX = 0
		}
	}

	if math.Abs(v.DY) > v.StopThreshold {
		if v.DY > 0 {
			v.DY -= decelleration
		} else if v.DY < 0 {
			v.DY += decelleration
		} else {
			v.DY = 0
		}
	}
}

// Limit the speed to max_speed
func (v *Velocity) ClampVelocity() {
	speed := math.Pow((math.Pow(v.DX, 2) + math.Pow(v.DY, 2)), 0.5)
	if speed > v.MaxSpeed {
		scale := v.MaxSpeed / speed
		v.DX = v.DX * scale
		v.DY = v.DY * scale
	}
}
