package main

import ("fmt"
	"html/template"
	"net/http"
)

const portNumber = ":8080"

func Home(w http.ResponseWrite, r *http.Request0 {
	renderTemplate(w, "home.page.tmpl")
}
func About(w http.ResponseWriter, r *http.Request){
}
func renderTemplate(w http.ResponseWriter, templ string){
	parsedTemplate, _ := template.ParseFiles("./templates"+tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil{
		fmt.Println("Error parsing template: ", err)
	}
}
