package clipboard2img

import (
	"bytes"
	"fmt"
	"github.com/fatih/color"
	"github.com/mitchellh/go-homedir"
	"github.com/urfave/cli/v2"
	"golang.design/x/clipboard"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

func Action(c *cli.Context) error {
	if runtime.GOOS != "Windows" {
		return fmt.Errorf(color.RedString("当前系统不支持该命令, 仅支持 Windows 系统"))
	}

	err := clipboard.Init()
	if err != nil {
		return fmt.Errorf(color.RedString("clipboard init error: %w", err))
	}

	b := clipboard.Read(clipboard.FmtImage)
	if b == nil {
		return fmt.Errorf(color.RedString("剪切板没有图片，请先复制图片到剪切板"))
	}

	path, _ := homedir.Expand("~/Downloads")
	filename := filepath.Join(path, fmt.Sprintf("%d.png", time.Now().Unix()))
	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf(color.RedString("create file error: %w", err))
	}
	defer f.Close()
	_, _ = io.Copy(f, bytes.NewReader(b))

	fmt.Println(color.GreenString("clipboard image save to %s", filename))
	return nil
}
