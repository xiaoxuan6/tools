package bookmarks

import "github.com/urfave/cli/v2"

var Flags = []cli.Flag{
	&cli.StringFlag{
		Name:     "type",
		Aliases:  []string{"t"},
		Usage:    "导出书签类型，chrome 或者 firefox",
		Required: true,
	},
}
