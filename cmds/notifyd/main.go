package main

import (
	"log"
	"flag"
	"github.com/ken5scal/notify"
	"github.com/matryer/filedb"
)

type path struct {
	Path string
	Hash string
}

func main() {
	var fatalErr error
	defer func() {
		if fatalErr != nil {
			log.Fatalln(fatalErr)
		}
	}()

	var (
		interval = flag.Int("interval", 10, "Check duration per sec")
		service = flag.String("service", "slack", "Notify service")
		dbpath = flag.String("db", "./db", "path to file db")
	)

	flag.Parse()

	m := &monitor.Monitor{
		Paths: make(map[string]string),
		Service: *service,
	}

	db, err := filedb.Dial(*dbpath)
	if err != nil {
		fatalErr = err
		return
	}
	defer db.Close()

	col, err := db.C("paths")
	if err != nil {
		fatalErr = err
		return
	}
}
