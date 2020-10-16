package controllers

import (
	"fmt"
	"html"
	"net/http"
)

type HomeController struct{}

func (c *HomeController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
