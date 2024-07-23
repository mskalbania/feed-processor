package main

import (
	_ "feed-processor/matchers"
	"feed-processor/search"
	"log"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	log.Println("Starting application...")
	search.Run("weather")
}
