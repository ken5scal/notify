package main

import (
	"flag"
	"log"
	"errors"
	"github.com/matryer/filedb"
	"strings"
	"encoding/json"
	"fmt"
)

type path struct {
	Path string
	Hash string
}

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
	args := flag.Args() // Return arguments as slice (excluding command line flag"
	if len(args) < 1 {
		fatalErr = errors.New("Error: Specify command")
		return
	}

	/*
		Persist Data
	 */
	db, err := filedb.Dial(*dbpath)
	if err != nil {
		fatalErr = err
		return
	}
	defer db.Close()

	// Obtain column
	col, err := db.C("paths")
	if err != nil {
		fatalErr = err
		return
	}

	switch strings.ToLower(args[0]) {
	case "list":
		var path path
		col.ForEach(func(i int, data []byte) bool {
			err := json.Unmarshal(data, &path)
			if err != nil {
				fatalErr = err
				return true
			}
			fmt.Printf("= %s\n", path)
			return false
		})
	case "add":
	case "remove":
	}
}
