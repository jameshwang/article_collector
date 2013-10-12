package controllers

import "github.com/robfig/revel"

type Test struct {
	*revel.Controller
}

func (c Test) Index() revel.Result {
	return c.Render()
}
