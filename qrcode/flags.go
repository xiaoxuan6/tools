package qrcode

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
	"net/url"
	"os"
	"path/filepath"
)

var ExtError = color.RedString("文件格式错误，请重新输入，支持的格式有：%s", "jpg, jpeg, png")

func Flags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringFlag{
			Name:    "content",
			Aliases: []string{"c"},
			Usage:   "qrcode content",
		},
		&cli.BoolFlag{
			Name:    "stdin",
			Aliases: []string{"s"},
			Usage:   "read from stdin",
		},
		&cli.StringFlag{
			Name:    "filename",
			Aliases: []string{"f"},
			Usage:   "qrcode filename",
			Action: func(context *cli.Context, s string) error {
				if _, err := os.Stat(s); os.IsNotExist(err) {
					return fmt.Errorf("文件 %s 不存在，请重新选择文件", s)
				}

				ext := filepath.Ext(s)
				if identify(ext) == false {
					return fmt.Errorf(ExtError)
				}

				return nil
			},
		},
		&cli.StringFlag{
			Name:    "url",
			Aliases: []string{"u"},
			Usage:   "qrcode url",
			Action: func(context *cli.Context, s string) error {
				u, err := url.Parse(s)
				if err != nil {
					return fmt.Errorf("url %s 解析失败，请重新输入", s)
				}

				ext := filepath.Ext(u.Path)
				if identify(ext) == false {
					return fmt.Errorf(ExtError)
				}

				return nil
			},
		},
	}

	return flags
}

func identify(ext string) bool {
	stat := false
	extAllow := []string{"jpg", "jpeg", "png"}
	for _, v := range extAllow {
		if ext == v {
			stat = true
		}
	}

	return stat
}
