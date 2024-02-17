package host

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var (
	path = filepath.Join(os.Getenv("windir"), "System32\\drivers\\etc\\hosts")

	Command = &cli.Command{
		Name:    "localhost",
		Usage:   "host 文件操作",
		Aliases: []string{"l"},
		Flags: []cli.Flag{
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
		},
		Action: Action,
	}
)

func Action(c *cli.Context) error {
	if runtime.GOOS != "Windows" {
		return fmt.Errorf(color.RedString("当前系统不支持该命令, 仅支持 Windows 系统"))
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		_ = os.MkdirAll(path, os.ModePerm)
	}

	b, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf(color.RedString("读取文件 %s 失败，请检查文件是否存在", path))
	}

	if strings.Contains(string(b), c.String("domain")) {
		return fmt.Errorf(color.RedString("域名 %s 已存在，请重新输入", c.String("domain")))
	}

	content := fmt.Sprintf("%s127.0.0.1 %s\n", string(b), c.String("domain"))
	_ = os.WriteFile(path, []byte(content), os.ModePerm)
	fmt.Println(color.GreenString("域名 %s 添加成功", c.String("domain")))

	return nil
}
