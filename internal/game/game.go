package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"log"
)

type Game struct {
	cfg   *Config
	image *ebiten.Image
}

type Config struct {
	ScreenWidth  int
	ScreenHeight int
}

func New(cfg *Config) *Game {
	img, _, err := ebitenutil.NewImageFromFile("assets/images/av-av.png")
	if err != nil {
		log.Fatal(err)
	}
	return &Game{cfg: cfg, image: img}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{
		R: 0xA0,
		G: 0x00,
		B: 0xA0,
		A: 0xFF,
	})
	screen.DrawImage(g.image, nil)
	ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.cfg.ScreenWidth, g.cfg.ScreenHeight
}
