package main

import (
	"encoding/csv"
	"log"
	"os"
	"time"

	"github.com/gen2brain/dlgs"
)

const title = "Productive"
const output = "results.csv"

var header = []string{"time", "task", "category", "enjoy", "impact", "comments"}

var categories = []string{
	"Clerical",
	"Internal meeting / conference call",
	"External meeting / conference call",
	"Constituent meeting / call",
	"Outreach",
	"Project",
	"Graphics",
	"Non-work",
	"Other",
}

var rating = []string{"1", "2", "3", "4", "5"}

func main() {
	logfile, err := os.OpenFile(output, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer logfile.Close()

	writer := csv.NewWriter(logfile)
	defer writer.Flush()

	lines, err := csv.NewReader(logfile).ReadAll()
	if err != nil {
		log.Fatal("Cannot read file", err)
	}
	if len(lines) == 0 {
		writer.Write(header)
	}

	done := make(chan bool, 1)
	for {
		select {
		case <-done:
			return
		default:
			prompt(done, writer)
		}
	}
}

func prompt(done chan bool, writer *csv.Writer) {
	log.Println("Prompting")

	task, _, _ := dlgs.Entry(title, "What task are you working on?", "")
	category, _, _ := dlgs.List(title, "How would you categorize this task?", categories)
	enjoy, _, _ := dlgs.List(title, "Are you enjoying this task?", rating)
	impact, _, _ := dlgs.List(title, "Do you find this task impactful?", rating)
	comment, _, _ := dlgs.Entry(title, "Comments", "")
	again, _ := dlgs.Question(title, "Do you want to continue?", false)

	timestamp := time.Now().Format(time.UnixDate)

	writer.Write([]string{timestamp, task, category, enjoy, impact, comment})

	if !again {
		done <- true
	} else {
		time.Sleep(2 * time.Hour)
	}
}
