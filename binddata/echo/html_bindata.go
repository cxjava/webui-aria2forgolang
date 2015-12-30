package main

import (
	"html/template"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
)

type (
	HTMLBinData struct {
		TemplateDir string
		Template    map[string]*template.Template
	}
)

const (
	ENV_ECHO_MODE = "ECHO_MODE"

	DebugMode   string = "debug"
	ReleaseMode string = "release"
	TestMode    string = "test"
)

var modeName string = DebugMode

func init() {
	mode := os.Getenv(ENV_ECHO_MODE)
	if len(mode) == 0 {
		modeName = DebugMode
	} else {
		modeName = mode
	}
}

func (r HTMLBinData) Render(w io.Writer, name string, data interface{}) error {
	var t *template.Template
	if modeName == "debug" {
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
				return err
			}
			tmpl = template.New(name)
			_, err = tmpl.Parse(string(templateBytes))
			if err != nil {
				log.Fatal(err)
				return err
			}
			r.Template[name] = tmpl
			t = tmpl
		}
	}
	return t.ExecuteTemplate(w, name, data)
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
