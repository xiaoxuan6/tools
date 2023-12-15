package translation

import (
	"github.com/urfave/cli/v2"
)

var Flags = []cli.Flag{
	&cli.StringFlag{
		Name:    "content",
		Aliases: []string{"c"},
		Usage:   "translation content",
	},
}
