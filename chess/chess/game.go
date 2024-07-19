package chess

import (
	
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten"
	// "github.com/hajimehoshi/ebiten/audio/wav"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Game 象棋窗口
type Game struct {
	sqSelected     int                   // 选中的格子
	mvLast         int                   // 上一步棋
	bFlipped       bool                  //是否翻转棋盘
	images         map[int]*ebiten.Image //图片资源
	audios         map[int]*audio.Player //音效
	audioContext   *audio.Context        //音效器
	singlePosition *PositionStruct       //棋局单例
}

// Update 更新状态，1秒60帧
func (g *Game) Update(screen *ebiten.Image) error {
	img, _, err := ebitenutil.NewImageFromFile("./res/ChessBoard.png")
	if err != nil {
		log.Print(err)
		return err
	}
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(img, op)
	return nil
}

// Layout 布局采用外部尺寸（例如，窗口尺寸）并返回（逻辑）屏幕尺寸，如果不使用外部尺寸，只需返回固定尺寸即可。
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 520, 576
}

// NewGame 创建象棋程序
func NewGame() bool {
	game := &Game{}
	if game == nil {
		return false
	}

	ebiten.SetWindowSize(520, 576)
	ebiten.SetWindowTitle("中国象棋")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
		return false
	}

	return true
}
//loadResource 加载资源
// func (g *Game) loadResource() bool {
// 	for k, v := range resMap {
// 		if k >= MusicSelect {
// 			//加载音效
// 			d, err := wav.Decode(g.audioContext, audio.BytesReadSeekCloser(v))
// 			if err != nil {
// 				fmt.Print(err)
// 				return false
// 			}
// 			player, err := audio.NewPlayer(g.audioContext, d)
// 			if err != nil {
// 				fmt.Print(err)
// 				return false
// 			}
// 			g.audios[k] = player
// 		} else {
// 			//加载图片
// 			img, _, err := image.Decode(bytes.NewReader(v))
// 			if err != nil {
// 				fmt.Print(err)
// 				return false
// 			}
// 			ebitenImage, _ := ebiten.NewImageFromImage(img, ebiten.FilterDefault)
// 			g.images[k] = ebitenImage
// 		}
// 	}

// 	return true
// }

// //playAudio 播放音效
// func (g *Game) playAudio(value int) {
// 	if player, ok := g.audios[value]; ok {
// 		player.Rewind()
// 		player.Play()
// 	}
// }