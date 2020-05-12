package language

import (
	"github.com/HongJaison/go-admin/modules/language"
	"html/template"
)

func Get(key string) string {
	return language.GetWithScope(key, "filemanager")
}

func GetHTML(key string) template.HTML {
	return template.HTML(language.GetWithScope(key, "filemanager"))
}
