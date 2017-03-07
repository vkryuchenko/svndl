package main

/*
author Kryuchenko Vyacheslav
*/

import (
	"flag"
	"helpers"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"sync"
)

const (
	minProcCount = 2
)

func main() {
	var workersCount string
	var profilePath string
	var revisionsPath string
	var processCount int
	var revisions helpers.Revisions
	//var err error
	cpuCount := runtime.NumCPU()
	runPath, err := os.Getwd()
	if err != nil {
		log.Panic(err)
	}

	flag.StringVar(&workersCount, "workers", "auto", "Workers count. Must be more then 0 or auto(is default).")
	flag.StringVar(&profilePath, "profile", "", "Path to work profile.")
	flag.StringVar(&revisionsPath, "revisions", "", "Path to file with checkout revisions")
	flag.Parse()

	log.SetFlags(0) // disable print date end time

	if profilePath == "" {
		log.Panic("Profile not set!")
	}

	profile := helpers.MetaProfile{}
	if err := profile.Read(profilePath); err != nil {
		log.Panic(err)
	}

	if workersCount == "auto" {
		if (cpuCount % 2) == 0 {
			processCount = cpuCount / 2
		} else {
			processCount = (cpuCount - 1) / 2
		}
	} else {
		processCount, err = strconv.Atoi(workersCount)
		if err != nil {
			log.Panic(err)
		}
	}
	if processCount <= minProcCount {
		processCount = minProcCount
	}
	if processCount >= len(profile.Tasks) {
		processCount = len(profile.Tasks)
	}

	revisions = helpers.Revisions{Map: make(map[string]string)}
	if revisionsPath != "" {
		revisions.Read(revisionsPath)
	} else {
		revisions.Map["all"] = "HEAD"
	}
	// remove invalid folders
	for index, task := range profile.Tasks {
		valid := task.LocalPathValid
		task.LocalPath = filepath.Join(runPath, task.LocalPath)
		valid = helpers.CheckLocalPathValid(task)
		if !valid {
			log.Printf("Invalid path -- %s", task.LocalPath)
			err = os.RemoveAll(task.LocalPath)
			if err != nil {
				panic(err)
			}
		} else {
			profile.Tasks[index].LocalPathValid = valid
		}
	}

	// get data
	getDataChanel := make(chan helpers.WorkTask, len(profile.Tasks))
	for _, task := range profile.Tasks {
		task.LocalPath = filepath.Join(runPath, task.LocalPath)
		getDataChanel <- task
	}
	getDataGroup := sync.WaitGroup{}
	getDataGroup.Add(processCount)
	for i := 0; i < processCount; i++ {
		go func() {
			defer getDataGroup.Done()
			for len(getDataChanel) > 0 {
				task := <-getDataChanel
				task.CheckRevision(revisions.Map)
				if err := helpers.GetData(task); err != nil {
					panic(err)
				}
			}
		}()
	}
	getDataGroup.Wait()
}
