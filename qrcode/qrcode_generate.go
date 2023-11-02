package qrcode

import (
	"github.com/mdp/qrterminal"
	"os"
)

func generateQrcode(content string) {
	cfg := qrterminal.Config{
		Level:     qrterminal.M,
		Writer:    os.Stdout,
		QuietZone: 0,
		BlackChar: qrterminal.BLACK,
		WhiteChar: qrterminal.WHITE,
	}

	qrterminal.GenerateWithConfig(content, cfg)
}
