package code

import (
	"bytes"
	"fmt"
	"github.com/OwO-Network/gdeeplx"
	"github.com/atotto/clipboard"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

var Command = &cli.Command{
	Name:    "code",
	Usage:   "code 人工智能来解释您不理解的任何代码的工具",
	Aliases: []string{"c"},
	Action:  Action,
}

func Action(c *cli.Context) error {
	text, err := clipboard.ReadAll()
	if err != nil {
		return fmt.Errorf(color.RedString("请复制您需要解释的代码，再执行该操作"))
	}

	fmt.Println(color.RedString("正文：%s", text))
	text = strings.ReplaceAll(strings.TrimSpace(text), `"`, `\"`)
	re := regexp.MustCompile(`(\r\n|\r|\n)`)
	text = re.ReplaceAllString(text, "\\n\\t")
	// 匹配换行符后的空格，然后替换为"\n\t"
	re = regexp.MustCompile(`\\n\\t[ \t]*`)
	text = re.ReplaceAllString(text, "\\n\\t")

	if len(text) < 1 {
		return fmt.Errorf(color.RedString("请复制您需要解释的代码，再执行该操作"))
	}

	var body bytes.Buffer
	body.WriteString(fmt.Sprintf(`{"code":"%s"}`, text))
	res, err := http.Post("https://whatdoesthiscodedo.com/api/stream-text", "application/json", &body)
	if err != nil {
		return fmt.Errorf(color.RedString("请求失败：%s", err.Error()))
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf(color.RedString("请求失败：%s", res.Status))
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf(color.RedString("读取响应失败：%s", err.Error()))
	}

	response := strings.Split(string(b), "\n")
	targetResponse, err := gdeeplx.Translate(response[0], "", "zh", 0)
	if err != nil {
		return fmt.Errorf(color.RedString("翻译失败：%s", err.Error()))
	}
	fmt.Println(color.GreenString("解释：%s", strings.TrimSpace(targetResponse.(map[string]interface{})["data"].(string))))
	return nil
}
