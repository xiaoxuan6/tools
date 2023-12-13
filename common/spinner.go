package common

import (
	"github.com/briandowns/spinner"
	"time"
)

var s = spinner.New(spinner.CharSets[0], 10*time.Millisecond)

func Start(prefix string) {
	s.Prefix = prefix + "doing..."
	s.Start()
}

func Stop() {
	s.Stop()
}
