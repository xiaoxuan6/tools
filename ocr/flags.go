package ocr

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/xiaoxuan6/tools/common"
	"os"
	"path/filepath"
)

func Flags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringFlag{
			Name:    "filename",
			Aliases: []string{"f"},
			Usage:   "图片文件名",
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
	}

	return flags
}
