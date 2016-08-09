package helpers

/*
author Kryuchenko Vyacheslav
*/

import (
	"bitbucket.org/slavyan85/svnwrapper"
	"os"
)

func GetData(uri string, target string, useHardReset bool) error {
	var err error
	svn := svnwrapper.Svn{}
	if _, err = os.Stat(target); err == nil {
		if useHardReset {
			err = svn.HardReset([]string{target}, false)
			if err != nil {
				return err
			}
		} else {
			err = svn.Cleanup([]string{target}, []string{})
			if err != nil {
				return err
			}
			err = svn.Revert([]string{target}, []string{"-R"})
			if err != nil {
				return err
			}
		}
	}
	return svn.CheckOut(uri, target, 0, []string{"--ignore-externals", "--force"})
}
