package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	app := app.New()

  tasks := []string{
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


	w := app.NewWindow("Productive")
  f := widget.NewForm()
  f.Append("What task are you working on?", widget.NewEntry())
	f.Append("How would you categorize this task?", widget.NewRadio(tasks, nil))
	f.Append("Are you enjoying this task?", widget.NewRadio(rating, nil))
	f.Append("Do you find this task impactful?", widget.NewRadio(rating, nil))
	f.Append("Comments", widget.NewMultiLineEntry())
  f.Append("", widget.NewButton("Submit", nil))

	w.SetContent(f)
	w.ShowAndRun()
}
