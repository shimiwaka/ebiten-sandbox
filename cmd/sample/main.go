package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"log"
	"image"
	_ "image/png"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	// "github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const DefaultTPS = 15

const (
	screenX		= 640
	screenY		= 480
	fontSize	= 16
)

var (
	sampleImg *ebiten.Image
)

//go:embed resource/nandakore.png
var byteSampleImg []byte

type Game struct {
	score		int
	objectX		float64
	objectY		float64
}

func NewGame() *Game {
	g := &Game{}
	g.objectX = 20
	g.objectY = 20
	return g
}

func (g *Game) Update() error {
	if g.isKeyPressed() {
		g.objectX += 1
	} else {
		g.objectX -= 1
	}
	return nil
}

func (g *Game) isKeyPressed() bool {
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		return true
	}
	return false
}

func (g *Game) Draw(screen *ebiten.Image) {
	img, _, err := image.Decode(bytes.NewReader(byteSampleImg))
	if err != nil {
		fmt.Println("IMAGE LOADING ERROR")
		log.Fatal(err)
	}
	sampleImg = ebiten.NewImageFromImage(img)

	tt, err := opentype.Parse(fonts.PressStart2P_ttf)
	if err != nil {
		log.Fatal(err)
	}
	const dpi = 72
	arcadeFont, _ := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    fontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

	screen.Fill(color.White)
	text.Draw(screen, fmt.Sprintf("TESUYA"), arcadeFont, 300, 20, color.Black)
	
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.objectX, g.objectY)
	op.Filter = ebiten.FilterLinear
	screen.DrawImage(sampleImg, op)
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