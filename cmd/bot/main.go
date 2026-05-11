package main

import (
	"fmt"
	"log"
	"runtime"
)

func main() {
	// Configure the standard logger to include timestamp and file:line.
	// We'll swap this for a structured logger (slog) later.
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("====================================")
	fmt.Println("  Go-Trade Bot — bootstrapping     ")
	fmt.Println("====================================")

	log.Printf("Go runtime: %s on %s/%s",
		runtime.Version(), runtime.GOOS, runtime.GOARCH)
	log.Println("Bot is alive. Nothing else to do yet — that's the point of Step 3.")
}
