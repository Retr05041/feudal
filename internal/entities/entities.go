package entities

import (
	"feudal/internal/components"

	"github.com/hajimehoshi/ebiten/v2"
)

// Defalt entity
type Entity interface {
	Update()
	Draw(screen *ebiten.Image, camX, camY float64)
}

// Shared attributes amongst entities
type BaseEntity struct {
	Position     *components.Position
	Velocity     *components.Velocity
	Acceleration *components.Acceleration
	Animation    *components.Animation
	Sprite       *components.Sprite
}
