package qrcode

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/tuotoo/qrcode"
	"os"
)

func scan(file *os.File) {
	qrMatrix, err := qrcode.Decode(file)
	if err != nil {
		fmt.Println(color.RedString("解析失败，请重新输入"))
		return
	}

	fmt.Println(color.GreenString("解析结果："), qrMatrix.Content)
}
