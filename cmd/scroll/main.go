// Copyright 2018 The Ebiten Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type Game struct {
	nowPage   int
	pages     []*Page
	stroke    *Stroke
	offset    int
	tmpOffset int
}

type Content struct {
	value string
	x     int
	y     int
}

type Page struct {
	contents []*Content
}

const (
	screenX = 640
	screenY = 480
)

type StrokeSource interface {
	Position() (int, int)
	IsJustReleased() bool
}

type MouseStrokeSource struct{}

func (m *MouseStrokeSource) Position() (int, int) {
	return ebiten.CursorPosition()
}

func (m *MouseStrokeSource) IsJustReleased() bool {
	return inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft)
}

type Stroke struct {
	source StrokeSource

	initX int
	initY int

	currentX int
	currentY int

	released bool
}

func (s *Stroke) Update() {
	if s.source.IsJustReleased() {
		s.released = true
	} else {
		s.released = false
	}
	s.currentX, s.currentY = s.source.Position()
}

func (s *Stroke) IsReleased() bool {
	return s.released
}

func NewStroke(source StrokeSource) *Stroke {
	cx, cy := source.Position()
	return &Stroke{
		source:   source,
		initX:    cx,
		initY:    cy,
		currentX: cx,
		currentY: cy,
	}
}

var (
	normalFont font.Face
)

func NewGame() *Game {
	page := &Page{}
	ct1 := &Content{
		value: "てすと１てすと１てすと１てすと１てすと１てすと１",
		x:     0,
		y:     25,
	}
	ct2 := &Content{
		value: "こんにちはありがとうさよならまた会いましょう",
		x:     0,
		y:     50,
	}
	page.contents = append(page.contents, ct1)
	page.contents = append(page.contents, ct2)

	game := &Game{
		nowPage: 0,
		stroke:  nil,
	}
	game.pages = append(game.pages, page)

	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	normalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    16,
		DPI:     dpi,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		log.Fatal(err)
	}

	return game
}

func (g *Game) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		s := NewStroke(&MouseStrokeSource{})
		g.stroke = s
	}

	if g.stroke != nil {
		g.stroke.Update()
		g.tmpOffset = g.stroke.currentY - g.stroke.initY
		if g.stroke.released {
			g.stroke = nil
			g.offset = g.offset + g.tmpOffset
			if g.offset > 0 {
				g.offset = 0
			}
			g.tmpOffset = 0
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)

	for _, v := range g.pages[g.nowPage].contents {
		text.Draw(screen, v.value, normalFont, 0, v.y+g.offset+g.tmpOffset, color.Black)
	}

	return
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenX, screenY
}

func main() {
	ebiten.SetWindowTitle("Sample Game")
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
