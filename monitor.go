package monitor

import (
	"net/http"
	"strings"
	"log"
	"bufio"
	"net/http/httputil"
	"fmt"
	"bytes"
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
		client := &http.Client{}
		jsonStr := []byte(`{"text": "hogehoge"}`)
		req, err := http.NewRequest("POST",
			"https://hooks.slack.com/services/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",
			bytes.NewBuffer(jsonStr))
		if err != nil {
			log.Println("Failed parsing", err)
			return err
		}
		req.Header.Set("Content-Type", "application/json")

		dump, err := httputil.DumpRequest(req, true)
		if err != nil {
			log.Println("Something happened", err)
			return err
		}
		fmt.Printf("%q\n", dump)

		resp, err := client.Do(req)
		if err != nil {
			log.Println("request failed", err)
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			s := bufio.NewScanner(resp.Body)
			s.Scan()
			log.Println(s.Text())
			log.Println("StatusCode =", resp.StatusCode)
		}

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