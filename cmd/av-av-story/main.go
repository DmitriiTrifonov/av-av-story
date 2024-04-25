package main

import (
	"github.com/DmitriiTrifonov/av-av-story/internal/app"
	"log"
)

func main() {
	application := app.New(nil)
	if err := application.Run(); err != nil {
		log.Fatal("cannot start application")
	}
}
