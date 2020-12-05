package v1

const MYSQL_TEMPLATE = `DROP TABLE IF EXISTS ` + "`{{ .TableName }}`;" + `
` + "CREATE TABLE `{{ .TableName }}` ( " + `
` + "`id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键'," + `
{{ range  $value := .datas }}{{$b := "int"}}
` + "`{{ $value.json }}` {{ $value.mysql_type }}  not null default {{if eq $value.kind $b}}0{{else}}''{{end}}  COMMENT '{{ $value.comment }}'," + `{{end}}

` + "`create_time` varchar(32)  not null DEFAULT ''  COMMENT '创建时间'," + `
` + "`update_time` varchar(32)  not null DEFAULT ''  COMMENT '更新时间'," + `
` + "`is_delete` tinyint(1) DEFAULT 0," + `
` + `
` + "PRIMARY KEY (`id`)" + `
` + `) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='{{ .TableName }}'

-- COLLATE 排序规则
-- utf8mb4_unicode_ci utf8mb4_general_ci 都行 英文字母大小写不敏感
-- utf8mb4_bin 的比较方法其实就是直接将所有字符看作二进制串，然后从最高位往最低位比对。所以很显然它是区分大小写的
`
