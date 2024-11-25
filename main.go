package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
	"github.com/xiaoxuan6/tools/bookmarks"
	"github.com/xiaoxuan6/tools/clipboard2img"
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
			bookmarks.Command,
			clipboard2img.Command,
			host.Command,
			translation.Command,
			qrcode.Command,
			ocr.Command,
			{
				Name:    "version",
				Usage:   "版本号",
				Aliases: []string{"v"},
				Action: func(c *cli.Context) error {
					fmt.Println("tools version:", color.GreenString(Version))
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
}
