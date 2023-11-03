package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/xiaoxuan6/tools/clipboard2img"
	"github.com/xiaoxuan6/tools/ocr"
	"github.com/xiaoxuan6/tools/qrcode"
	"github.com/xiaoxuan6/tools/translation"
	"os"
)

func main() {
	app := cli.App{
		Name:  "tools",
		Usage: "tools",
		Commands: []*cli.Command{
			{
				Name:    "translation",
				Usage:   "translation 翻译",
				Aliases: []string{"t"},
				Flags:   translation.Flags(),
				Action:  translation.Action,
			},
			{
				Name:    "qrcode",
				Usage:   "qrcode 二维码生成、解析",
				Aliases: []string{"q"},
				Flags:   qrcode.Flags(),
				Action:  qrcode.Action,
			},
			{
				Name:    "ocr",
				Usage:   "ocr 图片识别文字",
				Aliases: []string{"o"},
				Flags:   ocr.Flags(),
				Action:  ocr.Action,
			},
			{
				Name:    "clipboard2img",
				Usage:   "clipboard2img 粘贴板图片保存到本地",
				Aliases: []string{"c2i"},
				Action:  clipboard2img.Action,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
}
