package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"path"

	"github.com/go-chi/chi"
	// "github.com/go-chi/chi/middleware"

	LT "../like_trello"
)

type Profile struct {
	Name    string
	Hobbies []string
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Hi there, I love %s!\n", r.URL.Path[1:])
	fp := path.Join("templates", "index.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// w.Header().Set("Content-Type", "application/json")
	// d := fmt.Sprintf("<p>Go to <a>%s/boards</a> if want to look at list of boards!</p", r.Host)
	// w.Write([]byte(fmt.Sprintf("<html><body><p>Index page</p>%s</body></html>", d)))
	// profile := Profile{"Alex", []string{"snowboarding", "programming"}}
	if err := tmpl.Execute(w, r.Host); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func boardList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("List of all boards")

	// profile := Profile{"Alex", []string{"snowboarding", "programming"}}
	c := LT.Column{Name: "To do"}
	c2 := LT.Column{Name: "Doing"}
	c3 := LT.Column{Name: "To be Reviewed"}
	c4 := LT.Column{Name: "To be Tested"}
	b := LT.Board{Name: "The best board"}
	columns := []*LT.Column{&c, &c2, &c3, &c4}
	bm := LT.NewBoardManager(&b, columns...)

	js, err := json.Marshal(bm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.MethodFunc("GET", "/", index)
	r.MethodFunc("GET", "/boards", boardList)
	// r.MethodFunc("POST", "/{cloud}", h.api_handler)
	return r
}
