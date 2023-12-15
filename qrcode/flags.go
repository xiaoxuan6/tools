package qrcode

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/xiaoxuan6/tools/common"
	"net/url"
	"os"
	"path/filepath"
)

var Flags = []cli.Flag{
	&cli.StringFlag{
		Name:    "content",
		Aliases: []string{"c"},
		Usage:   "qrcode content",
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
			if common.Identify(ext) == false {
				return fmt.Errorf(common.ExtError)
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
			if common.Identify(ext) == false {
				return fmt.Errorf(common.ExtError)
			}

			return nil
		},
	},
}
