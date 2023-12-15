package translation

import (
	"fmt"
	"github.com/OwO-Network/gdeeplx"
	"github.com/abadojack/whatlanggo"
	"github.com/atotto/clipboard"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
	"github.com/xiaoxuan6/tools/common"
	"strings"
)

var Command = &cli.Command{
	Name:    "translation",
	Usage:   "翻译",
	Aliases: []string{"t"},
	Flags:   Flags,
	Action:  Action,
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

	common.Start("translations ")
	fmt.Println(color.RedString("翻译内容："), content)
	info := whatlanggo.Detect(content)
	lang := info.Lang.String()
	language := setLanguage(lang)
	response, err := gdeeplx.Translate(content, lang, language, 0)
	common.Stop()
	if err != nil {
		return err
	}

	fmt.Println(color.GreenString("翻译结果："), strings.TrimSpace(response.(map[string]interface{})["data"].(string)))
	return nil
}

func setLanguage(language string) string {
	languages := map[string]string{
		"English":  "zh",
		"Mandarin": "en",
	}

	if _, ok := languages[language]; ok {
		return languages[language]
	}

	return "zh"
}
