package template

import (
    "html/template"
    "io"
    "github.com/labstack/echo/v4"
    "fmt"

)
type Templates struct {
    templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error{
    return t.templates.ExecuteTemplate(w, name, data)
}

func NewTemplate() *Templates {
    pattern := "views/*.html"
    fmt.Println("Resolved pattern:", pattern)
    template.ParseGlob(pattern)
    return &Templates{
        templates: template.Must(template.ParseGlob(pattern)),
    }
}
