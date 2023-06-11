package main

import (
	// "fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenX		= 640
	screenY		= 480
)

type Game struct {
	score		int
}

func NewGame() *Game {
	g := &Game{}
	return g
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return screenX, screenY
}

func main() {
	ebiten.SetWindowSize(screenX, screenY)
	ebiten.SetWindowTitle("TEST HOGE")
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}