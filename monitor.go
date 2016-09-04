package monitor

import (
	"fmt"
	"errors"
)

type Monitor struct {
	Paths map[string]string
	Service string
}

func (m *Monitor) Now() (int, error) {
	var counter int
	for path, lastHash := range m.Paths {
		newHash, err := DirHash(path)
		if err != nil {
			fmt.Println("hogehoge")
			return 0, err
		}
		if newHash != lastHash {
			err := m.act(path)
			if err != nil {
				fmt.Println("fugafua")
				return counter, err
			}
			m.Paths[path] = newHash
			counter++
		}
	}
	return counter, nil
}

func (m *Monitor) act(path string) error {
	// TODO Something
	//dirname := filepath.Base(path)
	//filename := fmt.Sprintf("%d.zip", time.Now().UnixNano())
	//return m.Archiver.Archive(path, filepath.Join(m.Destination, dirname, filename))
	fmt.Println("Act has bee invoked...")
	return errors.New("hogehoge")
}