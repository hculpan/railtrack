package handlers

import (
	"bytes"
	"html/template"

	"github.com/gin-gonic/gin"
)

func AboutHandler(c *gin.Context) {
	t, _ := template.ParseFiles("templates/about.html")
	var tpl bytes.Buffer
	t.Execute(&tpl, nil)
	c.Data(200, "text/html; charset=utf-8", tpl.Bytes())
}
