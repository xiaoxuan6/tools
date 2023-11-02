package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
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
				Usage:   "translation",
				Aliases: []string{"t"},
				Flags:   translation.Flags(),
				Action:  translation.Action,
			},
			{
				Name:    "qrcode",
				Usage:   "qrcode",
				Aliases: []string{"q"},
				Flags:   qrcode.Flags(),
				Action:  qrcode.Action,
			},
			{
				Name:    "ocr",
				Usage:   "ocr",
				Aliases: []string{"o"},
				Flags:   ocr.Flags(),
				Action:  ocr.Action,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
}
