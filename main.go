package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/uguangtian/blog/controllers"
	"github.com/uguangtian/blog/models"
)

func main() {
	fmt.Println("Running")
	http.HandleFunc("/", controllers.DefaultHandle)
	http.HandleFunc("/letter", controllers.LetterHandle)
	http.HandleFunc("/blog", controllers.BlogHandle)
	http.HandleFunc("/post", controllers.PostHandle)
	// http.HandleFunc("/views/", controllers.StaticsHandle)
	// mux := http.NewServeMux()
	// wd, _ := os.Getwd()
	// mux.Handle("/view", http.StripPrefix("view/statics/css", http.FileServer(http.Dir(wd))))
	// http.Handle("/view/statics", http.FileServer(http.Dir("/views/statics")))
	// r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.HandleFunc("/views/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})
	b := models.CheckConnect()
	if b {
		fmt.Println("ok")
	} else {

		fmt.Println("Error")
	}
	log.Fatal(http.ListenAndServe(":8080", nil))
}
