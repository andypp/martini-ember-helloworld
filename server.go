package main

import (
  "github.com/codegangsta/martini"
  "github.com/codegangsta/martini-contrib/render"
  "html/template"
  "time"
)

type Person struct {
  Name string
}

func CurrentYear() int {
  return time.Now().Year()
}

func main() {
  m := martini.Classic()
  m.Use(render.Renderer(render.Options{
    Layout: "layout",
    Funcs: []template.FuncMap{
      template.FuncMap{"currentYear": CurrentYear},
    },
    Delims: render.Delims{Left: "{{%", Right: "%}}"},
  }))
  m.Get("/", func(r render.Render) {
    p := Person{Name: "lalalla"}
    r.HTML(200, "hello", p)
  })
  m.Run()
}
