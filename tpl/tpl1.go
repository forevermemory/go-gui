package tpl

import (
	"os"
	"text/template"
)

type Friend struct {
	Fname string
}
type Person struct {
	UserName string
	Emails   []string
	Friends  []*Friend
}

func Tpl1() {

	f1 := Friend{Fname: "xiaofang"}
	f2 := Friend{Fname: "wugui"}
	t := template.New("test")
	t = template.Must(t.Parse(
		`hello {{.UserName}}!
			{{ range .Emails }}
				an email {{ . }}
			{{- end }}
			{{ with .Friends }}
			{{- range . }}
				my friend name is {{.Fname}}
			{{- end }}
		{{ end }}`))
	p := Person{UserName: "longshuai",
		Emails:  []string{"a1@qq.com", "a2@gmail.com"},
		Friends: []*Friend{&f1, &f2}}
	t.Execute(os.Stdout, p)

}

func Tpl2() {
	tpl, err := template.ParseFiles("tpl/test.tmpl")
	if err != nil {
		return
	}
	result := make(map[string]interface{})
	result["name"] = "hello"
	result["ModelName"] = "Nihao"
	result["TableName"] = "Nihao"

	tpl.ExecuteTemplate(os.Stdout, "test.tmpl", result)
}
