package helpers

/*
author Kryuchenko Vyacheslav
*/

import (
	"log"
	"os"
	"svnwrapper"
)

func GetData(uri string, target string, revision string, useHardReset bool) error {
	var err error
	svn := svnwrapper.Svn{}
	if _, err = os.Stat(target); err == nil {
		if useHardReset {
			log.Printf("Hard reset %s", target)
			err = svn.HardReset([]string{target}, false)
			if err != nil {
				return err
			}
		} else {
			log.Printf("Cleanup %s", target)
			err = svn.Cleanup([]string{target}, []string{})
			if err != nil {
				return err
			}
			log.Printf("Revert %s", target)
			err = svn.Revert([]string{target}, []string{"-R"})
			if err != nil {
				return err
			}
		}
	}
	log.Printf("Checkout %s@%s", uri, revision)
	return svn.CheckOut(uri, target, revision, []string{"--ignore-externals", "--force"})
}
