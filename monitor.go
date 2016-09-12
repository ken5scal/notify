package monitor

import (
	"net/http"
	"strings"
	"log"
	"bufio"
	"net/http/httputil"
	"fmt"
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
		//log.Println("Slack")
		//resp, err := http.Post("https://hooks.slack.com/services/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",
		//	"application/json", strings.NewReader("text: hogehoge"))
		//if err != nil {
		//	log.Println(err.Error())
		//}
		//defer resp.Body.Close()
		//cli
		//
		//u, err := url.Parse("https://hooks.slack.com/services/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")
		//if err != nil {
		//	log.Println("Failed parsing",err)
		//	return err
		//}
		//
		////query := url.Values{"text": }
		//req, err := http.NewRequest("POST", u.String(), strings.NewReader("text: hogehoge"))
		//if err != nil {
		//	log.Println("Failed parsing",err)
		//	return err
		//}
		//
		////req.Body.Read(strings.NewReader("text: hogehoge"))
		//req.Header.Set("Content-Type", "application/json")
		////http.Client.Do(req)

		//{"text": "This is a line of text in a channel.\nAnd this is another line of text."}
		client := &http.Client{}
		req, err := http.NewRequest("POST",
			"https://hooks.slack.com/services/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",
			strings.NewReader("text: hogehoge"))
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
		fmt.Printf("%q", dump)

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