package monitor

import (
	"errors"
	"fmt"
	"github.com/matryer/filedb"
)

type path struct {
	Path string
	Hash string
}

func AddPath(dbPath string, monitorPaths []string) error {
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

	var prevPath path
	if len(monitorPaths) == 0 {
		return errors.New("Specify path to add")
	}
	for _, p := range monitorPaths {
		fmt.Printf("Argument: %s\n", p)
		if prevPath.Path == p {
			continue
		}

		path := &path{Path:p, Hash:"Not yet archived"}
		if err := col.InsertJSON(path); err != nil {
			return err
		}
		fmt.Printf("+ %s\n", path)
		prevPath = *path
	}

	return nil
}