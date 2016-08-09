package helpers

/*
author Kryuchenko Vyacheslav
*/

import (
	"encoding/json"
	"os"
)

type WorkTask struct {
	SvnURL    string `json:"SvnUrl"`
	LocalPath string `json:"LocalPath"`
	HardReset bool   `json:"HardReset"`
}

type WorkProfile struct {
	Tasks []WorkTask `json:"Tasks"`
}

func (wp *WorkProfile) Read(profilePath string) error {
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
