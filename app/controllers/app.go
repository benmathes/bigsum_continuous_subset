package controllers

import "github.com/robfig/revel"

type App struct {
  *revel.Controller
}

func (c App) Index() revel.Result {
  var m = map[string]string{
    "a": "aaa",
    "b": "bbb",
  }
  return c.Render(m)
}
