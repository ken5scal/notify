package monitor

import (
	"net/http"
	"strings"
	"log"
)

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
	log.Println("call alert")
	alert(path, "slack")
	return nil
}

func alert(path string, service string) error {
	log.Println("Alert initiated")
	switch service {
	case "slack":
		//u, err := url.Parse("https://hooks.slack.com/services/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")
		//if err != nil {
		//	log.Println("Failed in parsing URL", err)
		//	return
		//}
		//
		//query := url.Values{""}
		//query := url.{"text": "hogehoghoegegeg."}
		log.Println("Slack")
		resp, err := http.Post("https://hooks.slack.com/services/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",
			"application/json", strings.NewReader("text: hogehoge"))
		if err != nil {
			log.Println(err.Error())
		}
		defer resp.Body.Close()

	case "chatwork":
	case "email":
	case "empty":
	default:
		log.Println("default")
		// send nothing
		//query := url.Values{"text": "hogehoghoegegeg."}
		//http.Post("https://hooks.slack.com/services/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",
		//	"application/json", strings.NewReader(query.Encode()))
		http.Post("https://hooks.slack.com/services/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",
			"application/json", strings.NewReader("text: hogehoge"))
	}
	return nil
}