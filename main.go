package main

import (
	"image/color"
	"log"

	"feudal/internal/components"
	"feudal/internal/entities"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Game struct contains - Needed for life cycle
type Game struct {
	entities []entities.Entity
}

// Update handles game logic
func (g *Game) Update() error {
	for i := range g.entities {
		g.entities[i].Update()
	}

	return nil
}

// Draw renders the game screen
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 255}) // Clear the screen with black

	for i := range g.entities {
		g.entities[i].Draw(screen)
	}
}

// Layout specifies the game's screen size
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 320, 240
}

func main() {
	// Load the player sprite
	playerImg, _, err := ebitenutil.NewImageFromFile("assets/Prototype_Character/Prototype_Character.png")
	if err != nil {
		log.Fatal(err)
	}

	player := &entities.Player{
		BaseEntity: entities.BaseEntity{
			Position:     &components.Position{X: 160, Y: 120}, // Initial position in the center of the screen
			Velocity:     &components.Velocity{DX: 0, DY: 0, MaxSpeed: 1, StopThreshold: 0.01},
			Acceleration: &components.Acceleration{A: 0.1, DA: 0.05},
			Sprite:       &components.Sprite{Image: playerImg},
			Animation:    entities.LoadPlayerAnimations(),
		},
	}

	// Initialize the game
	game := &Game{entities: []entities.Entity{player}}

	// Start the game loop
	ebiten.SetWindowSize(1280, 960)   // Window size
	ebiten.SetWindowTitle("Astroids") // Window title
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
