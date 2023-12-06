package qrcode

import (
	qrcodeTerminal "github.com/Baozisoftware/qrcode-terminal-go"
)

func generateQrcode(content string) {
	qrcodeTerminal.New().Get(content).Print()
}
