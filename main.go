package main

import (
	//"image/color"
	"log"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	//"github.com/hajimehoshi/ebiten/v2/text"
	//"golang.org/x/image/font"
)

type Game struct{}

var (
	bg           *ebiten.Image
	rollSpeed    int = 97
	bgTranslateY int = 0
)

// Gameloop
func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyS){
		if bgTranslateY > (-5120 + 512) {
		bgTranslateY -= 10
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyW){
		if bgTranslateY != (0) {
		bgTranslateY += 10
		}
	}

	return nil
}

func init() {
	var err error
	bg, _, err = ebitenutil.NewImageFromFile("assets/bg.png")
	if err != nil {
		log.Fatal(err)
	}
}

// Render to screen
func (g *Game) Draw(screen *ebiten.Image) {
	bgop := &ebiten.DrawImageOptions{}
	bgop.GeoM.Translate(0, float64(bgTranslateY))
	screen.DrawImage(bg, bgop)
	//text.Draw(screen, "100 METERS", font.StretchCondensed, 50, 100, color.White)

	ebitenutil.DebugPrint(screen, strconv.Itoa((bgTranslateY * -1) / 8))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 512, 512
}

func main() {
	ebiten.SetWindowSize(720, 720)
	ebiten.SetWindowTitle("Go Fishing")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}

}