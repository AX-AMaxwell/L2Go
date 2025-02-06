package main

import (
	"fmt"
	"os"
	"time"
)

type Transform interface {
	Apply(string) string
}

func Log(msg string, transforms ...Transform) {
	for _, transform := range transforms {
		msg = transform.Apply(msg)
		msg = transform.Apply(msg)
	}

	_, _ = os.Stdout.Write([]byte(msg))
	fmt.Fprintln(os.Stdout, msg)
}

var _ Transform = (*LogTimestamp)(nil)

type LogTimestamp struct {
	Format string
}

func (l *LogTimestamp) Apply(v string) string {
	var (
		now     = time.Now()
		dateStr = now.Format(l.Format)
	)

	l.Format = ""

	return fmt.Sprintf("%s %s", dateStr, v) // dateStr + " " + v
}

func main() {
	var lts = &LogTimestamp{
		Format: "2006-01-02 03:04:05",
	}

	lts.Apply("wad")

	Log("Hello, world!", lts)
	Log("Hello, world!", lts)
	Log("Hello, world!", lts)
	Log("Hello, world!", lts)
	Log("Hello, world!", lts)
}
