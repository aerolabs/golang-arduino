package main

import (
	"fmt"
	//	"github.com/gorilla/mux"
	//	"io/ioutil"
	"net/http"
	//	"text/template"
)

type Config struct {
	Port   string
	Host   string
	Public string
	Root   string
}

var (
	//	Templates *template.Template
	config = Config{
		Port:   "8080",
		Public: "file:///c:AWS_Helice/public/",
		Root:   "c:/AWS_Helice/",
	}
)

func init() {
	//	LoadConfig()
	//	LoadTemplates()
}

/*
func LoadTemplates() {
	fmt.Println("------ Load Templates ")

	//defines the template dir pattern to look for templates
	templatesDirPattern := config.Root + "templates/*.html"
	//uses ParseGlob function to get all templates in this dir
	Templates = template.Must(template.ParseGlob(templatesDirPattern))

	templateList := Templates.Templates()
	//	nt := len(templateList)
	for j, _ := range templateList {
		fmt.Println(templateList[j].Name())
	}

	fmt.Println("------ Finished Loading Templates")
	fmt.Println("------")
}
*/

func HomeHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "<html>HELICE V1.0</html>")
}

func NewPropHandler(w http.ResponseWriter, req *http.Request) {

	fmt.Println("New Prop")
	fmt.Fprintln(w, "New Propeller")
}

func PropListHandler(w http.ResponseWriter, req *http.Request) {

	fmt.Println("PropListHandler")
	fmt.Fprintln(w, "PropListHandler")
}

func main() {

	fmt.Println("Helice Server Running v0.2")

	//	muxrouter := mux.NewRouter()

	//home
	//	muxrouter.HandleFunc("/", HomeHandler)

	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/new", NewPropHandler)
	http.HandleFunc("/lista", PropListHandler)
	http.HandleFunc("/send", SendToArduinoHandler)
	http.HandleFunc("/receive", ReceiveFromArduinoHandler)

	//	fs := http.FileServer(http.Dir("public"))
	//	http.Handle("/public/", http.StripPrefix("/public/", fs))

	//	fs := http.FileServer(http.Dir("public"))
	//  	http.Handle("/public", fs)

	s := &http.Server{
		Addr: ":8080",
	}
	s.ListenAndServe()

	//	http.Handle("/", muxrouter)

	//	siteport := ":7777"
	//	http.ListenAndServe(siteport, nil)

	fmt.Println("End Helice Server")
}
