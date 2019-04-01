package main

import (
	"encoding/csv"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"log"
	"os"
)

func main() {
	app := app.New()
	logfile, err := os.Create("results.csv")
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
				task.Text,
				category.Selected,
				enjoy.Selected,
				impact.Selected,
				comments.Text,
			})
			w.Close()
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
	w.ShowAndRun()
}
