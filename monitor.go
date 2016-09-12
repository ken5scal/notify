package monitor

import "net/http"

type Monitor struct {
	Paths   map[string]string
	Service string
}

func (m *Monitor) Now() (int, error) {
	var counter int
	for path, lastHash := range m.Paths {
		newHash, err := DirHash(path)
		if err != nil {
			return 0, err
		}
		if newHash != lastHash {
			err := m.act(path, m.Service)
			if err != nil {
				return counter, err
			}
			m.Paths[path] = newHash
			counter++
		}
	}
	return counter, nil
}

func (m *Monitor) act(path string, service string) error {
	// TODO Notify to slack
	//dirname := filepath.Base(path)
	//filename := fmt.Sprintf("%d.zip", time.Now().UnixNano())
	//return m.Archiver.Archive(path, filepath.Join(m.Destination, dirname, filename))
	alert(path, service)
	return nil
}

func alert(path string, service string) error {
	switch service {
	case "slack":
		resp, err := http.Get("https://slack.com/")
	case "chatwork":
	case "email":
	case "empty":
	default:
		// send nothing
	}
	return nil
}