package main

import (
	"github.com/antelman107/net-wait-go/wait"
	"time"
)

func waitForNetwork() {
	w := wait.New(
		wait.WithProto("tcp"),
		wait.WithWait(500*time.Millisecond),
		wait.WithBreak(100*time.Millisecond),
		wait.WithDeadline(24*time.Hour),
	)

	w.Do([]string{"example.com:80"})
}
