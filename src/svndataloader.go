package main

/*
author Kryuchenko Vyacheslav
*/

import (
	"flag"
	"log"
	"path/filepath"
	"runtime"
	"strconv"
	"sync"
	"os"
	"helpers"
)

const (
	MIN_PROC_COUNT = 2
)

func main() {
	var workersCount string
	var profilePath string
	var processCount int
	//var err error
	cpuCount := runtime.NumCPU()
	runPath, err := os.Getwd()
	if err != nil {
		log.Panic(err)
	}

	flag.StringVar(&workersCount, "workers", "auto", "Workers count. Must be more then 0 or auto(is default).")
	flag.StringVar(&profilePath, "profile", "", "Path to work profile.")
	flag.Parse()

	if profilePath == "" {
		log.Panic("Profile not set!")
	}

	profile := helpers.WorkProfile{}
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
	if processCount < MIN_PROC_COUNT {
		processCount = MIN_PROC_COUNT
	}
	if processCount > len(profile.Tasks) {
		processCount = len(profile.Tasks)
	}

	taskChanel := make(chan helpers.WorkTask, len(profile.Tasks))
	for _, task := range profile.Tasks {
		taskChanel <- task
	}

	wg := sync.WaitGroup{}
	wg.Add(processCount)

	for i := 0; i < processCount; i++ {
		go func() {
			defer wg.Done()
			for len(taskChanel) > 0 {
				task := <- taskChanel
				targetPath := filepath.Join(runPath, task.LocalPath)
				if err := helpers.GetData(task.SvnURL, targetPath, task.HardReset); err != nil {
					log.Panic(err)
				}
			}
		}()
	}
	wg.Wait()
}
