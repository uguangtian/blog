package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/labstack/gommon/log"
)

type Recipient struct {
	Name, Gift string
	Attended   bool
}

func checkError(err error) {
	if err != nil {
		log.Error(err)
	}
}
func LetterHandle(w http.ResponseWriter, r *http.Request) {
	// t := template.Must(template.New("letter").ParseFiles("index.html", "index.html"))
	// t.Execute(w, nil)
	// w.Header().Set("Content-Type", "text/html;charset=utf-8")
	var recipients = Recipient{"Aunt mildred", "bone china tes set", true}
	t := template.Must(template.ParseFiles("views/letter.html", "views/letter.html"))
	err := t.Execute(w, recipients)
	checkError(err)
	// fmt.Fprintf(w, "Work")
}

func PostHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html;charset=utf-8")
	switch r.Method {
	case "POST":
		{
			r.ParseForm()
			// fmt.Fprintf(w, "%s", r.PostForm)
			if r.PostForm != nil {

				fmt.Fprintf(w, "%s", r.PostFormValue("main-post-content"))
			}

		}

	case "DELETE":
		fmt.Fprintf(w, "DELETE")
	default:
		{
			t := template.Must(template.ParseFiles("views/post.html", "views/tpl_header.html", "views/tpl_footer.html", "views/tpl_base.html"))
			err := t.Execute(w, nil)
			checkError(err)

		}
	}
}
func BlogHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "blog work")
}

func DefaultHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	t := template.Must(template.ParseFiles("views/index.html", "views/tpl_header.html", "views/tpl_footer.html", "views/tpl_base.html"))
	err := t.Execute(w, nil)
	checkError(err)
}

// Handle the static file
func StaticsHandle(w http.ResponseWriter, r *http.Request) {
	wd, err := os.Getwd()
	checkError(err)
	http.StripPrefix("/views/", http.FileServer(http.Dir(wd)))
}
