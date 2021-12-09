package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/michaelahli/postgre-scheduler/cmd"
	"github.com/michaelahli/postgre-scheduler/helper"
	"github.com/michaelahli/postgre-scheduler/utils"

	"github.com/jasonlvhit/gocron"
)

func backup() {
	log.Println("[SYSTEM] Backing up DB...")

	var (
		executor         = cmd.NewTerminal("bash")
		s3session        = utils.NewS3Session()
		year, month, day = time.Now().Date()
	)

	res, err := executor.ExecuteBash(
		os.Getenv("EXE"),
		"-h", os.Getenv("HOST"),
		"-d", os.Getenv("DB"),
		"-f", os.Getenv("DIR"),
		"-k", os.Getenv("KEEP"),
		"-u", os.Getenv("USERNAME"),
		"-p", os.Getenv("PASSWORD"),
	)
	if err != nil {
		log.Fatalln("[ERROR]", err)
	}

	executor.ExecuteBash()
	if os.Getenv("AWS_S3_BACKUP") == "true" {
		_, err = s3session.UploadObjectbyFilePath(fmt.Sprintf("./static/database-%02d-%02d-%d.sql", day, int(month), year))
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Println("[OUTPUT]", res)
	log.Println("[OUTPUT]", s3session.GetURI())

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

	scheduler.Every(uint64(duration)).Days().Do(backup)
	<-scheduler.Start()
}
