/*
author Vyacheslav Kryuchenko
*/
package helpers

import (
	"log"
	"os"
	"strings"
	"svnwrapper"
)

type WorkTask struct {
	SvnURL         string `json:"SvnUrl"`
	LocalPath      string `json:"LocalPath"`
	HardReset      bool   `json:"HardReset"`
	Revision       string
	LocalPathValid bool
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

func CheckLocalPathValid(wt WorkTask) bool {
	svn := svnwrapper.Svn{}
	_, err := os.Stat(wt.LocalPath)
	if err != nil {
		log.Print(err)
		return false
	}
	_, err = svn.Info(wt.LocalPath, []string{})
	if err != nil {
		log.Print(err)
		return false
	}
	return true
}
