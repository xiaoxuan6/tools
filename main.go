package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"githug.com/xiaoxuan6/tools/translation"
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
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
}
