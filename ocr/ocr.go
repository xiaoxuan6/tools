package ocr

import (
	"bytes"
	"fmt"
	"github.com/fatih/color"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v2"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
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

	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	fileWriter, err := bodyWriter.CreateFormFile("file", filename)
	if err != nil {
		return fmt.Errorf("创建文件失败")
	}

	_, _ = io.Copy(fileWriter, file)
	contentType := bodyWriter.FormDataContentType()
	_ = bodyWriter.Close()

	response, err := http.Post("https://api.toolnb.com/api/ocr.html", contentType, bodyBuf)
	if err != nil {
		return fmt.Errorf("请求失败，请重新输入")
	}
	defer response.Body.Close()

	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("读取响应失败")
	}
	result := gjson.ParseBytes(b)
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
