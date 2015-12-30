package main

import (
	"html/template"
	"log"
	"path"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
)

type (
	HTMLBinData struct {
		TemplateDir string
		Template    map[string]*template.Template
	}
)

func (r HTMLBinData) Instance(name string, data interface{}) render.Render {
	var t *template.Template
	if gin.Mode() == "debug" {
		filename := path.Join(r.TemplateDir, name)
		name = filepath.Base(name)
		t = template.Must(template.New(name).ParseFiles(filename))
	} else {
		if tmpl, ok := r.Template[name]; ok {
			t = tmpl
		} else {
			filename := path.Join(r.TemplateDir, name)
			templateBytes, err := Asset(filename)
			if err != nil {
				log.Fatal(err)
			}
			tmpl = template.New(name)
			_, err = tmpl.Parse(string(templateBytes))
			if err != nil {
				log.Fatal(err)
			}
			r.Template[name] = tmpl
			t = tmpl
		}
	}

	return render.HTML{
		Template: t,
		Name:     name,
		Data:     data,
	}
}

func NewHTMLBinData(tempDir string) *HTMLBinData {
	return &HTMLBinData{
		TemplateDir: tempDir,
		Template:    map[string]*template.Template{},
	}
}

func DefaultHTMLBinData() *HTMLBinData {
	return NewHTMLBinData("templates")
}
