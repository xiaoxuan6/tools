package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
	"github.com/xiaoxuan6/tools/bookmarks"
	"github.com/xiaoxuan6/tools/clipboard2img"
	"github.com/xiaoxuan6/tools/code"
	"github.com/xiaoxuan6/tools/host"
	"github.com/xiaoxuan6/tools/ocr"
	"github.com/xiaoxuan6/tools/qrcode"
	"github.com/xiaoxuan6/tools/translation"
	"os"
)

var Version string

func main() {
	app := cli.App{
		Name:  "tools",
		Usage: "tools",
		Commands: []*cli.Command{
			{
				Name:    "translation",
				Usage:   "translation 翻译",
				Aliases: []string{"t"},
				Flags:   translation.Flags,
				Action:  translation.Action,
			},
			{
				Name:    "qrcode",
				Usage:   "qrcode 二维码生成、解析",
				Aliases: []string{"q"},
				Flags:   qrcode.Flags,
				Action:  qrcode.Action,
			},
			{
				Name:    "ocr",
				Usage:   "ocr 图片识别文字",
				Aliases: []string{"o"},
				Flags:   ocr.Flags,
				Action:  ocr.Action,
			},
			{
				Name:    "clipboard2img",
				Usage:   "clipboard2img 粘贴板图片保存到本地",
				Aliases: []string{"c2i"},
				Action:  clipboard2img.Action,
			},
			{
				Name:    "version",
				Usage:   "version 版本号",
				Aliases: []string{"v"},
				Action: func(c *cli.Context) error {
					fmt.Println("tools version:", color.GreenString(Version))
					return nil
				},
			},
			{
				Name:    "host",
				Usage:   "host 文件操作",
				Aliases: []string{"h"},
				Flags:   host.Flags,
				Action:  host.Action,
			},
			{
				Name:    "bookmarks",
				Usage:   "bookmarks 将书签导出到文件中",
				Aliases: []string{"b"},
				Flags:   bookmarks.Flags,
				Action:  bookmarks.Action,
			},
			{
				Name:    "code",
				Usage:   "code 人工智能来解释您不理解的任何代码的工具",
				Aliases: []string{"c"},
				Action:  code.Action,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
}
