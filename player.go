package main

import (
  // "fmt"
  // "math/rand"
  "log"
  "github.com/hajimehoshi/ebiten/v2"
  "github.com/hajimehoshi/ebiten/v2/ebitenutil"
  // "github.com/hajimehoshi/ebiten/v2/vector"
  // "github.com/hajimehoshi/ebiten/v2/inpututil"
  // "github.com/hajimehoshi/ebiten/v2/text"
  // "image"
  // "image/color"
  _ "image/png"
)

type player struct {
  name string
	path string
	x float64
	y float64
}

func (p *player) set_coordinates(x float64, y float64) {
	p.x = x
	p.y = y
}

func (p *player) draw_image(screen *ebiten.Image) {
  var err error
  var img *ebiten.Image 

  img, _, err = ebitenutil.NewImageFromFile(p.path)
  if err != nil {
    log.Fatal(err)
  }

  new_x := p.x - float64(img.Bounds().Dx() / 2)
  new_y := p.y - float64(img.Bounds().Dy())

  op := &ebiten.DrawImageOptions{}
  op.GeoM.Translate(new_x, new_y)

  screen.DrawImage(img, op)
}