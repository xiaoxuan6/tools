package qrcode

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"net/http"
	"os"
)

var Command = &cli.Command{
	Name:    "qrcode",
	Usage:   "二维码生成、解析",
	Aliases: []string{"q"},
	Flags:   Flags,
	Action:  Action,
}

func Action(c *cli.Context) error {
	content := c.String("content")
	stdin := c.Bool("stdin")

	if len(content) < 1 && stdin {
		b, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return fmt.Errorf(color.RedString("读取内容失败，请重新输入"))
		}
		content = string(b)
	}

	if len(content) > 0 {
		generateQrcode(content)
		return nil
	}

	filename := c.String("filename")
	if filename != "" {
		f, err := os.Open(filename)
		if err != nil {
			return fmt.Errorf(color.RedString("文件 %s 不存在，请重新选择文件", filename))
		}

		defer func() {
			_ = f.Close()
		}()

		scan(f)
		return nil
	}

	uri := c.String("url")
	if uri != "" {
		response, err := http.Get(uri)
		if err != nil {
			return fmt.Errorf(color.RedString("请求失败，请重新输入"))
		}

		b, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return fmt.Errorf(color.RedString("读取内容失败，请重新输入"))
		}

		f, _ := os.Create("qrcode.png")
		_, _ = f.Write(b)
		defer func() {
			_ = f.Close()
		}()

		scan(f)
		return nil
	}

	fmt.Println(color.RedString("参数错误，请重新输入: qrcode --help/q -h"))
	return nil
}
