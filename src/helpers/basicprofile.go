package helpers

/*
author Kryuchenko Vyacheslav
*/

import (
	"encoding/json"
	"log"
	"os"
	"strings"
)

type WorkTask struct {
	SvnURL    string `json:"SvnUrl"`
	LocalPath string `json:"LocalPath"`
	HardReset bool   `json:"HardReset"`
	Revision  string
}

type BasicProfile struct {
	Tasks []WorkTask `json:"Tasks"`
}

func (wt *WorkTask) CheckRevision(variants map[string]string) {
	for url, revision := range variants {
		if strings.Contains(wt.SvnURL, url) {
			wt.Revision = revision
			return
		}
	}
	wt.Revision = "HEAD"
}

func (wp *BasicProfile) Read(profilePath string) error {
	log.Printf("Read %s\n", profilePath)
	cf, err := os.Open(profilePath)
	defer cf.Close()
	if err != nil {
		return err
	}
	decoder := json.NewDecoder(cf)
	decode_err := decoder.Decode(wp)
	if decode_err != nil {
		return err
	}
	return nil
}
