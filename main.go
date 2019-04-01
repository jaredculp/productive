package main

import (
	"fmt"
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
	f := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "What task are you working on?", Widget: widget.NewEntry()},
			{Text: "How would you categorize this task?", Widget: widget.NewRadio(tasks, nil)},
			{Text: "Are you enjoying this task?", Widget: widget.NewRadio(rating, nil)},
			{Text: "Do you find this task impactful?", Widget: widget.NewRadio(rating, nil)},
			{Text: "Comments", Widget: widget.NewMultiLineEntry()},
		},
		OnSubmit: onSubmit,
		OnCancel: onCancel,
	}

	w.SetContent(f)
	w.ShowAndRun()
}

func onSubmit() {
	fmt.Println("Submitted")
}

func onCancel() {
	fmt.Println("Cancelled")
}
