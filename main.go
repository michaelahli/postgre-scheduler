package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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
			log.Fatalln(err)
		}
	}

	log.Println("[OUTPUT]", res)
	log.Println("[OUTPUT]", s3session.GetURI())
}

func cron(scheduler *gocron.Scheduler) error {
	duration, err := strconv.Atoi(os.Getenv("DURATION"))
	if err != nil {
		log.Fatalln(err)
	}

	isTimeSpecified, err := strconv.ParseBool(os.Getenv("IS_TIME_SPECIFIED"))
	if err != nil {
		log.Fatalln(err)
	}

	switch isTimeSpecified {
	case true:
		var (
			splitTime    []int
			timeSpecific = os.Getenv("TIME_SPECIFIC")
			splitTimeStr = strings.Split(timeSpecific, " ")
		)

		for _, ststr := range splitTimeStr {
			stint, err := strconv.Atoi(ststr)
			if err != nil {
				log.Fatalln(err)
			}
			splitTime = append(splitTime, stint)
		}
		t := time.Date(
			splitTime[0],
			time.Month(splitTime[1]),
			splitTime[2],
			splitTime[3],
			splitTime[4],
			splitTime[5],
			splitTime[6],
			time.UTC,
		)

		scheduler.Every(uint64(duration)).Day().From(&t).Do(backup)

	default:
		scheduler.Every(uint64(duration)).Days().Do(backup)
	}

	return nil
}

func main() {
	config := helper.New()
	config.SetUp()

	scheduler := gocron.NewScheduler()

	cron(scheduler)
	<-scheduler.Start()
}
