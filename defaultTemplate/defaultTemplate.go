package getdefaulttemplate

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

//Template to blah blah blah
type Template struct {
	templates *template.Template
}

//Render implments the "Render" interface provided by the "Template" struct
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}

//GetRenderer will return an instance of a template object contain the default go template engine.
func GetRenderer(templatesPathname string) *Template{
	return &Template{
		templates: template.Must(template.ParseGlob(templatesPathname)),
	}
}