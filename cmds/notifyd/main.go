package main

import (
	"log"
	"flag"
	"github.com/ken5scal/notify"
	"github.com/matryer/filedb"
	"encoding/json"
	"errors"
	"os"
	"os/signal"
	"syscall"
	"time"
	"fmt"
)

type path struct {
	Path string
	Hash string
}

func dialDb(db *filedb.DB, dbpath string) (*filedb.DB, error) {
	db, err := filedb.Dial(dbpath)
	if err != nil {
		if err == filedb.ErrDBNotFound {
			// TODO should do following
			// ./notify -db=dbpath add ${Something like ls -l monitorPath | grep -v db}
			// Then call Dial Again

			return dialDb(db, ".")
		} else {
			return nil, err
		}
	}
	return db, nil
}

func main() {
	var fatalErr error
	defer func() {
		if fatalErr != nil {
			log.Fatalln(fatalErr)
		}
	}()

	var (
		// These flag.* methods do not return actual types, but pointers
		interval = flag.Int("interval", 10, "Check duration per sec")
		service = flag.String("service", "slack", "Notify service")
		dbpath = flag.String("db", "./db", "path to file db")
		//monitorPath = flag.String("monitor", ".", "monitoring path")
	)

	flag.Parse()

	m := &monitor.Monitor{
		Paths: make(map[string]string),
		Service: *service,
	}

	db, _ := dialDb(nil, *dbpath)
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

	/*
		Infinite Loop
	 */
	check(m, col)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	Loop: for {
		// Loop: is a label of this for loop
		select {
		case <-time.After(time.Duration(*interval) * time.Second):
			check(m, col)
		case <-signalChan:
		// Finishing Loop
			fmt.Println()
			log.Printf("Ending...")
			break Loop
		}
	}
}

func check(m *monitor.Monitor, col *filedb.C) {
	log.Println("Checking...")
	counter, err := m.Now()
	if err != nil {
		log.Fatalln("failed to backup:", err)
	}
	if counter > 0 {
		log.Printf("  Archived %d directories\n", counter)
		// update hashes
		var path path
		col.SelectEach(func(_ int, data []byte) (bool, []byte, bool) {
			if err := json.Unmarshal(data, &path); err != nil {
				log.Println("failed to unmarshal data (skipping):", err)
				return true, data, false
			}
			path.Hash, _ = m.Paths[path.Path]
			newdata, err := json.Marshal(&path)
			if err != nil {
				log.Println("failed to marshal data (skipping):", err)
				return true, data, false
			}
			return true, newdata, false
		})
	} else {
		log.Println("  No changes")
	}
}
