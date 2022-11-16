package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func renderTemplate(w http.ResponseWriter, tmpl string) {
	str := fmt.Sprintf("./templates/" + tmpl)
	parsedTemp, _ := template.ParseFiles(str)
	err := parsedTemp.Execute(w, parsedTemp)
	if err != nil {
		fmt.Println(err)
	}
}
