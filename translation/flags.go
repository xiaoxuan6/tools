package translation

import (
	"github.com/urfave/cli/v2"
)

func Flags() (flags []cli.Flag) {
	flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "content",
			Aliases: []string{"c"},
			Usage:   "translation content",
		},
		&cli.StringFlag{
			Name:    "filename",
			Aliases: []string{"f"},
			Usage:   "translation filename",
		},
	}
	return
}
