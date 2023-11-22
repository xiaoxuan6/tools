package host

import (
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
	"net/url"
)

func Flags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:     "domain",
			Aliases:  []string{"d"},
			Usage:    "域名",
			Required: true,
			Action: func(context *cli.Context, s string) error {
				_, err := url.Parse(s)
				if err != nil {
					return errors.New(fmt.Sprintf("domain %s 解析失败，请重新输入", s))
				}
				return nil
			},
		},
	}
}
