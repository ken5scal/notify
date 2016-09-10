package monitor

import (
	//"errors"
	"fmt"
	"github.com/matryer/filedb"
)

type path struct {
	Path string
	Hash string
}

func AddPath(dbPath string, monitorPaths string) error {
	db, err := filedb.Dial(dbPath)
	if err != nil {
		return err
	}
	defer db.Close()

	// Obtain column
	col, err := db.C("paths")
	if err != nil {
		return err
	}

	fmt.Printf("Argument: %s\n", monitorPaths)

	path := &path{Path:monitorPaths, Hash:"Not yet archived"}
	if err := col.InsertJSON(path); err != nil {
		return err
	}
	fmt.Printf("+ %s\n", path)

	return nil
}