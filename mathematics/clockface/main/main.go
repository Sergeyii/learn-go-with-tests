package main

import (
	"../../clockface"
	"os"
	"time"
)

func main() {
	clockface.SVGWriter(os.Stdout, time.Now())
}
