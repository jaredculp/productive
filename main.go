package main

import (
	"encoding/csv"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"log"
	"os"
	"time"
)

func main() {
	app := app.New()
	logfile, err := os.OpenFile("results.csv", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer logfile.Close()

	writer := csv.NewWriter(logfile)
	defer writer.Flush()

	categories := []string{
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

	rating := []string{"1", "2", "3", "4", "5"}

	task := widget.NewEntry()
	category := widget.NewRadio(categories, nil)
	enjoy := widget.NewRadio(rating, nil)
	impact := widget.NewRadio(rating, nil)
	comments := widget.NewMultiLineEntry()

	w := app.NewWindow("Productive")
	f := &widget.Form{
		OnSubmit: func() {
			writer.Write([]string{
				time.Now().Format(time.UnixDate),
				task.Text,
				category.Selected,
				enjoy.Selected,
				impact.Selected,
				comments.Text,
			})

			w.Hide()
			time.Sleep(2 * time.Hour)
			w.RequestFocus()
		},
		OnCancel: func() {
			w.Close()
		},
	}
	f.Append("What task are you working on?", task)
	f.Append("How would you categorize this task?", category)
	f.Append("Are you enjoying this task?", enjoy)
	f.Append("Do you find this task impactful?", impact)
	f.Append("Comments", comments)

	w.SetContent(f)
	w.CenterOnScreen()
	w.ShowAndRun()
}
