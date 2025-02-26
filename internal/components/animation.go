package components

import "image"

type AnimationState int
type Direction int

const (
	Idle AnimationState = iota
	Walking
)

const (
	Up Direction = iota
	Down
	Left
	Right
)

type Animation struct {
	IdleFrames    []image.Rectangle // Rectangles for where on the sprite sheet it is
	WalkingFrames map[Direction][]image.Rectangle
	// Current state
	State AnimationState
	// current direction
	Direction Direction

	// What frame we are currently on
	FrameIndex int
	Timer      int
	FrameDelay int
}

// Update the animation
func (anim *Animation) Update(moving bool) {
	// Store the previous state.
	prevState := anim.State

	// Reset to idle if not moving
	if !moving {
		anim.State = Idle
	} else {
		anim.State = Walking
	}

	// If the state has changed, reset frame index and timer.
	if anim.State != prevState {
		anim.FrameIndex = 0
		anim.Timer = 0
	}

	// Update the animation every FrameDelay number of frames
	anim.Timer++
	if anim.Timer >= anim.FrameDelay {
		anim.Timer = 0
		if anim.State == Idle {
			anim.FrameIndex = (anim.FrameIndex + 1) % len(anim.IdleFrames)
		} else if anim.State == Walking {
			frames := anim.WalkingFrames[anim.Direction]
			anim.FrameIndex = (anim.FrameIndex + 1) % len(frames)
		}
	}
}

// Get the current sprite image location (image.Rectangle) and if it needs to be flipped
func (anim *Animation) GetCurrentFrame() (image.Rectangle, bool) {
	// if we are idle
	if anim.State == Idle {
		return anim.IdleFrames[anim.FrameIndex], false
	}

	// If we are walking a specific direction
	frames := anim.WalkingFrames[anim.Direction]
	if anim.Direction == Left {
		return frames[anim.FrameIndex], true // Flip for left movement
	}

	return frames[anim.FrameIndex], false
}

// Return the number of frames a specific state has
func (anim *Animation) frameCount() int {
	if anim.State == Idle {
		return len(anim.IdleFrames)
	}
	return len(anim.WalkingFrames[anim.Direction])
}
