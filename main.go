package main

import (
	"log"
	"os"
	"strconv"

	"github.com/michaelahli/postgre-scheduler/cmd"
	"github.com/michaelahli/postgre-scheduler/helper"

	"github.com/jasonlvhit/gocron"
)

func backup() {
	log.Println("[SYSTEM] Backing up DB...")
	executor := cmd.NewTerminal("bash")

	res, err := executor.ExecuteBash(
		os.Getenv("EXE"),
		"-d", os.Getenv("DB"),
		"-f", os.Getenv("DIR"),
		"-k", os.Getenv("KEEP"),
	)
	if err != nil {
		log.Fatalln("[ERROR]", err)
	}

	log.Println("[OUTPUT]", res)
	return
}

func main() {
	config := helper.New()
	config.SetUp()

	scheduler := gocron.NewScheduler()

	duration, err := strconv.Atoi(os.Getenv("DURATION"))
	if err != nil {
		log.Panicln(err)
	}

	scheduler.Every(uint64(duration)).Seconds().Do(backup)
	<-scheduler.Start()
}
