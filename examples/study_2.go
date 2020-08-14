package examples

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"net/url"
)

type Diagonal struct {
}

func (d *Diagonal) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	pos := fyne.NewPos(0, containerSize.Height-d.MinSize(objects).Height)
	// pos := fyne.NewPos(0, containerSize.Height-d.MinSize(objects).Height)
	for _, o := range objects {
		size := o.MinSize()
		o.Resize(size)
		o.Move(pos)

		pos = pos.Add(fyne.NewPos(0, size.Height)) // 横向不移动
		// pos = pos.Add(fyne.NewPos(size.Width, size.Height))
	}
}
func (d *Diagonal) MinSize(objects []fyne.CanvasObject) fyne.Size {
	w, h := 500, 800
	// for _, o := range objects {
	// 	childSize := o.MinSize()

	// 	w += childSize.Width
	// 	h += childSize.Height
	// }
	return fyne.NewSize(w, h)
}

func Raw2() {
	// 初始化一个app
	a := app.New()
	w := a.NewWindow("Disagonal")
	w.SetTitle("xxxxxx")

	text1 := widget.NewLabel("topleft")
	text2 := widget.NewLabel("Middle Label")
	text3 := widget.NewLabel("bottomright")
	// checkbox
	check1 := widget.NewCheck("hellp", func(ok bool) {
		fmt.Println("ok---", ok)
	})
	// 盒子一
	group1 := widget.NewGroup("group1", text1, text2, text3)
	group2 := widget.NewGroup("group2")

	// link
	link, _ := url.Parse("https://www.baidu.com")
	linkToBaidu := widget.NewHyperlink("go to baidu", link)
	// radio
	strs := []string{"eat", "drink", "play"}
	radio1 := widget.NewRadio(strs, func(s string) {
		fmt.Println(s)
	})
	box1 := widget.NewHBox(radio1)

	// 下拉
	select1 := widget.NewSelect(strs, func(s string) {
		fmt.Println(s)
	})

	// 音量滑动
	slider := widget.NewSlider(0, 100)
	// 密码输入框
	passwd := widget.NewPasswordEntry()
	// 文本输入框
	ent := widget.NewEntry()
	// 获取文本框的内容
	bt2 := widget.NewButton("text_button", btn)

	// 添加
	group1.Append(ent)
	group2.Append(bt2)
	group1.Append(linkToBaidu)
	group1.Append(check1)
	group1.Append(box1)
	group1.Append(select1)
	group1.Append(slider)
	group1.Append(passwd)
	w.SetContent(fyne.NewContainerWithLayout(&Diagonal{}, group1, group2))
	w.ShowAndRun()
}

func btn() {
	// func() {
	// fmt.Println(ent.Text)
	// }
	fmt.Println("aaa")
}
