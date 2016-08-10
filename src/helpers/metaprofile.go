package helpers

/*
author Kryuchenko Vyacheslav
*/

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type MetaProfile struct {
	Tasks         []WorkTask
	BasicProfiles []string `json:"Include"`
}

func (metaProfile *MetaProfile) Read(profilePath string) error {
	cf, err := os.Open(profilePath)
	defer cf.Close()
	if err != nil {
		return err
	}
	decoder := json.NewDecoder(cf)
	decode_err := decoder.Decode(metaProfile)
	if decode_err != nil {
		return err
	}
	if len(metaProfile.BasicProfiles) < 1 {
		return fmt.Errorf("%s not any basic profiles included", profilePath)
	}

	basePath := filepath.Dir(profilePath)
	for _, basicProfilePath := range metaProfile.BasicProfiles {
		basicProfile := BasicProfile{}
		basicProfile.Read(filepath.Join(basePath, basicProfilePath))
		if len(basicProfile.Tasks) > 0 {
			for _, item := range basicProfile.Tasks {
				metaProfile.Tasks = append(metaProfile.Tasks, item)
			}
		}
	}
	if len(metaProfile.Tasks) < 1 {
		return fmt.Errorf("Not any tasks for profile %s", profilePath)
	}
	return nil
}
