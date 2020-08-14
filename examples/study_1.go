package examples

import (
	rapp "fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func Raw1() {
	// 初始化一个app
	app := rapp.New()

	// 新建个窗口
	w := app.NewWindow("Hello")

	w.SetContent(widget.NewVBox(

		// label 标签

		widget.NewLabel("Hello Fyne!"),
		widget.NewLabel("hello2"),

		// 退出按钮
		widget.NewButton("Quit", func() {
			app.Quit()

		}),
		// text
	))
	// 展示窗口一
	w.ShowAndRun()
}
