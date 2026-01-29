package main

import (
	//"image/color"
	"log"
	"strconv"

	"go-fishing/collision"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	//"github.com/hajimehoshi/ebiten/v2/text"
	//"golang.org/x/image/font"
)

type Game struct{}

const (
	gameScreenSize int = 512
	gameWindowSize int = 720
)

var (
	bg           *ebiten.Image
	rollSpeed    int = 97
	bgTranslateY int = 0

	bgLength int = 5120

	hookMaxLength int = -bgLength + gameScreenSize
	hookCurrentLength int = -gameScreenSize + 100
	hookSpeed int = 5
)

// Gameloop
func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyS){
		if bgTranslateY > (hookCurrentLength) {
		bgTranslateY -= hookSpeed
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyW){
		if bgTranslateY != (0) {
		bgTranslateY += hookSpeed
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
	return gameScreenSize, gameScreenSize
}

func main() {
	ebiten.SetWindowSize(gameWindowSize, gameWindowSize)
	ebiten.SetWindowTitle("Go Fishing")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
	collision.Test()

}