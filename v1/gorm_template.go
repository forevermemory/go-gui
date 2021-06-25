package v1

const GORM_TEMPLATE = `

package dbaccess

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

/*
date:{{ .now }}

*/

// {{ .ModelName }} {{ .ModelName }}
type {{ .ModelName }} struct {
	ID int ` + "`json:\"id\" form:\"id\" gorm:\"column:id;primary_key;auto_increment;comment:'主键'\"`" + `
	{{ range  $value := .datas }}{{ $value.field }}	{{ $value.kind }} ` + "`json:\"{{ $value.json }}\" form:\"{{ $value.json }}\" gorm:\"column:{{ $value.json }};comment:'{{ $value.comment }}'\"`" + `
	{{end}}

	OtherDesc  string ` + "`json:\"other_desc\" form:\"other_desc\" gorm:\"column:other_desc;comment:'备注'\"`" + `
	CreateTime  string ` + "`json:\"-\" form:\"-\" gorm:\"column:create_time;comment:'创建时间'\"`" + `
	UpdateTime  string ` + "`json:\"-\" form:\"-\" gorm:\"column:update_time;comment:'更新时间'\"`" + `
	IsDelete uint ` + "`json:\"-\" form:\"-\" gorm:\"column:is_delete;default:0\"`" + ` // 0 未删除
	Page
}

// TableName 表名
func (o *{{ .ModelName }}) TableName() string {
	return "{{ .TableName }}"
}

// Delete{{ .ModelName }} 根据id删除
func Delete{{ .ModelName }}(id int, tx ...*gorm.DB) error {
	db := MYSQL
	if len(tx) > 0 {
		db = tx[0]
	}
	sql := "update {{ .TableName }} set is_delete = 1,update_time = ? where id = ?"
	err := db.Exec(sql, time.Now().Format("2006-01-02 15:04:05"), id).Error
	if err != nil {
        return err
    }
	return nil
}

// Get{{ .ModelName }}ByID 根据id查询一个
func Get{{ .ModelName }}ByID(id int) (*{{ .ModelName }}, error) {
	db := MYSQL
	o := &{{ .ModelName }}{}
	err := db.Table("{{ .TableName }}").Where("is_delete = 0").Where("id = ?", id).First(o).Error
	if err != nil {
        return nil, err
    }
	return o, nil
}

// Add{{ .ModelName }} 新增
func Add{{ .ModelName }}(o *{{ .ModelName }}, tx ...*gorm.DB) (*{{ .ModelName }},error) {
	db := MYSQL
	if len(tx) > 0 {
		db = tx[0]
	}
	o.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	o.UpdateTime = time.Now().Format("2006-01-02 15:04:05")
	err := db.Create(o).Error
	if err != nil {
        return nil, err
    }
	return o,nil
}

// Update{{ .ModelName }} 修改
func Update{{ .ModelName }}(o *{{ .ModelName }} , tx ...*gorm.DB) (*{{ .ModelName }},error) {
	db := MYSQL
	if len(tx) > 0 {
		db = tx[0]
	}
	o.UpdateTime = time.Now().Format("2006-01-02 15:04:05")
	err := db.Table("{{ .TableName }}").Where("id = ?", o.ID).Update(o).First(o).Error
	if err != nil {
        return nil, err
    }
	return o,nil
}

// List{{ .ModelName }} 分页条件查询
func List{{ .ModelName }}(o *CommonRequestParams) (*DataStore, error) {
	db := SQLITE
	var err error
	list := make([]*{{ .ModelName }}, 0)
	// 条件
	db = db.Debug().Table("{{ .TableName }}").Where("is_delete = 0").Where(o)
	if o.Status != "" {
		db = db.Where("status = ?", o.Status)
	}

	// 数量
	var total int
	if err = db.Count(&total).Error; err != nil {
		return nil, err
	}

	// 分页排序
	db = db.Offset((o.PageNo - 1) * o.PageSize)
	if len(o.Orders) > 0 {
		for _, val := range o.Orders {
			db = db.Order(fmt.Sprintf("%s %s", val.OrderField, val.OrderMethod))
		}
	}
	if err = db.Limit(o.PageSize).Find(&list).Error; err != nil {
		return nil, err
	}

	return &DataStore{Total: total, Data: list, TotalPage: (int(total) + o.PageSize - 1) / o.PageSize}, nil
}

// Count{{ .ModelName }} 条件数量
func Count{{ .ModelName }}(o *{{ .ModelName }}) (int64, error) {
	db := MYSQL
	var count int64
	err := db.Table("{{ .TableName }}").Where("is_delete = 0").Where(o).Count(&count).Error
	if err != nil {
        return 0, err
    }
	return count, err
}



` + `
// DataStore list的结构体
type DataStore struct {
	Total     int64     ` + "  `json:\"total\"` " + `
	TotalPage int       ` + "  `json:\"total_page\"`" + `
	Data      interface{}  ` + " `json:\"data\"`" + `
}

// Page 分页参数
type Page struct {
	PageNo   int ` + "  `gorm:\"-\" json:\"page_no,default=1,omitempty\" form:\"page_no,default=1\"` " + `
	PageSize int ` + "  `gorm:\"-\" json:\"page_size,default=10,omitempty\" form:\"page_size,default=10\"` " + `
}


Lon             string ` + " `json:\"lon\" form:\"lon\" gorm:\"column:lon\"` " + `
Lat             string ` + " `json:\"lat\" form:\"lat\" gorm:\"column:lat\"` " + `
GeoInfo         string  ` + " `json:\"geo_info\" form:\"geo_info\" gorm:\"column:GEO_INFO\"` " + `
CenterPoinit    string  ` + " `json:\"center_poinit\" form:\"center_poinit\" gorm:\"column:CENTER_POINIT\"` " + `
NodeCode        string  ` + " `json:\"node_code\" form:\"node_code\" gorm:\"column:NODE_CODE\"` " + `
XzqhID          int     ` + " `json:\"xzqh_id\" form:\"xzqh_id\" gorm:\"column:XZQH_ID\"` " + `
XzqhName        string  ` + " `json:\"xzqh_name\" form:\"xzqh_name\" gorm:\"column:XZQH_NAME\"` " + `


func PercentWrap(a string) string {
	return "%" + a + "%"
}

` + `

type OrderParams struct {
	OrderField  string ` + " `json:\"order_field\" form:\"order_field\" gorm:\"-\"` " + `
	OrderMethod string ` + " `json:\"order_method\" form:\"order_method\" gorm:\"-\"` " + `
} 
` + `

type CommonRequestParams struct {
	Page
	Orders      []*OrderParams ` + " `json:\"alarm_orders\" form:\"alarm_orders\" gorm:\"-\"` " + `

	Start     string ` + " `json:\"start\" form:\"start\" gorm:\"-\"` " + `
}
`
