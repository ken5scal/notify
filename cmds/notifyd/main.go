package main

import (
	"log"
	"flag"
)

func main() {
	var fatalErr error
	defer func() {
		if fatalErr != nil {
			log.Fatalln(fatalErr)
		}
	}()

	var (
		interval = flag.Int("interval", 10, "Check duration per sec")
		// service = flag.String("service", "slack", "Notify service")
	)

	flag.Parse()
}
