package main

import (
	"image/color"
	"log"
	"math/rand/v2"

	"feudal/internal/camera"
	"feudal/internal/components"
	"feudal/internal/entities"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	ScreenWidth  = 1280
	ScreenHeight = 960
	CameraWidth  = 320
	CameraHeight = 240
	WorldWidth   = 1000 // Large world for testing camera movement
	WorldHeight  = 1000
)

// Dot struct - Just a static point
type Dot struct {
	X, Y float64
}

// Game struct contains - Needed for life cycle
type Game struct {
	entities []entities.Entity
	player   *entities.Player

	camera *camera.Camera

	dots []Dot
}

// Update handles game logic
func (g *Game) Update() error {
	for i := range g.entities {
		g.entities[i].Update()
	}

	g.camera.Update()

	return nil
}

// Draw renders the game screen
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 255}) // Clear the screen with black

	// Draw dots with camera offset
	for _, dot := range g.dots {
		x := dot.X - g.camera.X
		y := dot.Y - g.camera.Y
		ebitenutil.DrawRect(screen, x, y, 3, 3, color.White)
	}

	for i := range g.entities {
		g.entities[i].Draw(screen, g.camera.X, g.camera.Y)
	}
}

// Layout specifies the game's screen size
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return CameraWidth, CameraHeight
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

	// Generate random dots
	dots := make([]Dot, 50)
	for i := range dots {
		dots[i] = Dot{
			X: rand.Float64() * WorldWidth,
			Y: rand.Float64() * WorldHeight,
		}
	}

	camera := &camera.Camera{
		X:               0,
		Y:               0,
		FollowingPlayer: player,
		ScreenWidth:     CameraWidth,
		ScreenHeight:    CameraHeight,
		WorldWidth:      WorldWidth,
		WorldHeight:     WorldHeight,
	}

	// Initialize the game
	game := &Game{
		entities: []entities.Entity{player},
		player:   player, // Store reference to player
		camera:   camera,
		dots:     dots,
	}

	// Start the game loop
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight) // Window size
	ebiten.SetWindowTitle("Feudal")                 // Window title
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
