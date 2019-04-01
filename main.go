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
	w.SetContent(widget.NewVBox(
		widget.NewLabel("What task are you working on?"),
    widget.NewEntry(),

		widget.NewLabel("How would you categorize this task?"),
    widget.NewRadio(tasks, nil),

    widget.NewLabel("Are you enjoying this task?"),
    widget.NewRadio(rating, nil),

    widget.NewLabel("Do you find this task impactful?"),
    widget.NewRadio(rating, nil),

    widget.NewLabel("Comments"),
    widget.NewMultiLineEntry(),

		widget.NewButton("Submit", func() {
			app.Quit()
		}),
	))

	w.ShowAndRun()
}
