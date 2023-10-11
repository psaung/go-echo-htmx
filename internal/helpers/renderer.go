package helpers

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
)

// reference from https://gist.github.com/rakd/5073f57e5053ce443cd8de070e623d63

// Template ...
type Template struct {
	templates map[string]*template.Template
}

// Add ...
func (t Template) Add(name string, tmpl *template.Template) {
	if tmpl == nil {
		panic("template can not be nil")
	}
	if len(name) == 0 {
		panic("template name cannot be empty")
	}
	t.templates[name] = tmpl
}

// Render ...
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if _, ok := t.templates[name]; ok == false {
		// not such view
		return fmt.Errorf("no such view. (%s)", name)
	}
	return t.templates[name].Execute(w, data)
}

// NewTemplate creates a new template
func NewTemplateRenderer(e *echo.Echo, templatesDir string) {
	ext := ".html"
	ins := Template{
		templates: map[string]*template.Template{},
	}

	layout := templatesDir + "layouts/master" + ext
	adminLayout := templatesDir + "layouts/admin" + ext

	_, err := os.Stat(layout)
	if err != nil {
		log.Panicf("cannot find %s", layout)
		os.Exit(1)
	}
	_, err = os.Stat(adminLayout)
	if err != nil {
		log.Printf("cannot find %s", adminLayout)
		os.Exit(1)
	}

	partials, err := filepath.Glob(templatesDir + "partials/" + "*" + ext)
	if err != nil {
		log.Print("cannot find " + templatesDir + "partials/" + "*" + ext)
		os.Exit(1)
	}

	funcMap := template.FuncMap{
		"safehtml": func(text string) template.HTML { return template.HTML(text) },
	}

	views, _ := filepath.Glob(templatesDir + "**/*" + ext)
	for _, view := range views {
		dir, file := filepath.Split(view)
		dir = strings.Replace(dir, templatesDir, "", 1)
		file = strings.TrimSuffix(file, ext)
		renderName := dir + file

		if strings.Contains(renderName, "partials") || renderName == "layouts/master" ||
			renderName == "layouts/admin" {
			// Do Nothing
		} else if strings.Contains(renderName, "admin") {
			tmplfiles := append([]string{adminLayout, view}, partials...)
			tmpl := template.Must(
				template.New(filepath.Base(adminLayout)).Funcs(funcMap).ParseFiles(tmplfiles...),
			)
			ins.Add(renderName, tmpl)
			log.Printf("renderName: %s, layout: %s", renderName, adminLayout)
		} else if strings.Contains(renderName, "pages") {
			tmplfiles := append([]string{layout, view}, partials...)
			tmpl := template.Must(template.New(filepath.Base(layout)).Funcs(funcMap).ParseFiles(tmplfiles...))
			ins.Add(renderName, tmpl)
			log.Printf("renderName: %s, layout: %s", renderName, layout)
		} else {
			tmplfiles := append([]string{view}, partials...)
			tmpl := template.Must(template.New(filepath.Base(view)).Funcs(funcMap).ParseFiles(tmplfiles...))
			ins.Add(renderName, tmpl)
			log.Printf("renderName: %s, layout: %s", renderName, "none")
		}
	}
	e.Renderer = &ins
}
