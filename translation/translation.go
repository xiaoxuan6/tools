package translation

import (
	"errors"
	"fmt"
	"github.com/abadojack/whatlanggo"
	"github.com/atotto/clipboard"
	"github.com/avast/retry-go"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
	"github.com/xiaoxuan6/deeplx"
	"github.com/xiaoxuan6/tools/common"
	"strings"
)

var Command = &cli.Command{
	Name:    "translation",
	Usage:   "翻译",
	Aliases: []string{"t"},
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "content",
			Aliases: []string{"c"},
			Usage:   "translation content",
		},
	},
	Action: Action,
}

func Action(c *cli.Context) error {
	content := c.String("content")
	if len(content) < 1 {
		content = common.ScannerF(func() string {
			body, _ := clipboard.ReadAll()
			return strings.TrimSpace(body)
		})
	}

	if len(content) < 1 {
		fmt.Println(color.RedString("翻译内容不能为空"))
		return nil
	}

	fmt.Println(color.RedString("翻译内容："), content)
	common.Start("translations ")

	info := whatlanggo.Detect(content)
	sourceLang := info.Lang.String()
	var targetLang string
	switch sourceLang {
	case "Mandarin", "Chinese", "Zh", "zh-cn":
		targetLang = "en"
	case "English", "En":
		targetLang = "zh"
	default:
		targetLang = ""
	}

	var result string
	err := retry.Do(
		func() error {
			response := deeplx.Translate(content, sourceLang, targetLang)
			if response.Code != 200 {
				return errors.New(response.Msg)
			}

			result = response.Data
			return nil
		},
		retry.Attempts(3),
		retry.LastErrorOnly(true),
	)
	common.Stop()

	if err != nil {
		result = "翻译失败，翻译原文 => " + content
	}

	fmt.Println(color.GreenString("翻译结果："), strings.TrimSpace(result))
	return nil
}
