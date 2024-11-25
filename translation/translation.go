package translation

import (
	"fmt"
	"github.com/abadojack/whatlanggo"
	"github.com/atotto/clipboard"
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

	lang := whatlanggo.DetectLang(content)
	sourceLang := strings.ToLower(lang.Iso6391())
	var targetLang string
	switch sourceLang {
	case "zh":
		targetLang = "en"
	case "so":
		targetLang = "zh"
	default:
		targetLang = ""
	}

	response := deeplx.Translate(content, sourceLang, targetLang)
	common.Stop()

	result := response.Data
	if response.Code != 200 {
		result = "翻译失败 => " + response.Msg
	}

	fmt.Println(color.GreenString("翻译结果："), strings.TrimSpace(result))
	return nil
}
