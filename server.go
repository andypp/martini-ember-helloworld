package main

import (
    "github.com/codegangsta/martini"
    "github.com/codegangsta/martini-contrib/render"
    "html/template"
    "strings"
)

type Person struct {
    Name string
}

func ToUpper(args ...interface{}) string {
    return strings.ToUpper(args[0].(string))
}

func main() {
    m := martini.Classic()
    m.Use(render.Renderer(render.Options{
        Layout: "layout",
        Funcs: []template.FuncMap{
            template.FuncMap{"toUpper": ToUpper},
        },
        Delims: render.Delims{Left: "{{%", Right: "%}}"},
    }))
    m.Get("/", func(r render.Render) {
        p := Person{Name: "lalalla"}
        r.HTML(200, "hello", p)
    })
    m.Run()
}
