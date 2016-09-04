package main

import (
	"log"
	"flag"
	"github.com/ken5scal/notify"
	"github.com/matryer/filedb"
	"encoding/json"
	"errors"
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

	/*
		Cash Data
	 */
	var path path
	col.ForEach(func(_ int, data []byte) bool {
		if err := json.Unmarshal(data, &path); err != nil {
			fatalErr = err
			return true
		}

		m.Paths[path.Path] = path.Hash
		return false
	})
	if fatalErr != nil {
		return
	} else if len(m.Paths) < 1 {
		fatalErr = errors.New("Nopath exists. add path")
		return
	}
}
