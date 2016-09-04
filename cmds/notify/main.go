package main

import (
	"flag"
	"log"
	"errors"
)

func main() {
	var fatalErr error
	defer func() {
		if fatalErr != nil {
			flag.PrintDefaults()
			log.Fatalln(fatalErr)
		}
	}()

	var (
		// Command line flag
		dbpath = flag.String("db", "./backupdata", "bpath to db dir")
	)
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fatalErr = errors.New("Error: Specify command")
		return
	}
}
