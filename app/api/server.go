package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	H "dadJokesApiApp/app/api/handlers"

	"github.com/gorilla/mux"
)

func main() {

	// in order to serve the static folder, create file server object.
	// fileServer := http.FileServer(http.Dir("./static"))
	// accepts a path and the fileserver
	// http.Handle("/", fileServer)

	// http://localhost:8080/hello
	// listening at this route and calls the hello handler function
	// http.HandleFunc("/hello", H.HelloHandler)

	// joke route
	// http.HandleFunc("/getJoke", H.GetJokeHandler)

	// ----------------------------------------//
	// Start of Gorilla Mux version of things: //

	r := mux.NewRouter()

	// path prefix takes the folder that is holding static files
	// strip prefix lets url not have "static" in it
	directory := "./static"
	pathPrefix := "/static/"
	r.PathPrefix(pathPrefix).Handler(http.StripPrefix(pathPrefix, http.FileServer(http.Dir(directory))))

	r.HandleFunc("/", H.HomeHandler)

	r.HandleFunc("/findJoke/{id}", H.FindJokeByIDHandler)
	r.HandleFunc("/findJoke", H.FindJokeHandler)

	// doesn't really make sense to not have a form for user to update a joke
	// r.HandleFunc("/updateJoke/{id}", H.UpdateJokeByIDHandler)
	r.HandleFunc("/updateJoke", H.UpdateJokeHandler)

	r.HandleFunc("/deleteJoke/{id}", H.DeleteJokeByIDHandler)
	r.HandleFunc("/deleteJoke", H.DeleteJokeHandler)

	r.HandleFunc("/getJoke", H.GetJokeHandler)

	srv := &http.Server{
		Addr: "0.0.0.0:8080",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}

	// End of Gorilla Mux things. //
	// ---------------------------//

	fmt.Printf("Starting server at port 8080\n")

	// starts server and checks if there is an error
	// if err := http.ListenAndServe(":8080", nil); err != nil {
	// 	log.Fatal(err)
	// }

	if err := srv.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
