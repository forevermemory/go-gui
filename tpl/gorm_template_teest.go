package tpl

const TEMPLATE_TEST_GORM = `
package dbaccess

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"testing"
)
//执行整个测试文件
// go test -v {{ .TableName }}_gorm_test.go {{ .TableName }}_gorm.go
// //若文件存在多级依赖，可以直接在包目录下执行go test，运行包下所有的测试文件
// //执行测试文件中的指定方法
//  go test -v -count=1  {{ .TableName }}_gorm_test.go {{ .TableName }}_gorm.go -test.run TestMigration

func TestMigration(t *testing.T) {
	db, err := gorm.Open("mysql", "root:123456@(101.133.168.208)/codenai?charset=utf8&parseTime=True&loc=Local")
	// db, err := gorm.Open("mysql", "root:123456@(localhost)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("err--", err)
		panic(err)
	}
	defer db.Close()
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&{{ .ModelName }}{})
	db.Model(&{{ .ModelName }}{}).AddIndex("idx_{{ .TableName }}", "id")
}

`
