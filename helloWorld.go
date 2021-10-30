package main
/*
For this app to run you need to install fyne.io/fyne
This can be done with the below command
go get fyne.io/fyne

Then compile and run
*/
import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"fmt"
}

func main() {
	fmt.Println("Hello, World!")
	a := app.New()

	w := a.NewWindows("Hello World")
	w.Setcontent(widget.NewVBox(
		widget.NewLabel("Hello World"),
		widget.NewButton("Quit", func(){a.Quit()
		}),
	))
w.ShowAndRun()
}
