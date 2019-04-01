package main

import (
	"fmt"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	app := app.New()

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

	task := &widget.FormItem{
		Text:   "What task are you working on?",
		Widget: widget.NewEntry(),
	}

	category := &widget.FormItem{
		Text:   "How would you categorize this task?",
		Widget: widget.NewRadio(categories, nil),
	}

	enjoy := &widget.FormItem{
		Text:   "Are you enjoying this task?",
		Widget: widget.NewRadio(rating, nil),
	}

	impact := &widget.FormItem{
		Text:   "Do you find this task impactful?",
		Widget: widget.NewRadio(rating, nil),
	}

	comments := &widget.FormItem{
		Text:   "Comments",
		Widget: widget.NewMultiLineEntry(),
	}

	w := app.NewWindow("Productive")
	f := &widget.Form{
		Items:    []*widget.FormItem{task, category, enjoy, impact, comments},
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
