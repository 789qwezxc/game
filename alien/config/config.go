package config

import (
	"encoding/json"
	"image/color"
	"log"
	"os"
)

type Config struct {
	ScreenWidth        int        `json:"screen_width"`         // 屏幕宽度
	ScreenHeight       int        `json:"screen_height"`        // 屏幕高度
	Title              string     `json:"title"`                // 标题
	BgColor            color.RGBA `json:"bg_color"`             // 背景颜色
	MoveSpeed          int        `json:"move_speed"`           // 移动速度
	BulletWidth        int        `json:"bullet_width"`         // 子弹宽度
	BulletHeight       int        `json:"bullet_height"`        // 子弹高度
	BulletSpeed        int        `json:"bullet_speed"`         // 子弹速度
	BulletColor        color.RGBA `json:"bullet_color"`         // 子弹颜色
	MaxBulletNum       int        `json:"max_bullet_num"`       //页面中最多子弹数量
	BulletInterval     int64      `json:"bullet_interval"`      //发射子弹间隔时间
	MonsterSpeedFactor int        `json:"monster_speed_factor"` // 怪物速度因子
	TitleFontSize      int        `json:"title_font_size"`      //标题字体大小
	FontSize           int        `json:"font_size"`            //字体大小
	SmallFontSize      int        `json:"small_font_size"`      //小字体大小
	FailedCountLimit   int        `json:"failed_count_limit"`   //最多能遗漏多少怪物
}

func LoadConfig() *Config {
	file, err := os.Open("./config.json")
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer file.Close()
	config := new(Config)
	err = json.NewDecoder(file).Decode(config)
	if err != nil {
		log.Fatalf("%v", err)
	}
	return config
}
