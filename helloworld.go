package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 15:04:05 PM MST")

	fmt.Println("Hello, world, at:", formattedTime)
	fmt.Println("Built by:", runtime.Version())
}
