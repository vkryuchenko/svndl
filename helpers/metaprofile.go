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
	if err != nil {
		return err
	}
	defer cf.Close()
	decoder := json.NewDecoder(cf)
	decodeErr := decoder.Decode(metaProfile)
	if decodeErr != nil {
		return err
	}
	if len(metaProfile.BasicProfiles) < 1 {
		return fmt.Errorf("%s not any basic profiles included", profilePath)
	}

	basePath := filepath.Dir(profilePath)
	for _, basicProfilePath := range metaProfile.BasicProfiles {
		basicProfile := BasicProfile{}
		err := basicProfile.Read(filepath.Join(basePath, basicProfilePath))
		if err != nil {
			fmt.Println(err)
			continue
		}
		for _, item := range basicProfile.Tasks {
			metaProfile.Tasks = append(metaProfile.Tasks, item)
		}
	}
	if len(metaProfile.Tasks) < 1 {
		return fmt.Errorf("not any tasks for profile %s", profilePath)
	}
	return nil
}
