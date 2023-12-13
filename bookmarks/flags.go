package bookmarks

import "github.com/urfave/cli/v2"

var Flags = []cli.Flag{
	&cli.StringFlag{
		Name:     "browser",
		Aliases:  []string{"b"},
		Required: true,
		Usage:    "支持浏览器 chrome、edge 或 firefox",
	},
}
