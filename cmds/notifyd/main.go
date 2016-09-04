package main

import (
	"log"
	"flag"
	"github.com/ken5scal/goblueprints/chapter8/backup"
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

	m := &backup.Monitor{
		Paths: make(map[string]string),
	}
}
