package host

import (
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var path = filepath.Join(os.Getenv("windir"), "System32\\drivers\\etc\\hosts")

func Action(c *cli.Context) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		_ = os.MkdirAll(path, os.ModePerm)
	}

	b, err := ioutil.ReadFile(path)
	if err != nil {
		return errors.New(fmt.Sprintf("读取文件 %s 失败，请检查文件是否存在", path))
	}

	if strings.Contains(string(b), c.String("domain")) {
		return errors.New(fmt.Sprintf("域名 %s 已存在，请重新输入", c.String("domain")))
	}

	content := fmt.Sprintf("%s \n127.0.0.1 %s\n", string(b), c.String("domain"))
	_ = os.WriteFile(path, []byte(content), os.ModePerm)
	fmt.Println(fmt.Sprintf("域名 %s 添加成功", c.String("domain")))

	return nil
}
