package entities

import (
	"feudal/internal/components"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	BaseEntity
}

func (p *Player) Update() {
	moving := false
	// Update velocity based on key presses
	if ebiten.IsKeyPressed(ebiten.KeyA) { // Move left
		p.Velocity.DX -= p.Acceleration.A
		moving = true
		p.Animation.Direction = components.Left
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) { // Move right
		p.Velocity.DX += p.Acceleration.A
		moving = true
		p.Animation.Direction = components.Right
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) { // Move up
		p.Velocity.DY -= p.Acceleration.A
		moving = true
		p.Animation.Direction = components.Up
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) { // Move down
		p.Velocity.DY += p.Acceleration.A
		moving = true
		p.Animation.Direction = components.Down
	}

	// Clamp velocity to prevent excessive speed
	p.Velocity.ClampVelocity()

	// Apply deceleration if no keys are pressed
	if !ebiten.IsKeyPressed(ebiten.KeyA) && !ebiten.IsKeyPressed(ebiten.KeyD) &&
		!ebiten.IsKeyPressed(ebiten.KeyW) && !ebiten.IsKeyPressed(ebiten.KeyS) {
		p.Velocity.ApplyDecelleration(p.Acceleration.DA)
	}

	// Handle idle animation logic
	p.Animation.Update(moving)

	// Update position using velocity
	p.Position.X += p.Velocity.DX
	p.Position.Y += p.Velocity.DY
}

func (p *Player) Draw(screen *ebiten.Image, camX, camY float64) {
	op := &ebiten.DrawImageOptions{}

	srcRect, flip := p.Animation.GetCurrentFrame()

	// If it needs to be flipped
	if flip {
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(32, 0) // Remember to shift
	}

	spriteSubImage := p.Sprite.Image.SubImage(srcRect).(*ebiten.Image)
	w, h := spriteSubImage.Bounds().Dx(), spriteSubImage.Bounds().Dy()

	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)         // Align center of image to player
	op.GeoM.Translate(p.Position.X-camX, p.Position.Y-camY) // Set the position of the image on the screen

	screen.DrawImage(spriteSubImage, op) // Draw the image on the screen
}

// Creates the animation component for the player
func LoadPlayerAnimations() *components.Animation {
	// Idle frames: first 2 frames in row 1.
	idleFrames := []image.Rectangle{
		image.Rect(0, 0, 32, 32),
		image.Rect(32, 0, 64, 32),
	}

	// Walking animations use all 4 frames per row.
	walkDownFrames := []image.Rectangle{
		image.Rect(0, 96, 32, 128),
		image.Rect(32, 96, 64, 128),
		image.Rect(64, 96, 96, 128),
		image.Rect(96, 96, 128, 128),
	}
	walkRightFrames := []image.Rectangle{
		image.Rect(0, 128, 32, 160),
		image.Rect(32, 128, 64, 160),
		image.Rect(64, 128, 96, 160),
		image.Rect(96, 128, 128, 160),
	}
	walkUpFrames := []image.Rectangle{
		image.Rect(0, 160, 32, 192),
		image.Rect(32, 160, 64, 192),
		image.Rect(64, 160, 96, 192),
		image.Rect(96, 160, 128, 192),
	}

	// Create the animation component.
	anim := components.Animation{
		IdleFrames: idleFrames,
		WalkingFrames: map[components.Direction][]image.Rectangle{
			components.Down:  walkDownFrames,
			components.Right: walkRightFrames,
			components.Up:    walkUpFrames,
			components.Left:  walkRightFrames,
		},
		State:      components.Idle,
		Direction:  components.Down, // Default facing direction
		FrameIndex: 0,
		Timer:      0,
		FrameDelay: 10, // Adjust for desired animation speed
	}
	return &anim
}
