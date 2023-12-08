package common

import (
	"github.com/briandowns/spinner"
	"time"
)

var s = spinner.New(spinner.CharSets[0], 10*time.Millisecond)

func Start() {
	s.Prefix = "doing..."
	s.Start()
}

func Stop() {
	s.Stop()
}
