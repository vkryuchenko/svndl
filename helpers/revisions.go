package helpers

/*
author Kryuchenko Vyacheslav
*/

import (
	"io/ioutil"
	"log"
	"strings"
)

type Revisions struct {
	Map map[string]string
}

func (r *Revisions) Read(revisionsPath string) error {
	log.Printf("Read %s\n", revisionsPath)
	content, err := ioutil.ReadFile(revisionsPath)
	if err != nil {
		return err
	}

	for _, line := range strings.Split(string(content[:]), "\n") {
		data := strings.Split(line, "=")
		if len(data) > 1 {
			r.Map[data[0]] = strings.TrimSuffix(data[1], "\r")
		}
	}
	return nil
}
