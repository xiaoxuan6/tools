package qrcode

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"net/http"
	"os"
)

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
		b, err := ioutil.ReadFile(filename)
		if err != nil {
			return fmt.Errorf(color.RedString("读取文件失败，请重新输入"))
		}

		scan(b)
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

		scan(b)
		return nil
	}

	return nil
}