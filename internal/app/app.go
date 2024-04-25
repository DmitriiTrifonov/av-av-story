package app

import (
	"fmt"
	"github.com/DmitriiTrifonov/av-av-story/internal/game"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	defaultConfig = &Config{
		AppTitle: "av-av-story",
		Window: Window{
			Width:  1280,
			Height: 720,
		},
		ScaleFactor: 1,
	}
)

type App struct {
	cfg *Config
}

type Config struct {
	AppTitle    string
	Window      Window
	ScaleFactor int
}

type Window struct {
	Height int
	Width  int
}

func New(cfg *Config) *App {
	a := &App{}
	if cfg == nil {
		a.cfg = defaultConfig
	}

	return a
}

func (a *App) Run() error {
	fmt.Println("app started")

	ebiten.SetWindowSize(a.cfg.Window.Width, a.cfg.Window.Height)
	ebiten.SetWindowTitle(a.cfg.AppTitle)
	if err := ebiten.RunGame(game.New(&game.Config{
		ScreenHeight: a.cfg.Window.Height / a.cfg.ScaleFactor,
		ScreenWidth:  a.cfg.Window.Width / a.cfg.ScaleFactor})); err != nil {

		return err
	}

	return nil
}
