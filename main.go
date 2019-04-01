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

	task := widget.NewEntry()
	category := widget.NewRadio(categories, nil)
	enjoy := widget.NewRadio(rating, nil)
	impact := widget.NewRadio(rating, nil)
	comments := widget.NewMultiLineEntry()

	w := app.NewWindow("Productive")
	f := &widget.Form{
    OnSubmit: func() {
      fmt.Println("Submitted")
      fmt.Println(task.Text)
      fmt.Println(category.Selected)
      fmt.Println(enjoy.Selected)
      fmt.Println(impact.Selected)
      fmt.Println(comments.Text)
    },
    OnCancel: func() {
      fmt.Println("Cancelled")
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
