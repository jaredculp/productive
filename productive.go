package main

import (
	"encoding/csv"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"log"
	"os"
	"time"
)

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
	logfile, err := os.OpenFile("results.csv", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0600)
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
		writer.Write([]string{"time", "task", "category", "enjoy", "impact", "comments"})
	}

	w := app.New().NewWindow("Productive")
	w.SetContent(newForm(w, writer))
	w.CenterOnScreen()
	w.ShowAndRun()
}

func newForm(window fyne.Window, writer *csv.Writer) fyne.CanvasObject {
	task := widget.NewEntry()
	category := widget.NewRadio(categories, nil)
	enjoy := widget.NewRadio(rating, nil)
	impact := widget.NewRadio(rating, nil)
	comments := widget.NewMultiLineEntry()

	f := &widget.Form{
		OnSubmit: func() {
			log.Println("Form submitted")
			writer.Write([]string{
				time.Now().Format(time.UnixDate),
				task.Text,
				category.Selected,
				enjoy.Selected,
				impact.Selected,
				comments.Text,
			})
			writer.Flush()

			window.Hide()
			time.Sleep(2 * time.Second)
			log.Println("Refreshing form")
			window.SetContent(newForm(window, writer))
			window.RequestFocus()
		},
		OnCancel: func() {
			log.Println("Form cancelled")
			window.Close()
		},
	}

	f.Append("What task are you working on?", task)
	f.Append("How would you categorize this task?", category)
	f.Append("Are you enjoying this task?", enjoy)
	f.Append("Do you find this task impactful?", impact)
	f.Append("Comments", comments)

	return f
}
