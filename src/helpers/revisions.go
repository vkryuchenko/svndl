package helpers

/*
author Kryuchenko Vyacheslav
*/

import (
	"io/ioutil"
	"log"
	"runtime"
	"strings"
)

type Revisions struct {
	Map map[string]string
}

func (r *Revisions) Read(revisionsPath string) error {
	var lineSeparator string
	log.Printf("Read %s\n", revisionsPath)
	content, err := ioutil.ReadFile(revisionsPath)
	if err != nil {
		return err
	}
	if runtime.GOOS == "windows" {
		lineSeparator = "\r\n"
	} else {
		lineSeparator = "\n"
	}

	for _, line := range strings.Split(string(content[:]), lineSeparator) {
		data := strings.Split(line, "=")
		if len(data) > 1 {
			r.Map[data[0]] = data[1]
		}
	}
	return nil
}
