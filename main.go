package main

import (
	"flag"
	"fmt"
	"log"

	// LT "./like_trello"
	"net/http"

	H "./handlers"
)

// func Handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Printf("Hi there, I love %s!", r.URL.Path[1:])
// }

func main() {

	// fmt.Println("Trello")
	// c := LT.Column{Name: "To do"}
	// c2 := LT.Column{Name: "Doing"}
	// c3 := LT.Column{Name: "To be Reviewed"}
	// c4 := LT.Column{Name: "To be Tested"}
	// b := LT.Board{Name: "The best board"}
	// columns := []*LT.Column{&c, &c2, &c3, &c4}
	// bm := LT.NewBoardManager(&b, columns...)
	// fmt.Println(b)
	// fmt.Println(bm, "bm")

	// bm.CreateTask(&c, "Init task$$$$$", &LT.First)
	// bm.CreateTask(&c, "Second task$$$$$", &LT.First)
	// bm.CreateTask(&c, "Third task$$$$$", &LT.First)
	// bm.CreateTask(&c, "Fourth task$$$$$", &LT.First)
	// bm.CreateTask(&c, "Fifth task$$$$$", &LT.First)

	// for _, t := range c.Tasks {
	// 	fmt.Println(*t, "tasks - columns")
	// }
	// fmt.Println("--------------------")
	// bm.MoveTop(&c, c.Tasks[4])
	// for _, t := range c.Tasks {
	// 	fmt.Println(*t, "tasks - columns")
	// }
	// bm.MoveToColumn(c.Tasks[4], &c2)
	// // bm.MoveToColumn(c.Tasks[2], &c2)
	// // fmt.Println(*c.Tasks[1], "tasks - columns")
	// fmt.Println(len(c2.Tasks), c2.Tasks[0].Priority, "c2")
	// // fmt.Println(len(c2.Tasks), c2.Tasks[1].Priority, "c2")

	// fmt.Println(len(c.Tasks), "c")

	// bm.AddComment(c.Tasks[0], "What is up?")
	// fmt.Println(c.Tasks[0].Comments[0], "c")

	// // bm2 := LT.NewBoardManager(&b)
	// // fmt.Println(bm2.Columns[0], "bm2")
	var httpHost = flag.String("http.host", "localhost:8081", "Host+port.")

	flag.Parse()
	fmt.Println("Server started on", *httpHost)

	server := &http.Server{
		Addr:    *httpHost,
		Handler: H.NewRouter(),
	}

	if err := server.ListenAndServe(); err != nil {
		//str_err := fmt.Sprintf("Error during starting server: %s, %v\n", *httpHost, err)
		//custom_logger.WriteErrorLog(&str_err)
		log.Fatal("Could not start listening for HTTP messages.")
	}
}
