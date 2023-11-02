package common

import "github.com/fatih/color"

var ExtError = color.RedString("文件格式错误，请重新输入，支持的格式有：%s", "jpg, jpeg, png")

func Identify(ext string) bool {
	stat := false
	extAllow := []string{".jpg", ".jpeg", ".png"}
	for _, v := range extAllow {
		if ext == v {
			stat = true
		}
	}

	return stat
}
