package previewer

import (
	"github.com/HongJaison/filemanager/modules/language"
	"github.com/HongJaison/html"
	"github.com/h2non/filetype"
	"html/template"
	"io/ioutil"
	"path/filepath"
)

type Previewer interface {
	Preview(content []byte) template.HTML
}

func Preview(path string) (template.HTML, error) {
	buf, err := ioutil.ReadFile(filepath.FromSlash(path))

	if err != nil {
		return "", err
	}

	if filetype.IsImage(buf) {
		return image.Preview(buf), nil
	}

	ext := filepath.Ext(path)

	if IsCode(ext) {
		return NewCode(ext).Preview(buf), nil
	}

	return html.DivEl().SetClass("preview-content").
		SetStyle("margin", "20px auto 20px auto").
		SetStyle("width", "500px").
		SetStyle("text-align", "center").
		SetContent(html.H1(language.GetHTML("no supported"))).
		Get(), nil
}

var image = new(Image)

var codeExtensions = [...]string{
	".go", ".php", ".html", ".css", ".js", ".py", ".md",
	".c", ".cpp", ".java", ".sh", ".tmpl", ".mod", ".sum",
	".sql", ".json", ".txt",
}

func IsCode(ext string) bool {
	for _, e := range codeExtensions {
		if ext == e {
			return true
		}
	}
	return false
}
