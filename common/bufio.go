package common

import (
	"bufio"
	"os"
	"strings"
	"time"
)

func ScannerF(fn func() string) string {
	var ch = make(chan string)
	go func() {
		scanner := bufio.NewScanner(os.Stdin)

		var body []string
		for scanner.Scan() {
			body = append(body, scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			ch <- ""
			return
		}
		ch <- strings.Join(body, "\n")
	}()

	select {
	case body := <-ch:
		return strings.TrimSpace(body)
	case <-time.After(2 * time.Second):
		return fn()
	}
}
