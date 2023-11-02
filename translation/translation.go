package translation

import (
	"fmt"
	"github.com/OwO-Network/gdeeplx"
	"github.com/abadojack/whatlanggo"
	"github.com/atotto/clipboard"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

func Action(c *cli.Context) error {
	content := c.String("content")
	if content == "" {
		content = setContent()
	}

	info := whatlanggo.Detect(content)
	lang := info.Lang.String()
	language := setLanguage(lang)
	response, err := gdeeplx.Translate(content, lang, language, 0)
	if err != nil {
		return err
	}

	fmt.Println(color.GreenString("翻译结果："), response.(map[string]interface{})["data"])
	return nil
}

func setContent() string {
	content, _ := clipboard.ReadAll()
	fmt.Println(color.RedString("翻译内容："), content)
	return content
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
