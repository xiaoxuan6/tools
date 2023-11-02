package translation

import (
	"fmt"
	"github.com/OwO-Network/gdeeplx"
	"github.com/abadojack/whatlanggo"
	"github.com/atotto/clipboard"
	"github.com/fatih/color"
	"github.com/sqweek/dialog"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"os"
	"strings"
)

func Action(c *cli.Context) error {
	var content string
	filename := c.String("filename")
	if filename != "" {
		if _, err := os.Stat(filename); os.IsNotExist(err) {
		SetFilename:
			fmt.Println(color.RedString("文件不存在，请重新选择文件"))
			filename = setFilename()

			if filename == "" {
				goto SetFilename
			}

			f, _ := os.Stat(filename)
			fileSize := f.Size()
			fileSizeInMB := float64(fileSize) / (1024 * 1024)
			if fileSizeInMB > 1 {
				return fmt.Errorf("文件大小超过 1MB，无法翻译")
			}

			content = fileGetContent(filename)
		}
	} else {
		content = c.String("content")
		if len(content) < 1 {
			stdin := c.Bool("stdin")
			content = setContent(stdin)
		}
	}

	info := whatlanggo.Detect(content)
	lang := info.Lang.String()
	language := setLanguage(lang)
	response, err := gdeeplx.Translate(content, lang, language, 0)
	if err != nil {
		return err
	}

	fmt.Println(color.GreenString("翻译结果："), strings.TrimSpace(response.(map[string]interface{})["data"].(string)))
	return nil
}

func setFilename() string {
	filename, err := dialog.File().Filter("txt file", "txt").Title("选择文件").Load()
	if err != nil {
		return ""
	}

	return filename
}

func fileGetContent(filename string) string {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		return ""
	}

	fmt.Println(color.RedString("翻译文件："), filename)
	return string(f)
}

func setContent(stdin bool) string {
	if stdin {
		b, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return ""
		}
		return string(b)
	}

	content, _ := clipboard.ReadAll()
	fmt.Println(color.RedString("翻译内容："), strings.TrimSpace(content))
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
