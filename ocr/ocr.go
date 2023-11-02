package ocr

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v2"
	"github.com/xiaoxuan6/tools/common"
	"os"
	"path/filepath"
	"strings"
)

func Action(c *cli.Context) error {
	filename := c.String("filename")
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("文件 %s 不存在，请重新选择文件", filename)
	}
	defer file.Close()

	if strings.HasPrefix(filename, ".") {
		dir, _ := os.Getwd()
		filename = filepath.Join(dir, filename)
	}

	response, err := common.PostWithMultipart("https://api.toolnb.com/api/ocr.html", "file", filename, file)
	if err != nil {
		return fmt.Errorf(color.RedString(err.Error()))
	}

	result := gjson.ParseBytes(response)
	if result.Get("code").Int() != 1 {
		return fmt.Errorf(result.Get("msg").String())
	}

	fmt.Println(color.GreenString("识别结果："))
	result.Get("data.list").ForEach(func(key, value gjson.Result) bool {
		fmt.Println(value.Get("text").String())
		return true
	})

	return nil
}
