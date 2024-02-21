package translation

import (
	"fmt"
	"github.com/OwO-Network/gdeeplx"
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

	info := whatlanggo.Detect(content)
	var targetLang string
	switch info.Lang.String() {
	case "Mandarin", "Chinese", "Zh", "zh-cn":
		targetLang = "en"
	case "English", "En":
		targetLang = "zh"
	default:
		targetLang = ""
	}

	response, err := deeplx.Translate(content, "", targetLang)
	num := 0
RETRY:
	if err != nil {
		result, errs := gdeeplx.Translate(content, "", targetLang, 0)
		response = strings.TrimSpace(result.(map[string]interface{})["data"].(string))
		err = errs

		num++
		if num < 3 {
			goto RETRY
		}

		response = err.Error()
	}
	common.Stop()

	fmt.Println(color.GreenString("翻译结果："), strings.TrimSpace(response))
	return nil
}
