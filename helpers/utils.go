package helpers

/*
author Kryuchenko Vyacheslav
*/

import (
	"github.com/vkryuchenko/svnwrapper"
	"log"
)

func GetData(task WorkTask) error {
	var err error
	svn := svnwrapper.Svn{}
	if task.LocalPathValid {
		if task.HardReset {
			log.Printf("Hard reset %s", task.LocalPath)
			err = svn.HardReset([]string{task.LocalPath}, false)
			if err != nil {
				return err
			}
		} else {
			log.Printf("Cleanup %s", task.LocalPath)
			err = svn.Cleanup([]string{task.LocalPath}, []string{})
			if err != nil {
				return err
			}
			log.Printf("Revert %s", task.LocalPath)
			err = svn.Revert([]string{task.LocalPath}, []string{"-R"})
			if err != nil {
				return err
			}
		}
		log.Printf("Update %s from %s@%s", task.LocalPath, task.SvnURL, task.Revision)
	} else {
		log.Printf("Fresh checkout from %s@%s", task.SvnURL, task.Revision)
	}
	return svn.CheckOut(task.SvnURL, task.LocalPath, task.Revision, []string{"--ignore-externals", "--force"})
}
