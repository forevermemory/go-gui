package tpl

const GORM_TEMPLATE = `

package dbaccess

/*
date:{{ .now }}
*/

// {{ .ModelName }} {{ .ModelName }}
type {{ .ModelName }} struct {
	ID int ` + "`json:\"id\" form:\"id\" gorm:\"column:id;primary_key;auto_increment;comment:'主键'\"`" + `
	{{ range  $value := .datas }}{{ $value.field }}	{{ $value.kind }} ` + "`json:\"{{ $value.json }}\" form:\"{{ $value.json }}\" gorm:\"column:{{ $value.json }};comment:'{{ $value.comment }}'\"`" + `
	{{end}}
	IsDelete uint ` + "`json:\"-\" form:\"-\" gorm:\"column:is_delete;default:0\"`" + ` // 0 未删除
	PageNo   int   ` + " `json:\"page\" form:\"page\" gorm:\"-\"`" + `
	PageSize int    ` + "`json:\"page_size\" form:\"page_size\" gorm:\" - \"`" + `
}

// TableName 表名
func (o *{{ .ModelName }}) TableName() string {
	return "{{ .TableName }}"
}

// Delete{{ .ModelName }} 根据id删除
func Delete{{ .ModelName }}(id int, tx ...*gorm.DB) error {
	db := OpenGorm()
	if len(tx) > 0 {
		db = tx[0]
	}
	return db.Table("{{ .TableName }}").Where("id = ?", id).Update("is_delete", 1).Error
}

// Get{{ .ModelName }}ByID 根据id查询一个
func Get{{ .ModelName }}ByID(id int) (*{{ .ModelName }}, error) {
	db := OpenGorm()
	o := &{{ .ModelName }}{}
	err := db.Table("{{ .TableName }}").Where("is_delete = 0").Where("id = ?", id).First(o).Error
	return o, err
}

// Add{{ .ModelName }} 新增
func Add{{ .ModelName }}(o *{{ .ModelName }}, tx ...*gorm.DB) (*{{ .ModelName }},error) {
	db := OpenGorm()
	if len(tx) > 0 {
		db = tx[0]
	}
	err := db.Create(o).Error
	return o,err
}

// Update{{ .ModelName }} 修改
func Update{{ .ModelName }}(o *{{ .ModelName }} , tx ...*gorm.DB) (*{{ .ModelName }},error) {
	db := OpenGorm()
	if len(tx) > 0 {
		db = tx[0]
	}
	err := db.Table("{{ .TableName }}").Where("id=?", o.ID).Update(o).First(o).Error
	return o,err
}

// List{{ .ModelName }} 分页条件查询
func List{{ .ModelName }}(o *{{ .ModelName }}) ([]*{{ .ModelName }}, error) {
	db := OpenGorm()
	res := make([]*{{ .ModelName }}, 0)
	err := db.Table("{{ .TableName }}").Where("is_delete = 0").Where(o).Offset((o.PageNo - 1) * o.PageSize).Limit(o.PageSize).Find(&res).Error
	return res, err
}

// Count{{ .ModelName }} 条件数量
func Count{{ .ModelName }}(o *{{ .ModelName }}) (int64, error) {
	db := OpenGorm()
	var count int64
	err := db.Table("{{ .TableName }}").Where("is_delete = 0").Where(o).Count(&count).Error
	return count, err
}



`
