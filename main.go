package main

import (
	"fmt"
	"gui/examples"

	// "gui/spider"
	"net/url"
)

func main() {

	examples.NewWindow()
	fmt.Println(url.QueryEscape("你好"))
	// spider.ProcessXiaohuoshuan()
	// spider.ProcessXiaofangshuan()
	// spider.ProcessShuihe()
	// spider.ProcessShuichi()
	// spider.ProcessQushuimatou()
}

// http://gorm.book.jasperxu.com/crud.html#u
