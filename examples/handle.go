package examples

import (
	"fmt"
	"gui/utils"
	gtpl "gui/v1"
	"os"
	"text/template"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

var (
	kinds        = []string{"string", "int", "uint"}
	boxChan      = make(chan *widget.Box)
	results      = make([]map[string]interface{}, 0)
	count        = 0
	tableNamePtr *widget.Entry
)

// NewWindow ...
func NewWindow() {
	// 初始化一个app
	a := app.New()
	w := a.NewWindow("w1")
	w.SetTitle("🌵🌴🌲🌳👀🔞🈲👀🌳🌲🌴🌵")

	box1 := widget.NewHBox()
	box2 := widget.NewHBox()

	tname := widget.NewEntry()
	tableNamePtr = tname
	labelTname := widget.NewLabel("table_name")
	box1.Append(labelTname)
	box1.Append(tname)

	button := widget.NewButton("new item", addNewItem)
	button2 := widget.NewButton("gene", aggerateeAndGenerate)
	buttonExit := widget.NewButton("exit", func() {
		w.Close()
		os.Exit(0)
	})

	// button.
	box2.Append(button)
	box2.Append(button2)
	box2.Append(buttonExit)
	container := fyne.NewContainerWithLayout(&Diagonal{}, box1, box2)
	// 初始化一个布局容器

	// 从管道取box 往这个容器中放
	go func() {
		for {
			select {
			case box := <-boxChan:
				container.AddObject(box)
			default:
				time.Sleep(time.Millisecond * 200)
			}
		}
	}()

	w.SetContent(container)
	w.ShowAndRun()
}

func newBox() {
	count++
	entry := make(map[string]interface{})

	name := widget.NewEntry()
	label := widget.NewLabel("name")

	entry["name_ptr"] = name // 动态获取不到 存入指针
	entry["name"] = ""
	label2 := widget.NewLabel("kind")
	select1 := widget.NewSelect(kinds, func(s string) {
		entry["kind"] = s
		fmt.Println(s)
	})

	box := widget.NewHBox() // vbox 内容上下排列。// hbox 内容左右排列
	box.Append(label)
	box.Append(name)
	box.Append(label2)
	box.Append(select1)

	// comment
	comment := widget.NewEntry()
	label3 := widget.NewLabel("comment")

	entry["comment_ptr"] = comment // 动态获取不到 存入指针
	entry["comment"] = ""
	box.Append(label3)
	box.Append(comment)
	results = append(results, entry)

	// res := make(map[int]interface{})
	// res[count] = entry
	// results = append(results, res)
	boxChan <- box
}

func addNewItem() {
	// TODO 这里按下了只有鼠标点击了拿开后才会触发的bug
	// 按下多次后新增多个

	fmt.Println("add new item")
	newBox()

}

// 获取所有的字段和对应的类型 并渲染模版
func aggerateeAndGenerate() {
	// 获取所有的entry的text值
	length := len(results)
	if length <= 0 {
		return
	}
	// 获取name的值
	for i := 0; i < length; i++ {
		item := results[i]
		namePtr := item["name_ptr"].(*widget.Entry)
		item["name"] = namePtr.Text
		commentPtr := item["comment_ptr"].(*widget.Entry)
		item["comment"] = commentPtr.Text
	}
	fmt.Println(results)
	ModelName := utils.Marshal(tableNamePtr.Text)
	TableName := utils.UnMarshal(tableNamePtr.Text)
	//    整合自动生成sql的代码
	go renderController(ModelName, TableName)
	go renderService(ModelName, TableName)
	go renderGorm(ModelName, TableName)
	go renderGormTest(ModelName, TableName)
	go renderMysql(ModelName, TableName)
}

func renderMysql(m, t string) {
	// tpl, err := template.ParseFiles("tpl/gorm.tmpl")
	tpl, err := template.New("render_mysql").Parse(gtpl.MYSQL_TEMPLATE)

	if err != nil {
		fmt.Println("err-", err)
		// return
	}
	content := make(map[string]interface{})
	content["ModelName"] = m
	content["TableName"] = t
	content["now"] = time.Now().Format("2006-01-02 15:04:05")
	// 处理results
	for _, v := range results {
		v["json"] = utils.UnMarshal(v["name"].(string))
		mysqlType := v["kind"].(string)
		switch mysqlType {
		case "string":
			mysqlType = "varchar(255)"
		case "int":
			mysqlType = "int(11) "
		case "uint":
			mysqlType = " tinyint(1) "
		}
		// "string", "int", "uint"
		v["mysql_type"] = mysqlType
		// v["field"] = utils.Marshal(v["name"].(string))
	}
	content["datas"] = results
	// 输出到控制台
	// tpl.ExecuteTemplate(os.Stdout, "test.tmpl", content)
	// 写入到文件
	f, err := os.Create(fmt.Sprintf("./%s_mysql.sql", t))
	if err != nil {
		fmt.Println("create file: ", err)
		return
	}
	err = tpl.Execute(f, content)
	if err != nil {
		fmt.Print("execute: ", err)
		return
	}
	f.Close()
}
func renderGormTest(m, t string) {
	// tpl, err := template.ParseFiles("tpl/gorm_test.tmpl")
	tpl, err := template.New("gorm_test").Parse(gtpl.TEMPLATE_TEST_GORM)
	if err != nil {
		fmt.Println("err-", err)
		// return
	}
	content := make(map[string]interface{})
	content["ModelName"] = m
	content["TableName"] = t
	content["now"] = time.Now().Format("2006-01-02 15:04:05")
	// 输出到控制台
	// tpl.ExecuteTemplate(os.Stdout, "test.tmpl", content)
	// 写入到文件
	f, err := os.Create(fmt.Sprintf("./%s_gorm_test.go", t))
	if err != nil {
		fmt.Println("create file: ", err)
		return
	}
	err = tpl.Execute(f, content)
	if err != nil {
		fmt.Print("execute: ", err)
		return
	}
	f.Close()
}

func renderGorm(m, t string) {
	// tpl, err := template.ParseFiles("tpl/gorm.tmpl")
	tpl, err := template.New("gorm_test").Parse(gtpl.GORM_TEMPLATE)

	if err != nil {
		fmt.Println("err-", err)
		// return
	}
	content := make(map[string]interface{})
	content["ModelName"] = m
	content["TableName"] = t
	content["now"] = time.Now().Format("2006-01-02 15:04:05")
	// 处理results
	for _, v := range results {
		v["json"] = utils.UnMarshal(v["name"].(string))
		v["field"] = utils.Marshal(v["name"].(string))
	}
	content["datas"] = results
	// 输出到控制台
	// tpl.ExecuteTemplate(os.Stdout, "test.tmpl", content)
	// 写入到文件
	f, err := os.Create(fmt.Sprintf("./%s_gorm.go", t))
	if err != nil {
		fmt.Println("create file: ", err)
		return
	}
	err = tpl.Execute(f, content)
	if err != nil {
		fmt.Print("execute: ", err)
		return
	}
	f.Close()
}
func renderService(m, t string) {
	// tpl, err := template.ParseFiles("tpl/service.tmpl")
	tpl, err := template.New("service").Parse(gtpl.SERVICE)
	if err != nil {
		fmt.Println("err-", err)
		// return
	}
	content := make(map[string]interface{})
	content["ModelName"] = m
	content["now"] = time.Now().Format("2006-01-02 15:04:05")
	// 输出到控制台
	// tpl.ExecuteTemplate(os.Stdout, "controller.tmpl", content)
	// 写入到文件
	f, err := os.Create(fmt.Sprintf("./%s_service.go", t))
	if err != nil {
		fmt.Println("create file: ", err)
		return
	}
	err = tpl.Execute(f, content)
	if err != nil {
		fmt.Print("execute: ", err)
		return
	}
	f.Close()
}
func renderController(m, t string) {
	// tpl, err := template.ParseFiles("tpl/controller.tmpl")
	tpl, err := template.New("controller").Parse(gtpl.CONTROLLER)
	if err != nil {
		fmt.Println("err-", err)
		// return
	}
	content := make(map[string]interface{})
	content["ModelName"] = m
	content["now"] = time.Now().Format("2006-01-02 15:04:05")
	// 输出到控制台
	// tpl.ExecuteTemplate(os.Stdout, "controller.tmpl", content)
	// 写入到文件
	f, err := os.Create(fmt.Sprintf("./%s_controller.go", t))
	if err != nil {
		fmt.Println("create file: ", err)
		return
	}
	err = tpl.Execute(f, content)
	if err != nil {
		fmt.Print("execute: ", err)
		return
	}
	f.Close()
}
