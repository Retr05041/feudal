package camera

import "feudal/internal/entities"

type Camera struct {
	FollowingPlayer *entities.Player

	X, Y float64

	ScreenWidth, ScreenHeight, WorldWidth, WorldHeight float64
}

func (c *Camera) Update() {
	const cameraLerpFactor = 0.1 // Adjust for smoothness

	targetX := c.FollowingPlayer.Position.X - float64(c.ScreenWidth)/2
	targetY := c.FollowingPlayer.Position.Y - float64(c.ScreenHeight)/2

	// Lerp towards target position
	c.X += (targetX - c.X) * cameraLerpFactor
	c.Y += (targetY - c.Y) * cameraLerpFactor

	// Clamp camera to world bounds (if applicable)
	if c.X < 0 {
		c.X = 0
	}
	if c.Y < 0 {
		c.Y = 0
	}
	if c.X > c.WorldWidth-c.ScreenWidth {
		c.X = c.WorldWidth - c.ScreenWidth
	}
	if c.Y > c.WorldHeight-c.ScreenHeight {
		c.Y = c.WorldHeight - c.ScreenHeight
	}
}
