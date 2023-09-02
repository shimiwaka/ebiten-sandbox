package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	return
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 100, 200
}

func main() {
	ebiten.SetWindowTitle("ほぎゃあああああああ")
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
