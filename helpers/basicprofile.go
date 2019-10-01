package helpers

/*
author Kryuchenko Vyacheslav
*/

import (
	"encoding/json"
	"log"
	"os"
)

type BasicProfile struct {
	Tasks []WorkTask `json:"Tasks"`
}

func (wp *BasicProfile) Read(profilePath string) error {
	log.Printf("Read %s\n", profilePath)
	cf, err := os.Open(profilePath)
	if err != nil {
		return err
	}
	defer cf.Close()
	decoder := json.NewDecoder(cf)
	decodeErr := decoder.Decode(wp)
	if decodeErr != nil {
		return err
	}
	return nil
}
