package handlers

import (
	Broker "dadJokesApiApp/internal/jokes"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// this struct is declaring the fields (data types) that belong to this struct
// for home handler (home handler model)
type homeData struct {
	Message string
}

func loadMessageInIndexHTML(w http.ResponseWriter, joke string) {
	// template from parsing the html file
	// the relative path from server.go, where everything starts execution
	templateObject, err := template.ParseFiles("./static/index.html")

	if err != nil {
		fmt.Print(err.Error())
		return
	}

	// home data is what we can render on the html page; passing it to the front end
	hd := homeData{
		Message: joke,
	}

	// specifying that what we are sending is html so browsers can render it
	w.Header().Set("Content-Type", "text/html")

	err = templateObject.Execute(w, hd)

	if err != nil {
		fmt.Print(err.Error())
		return
	}
}

// TODO: finish
func HomeHandler(w http.ResponseWriter, r *http.Request) {

	loadMessageInIndexHTML(w, "hello world, 你好世界，hallo welt")

}

func GetJokeHandler(w http.ResponseWriter, r *http.Request) {

	// if the path isn't right
	if r.URL.Path != "/getJoke" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	// if the method isn't right (must be GET for the API)
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	fmt.Println("Calling Dad Jokes API...")

	// using the new Broker object,
	// gets the response object (JokeStructure struct from the client file)
	responseObject, err := Broker.BrokerConstructor().RetrieveJoke()

	if err != nil {
		fmt.Print(err.Error())
		return
	}

	dadJoke := responseObject.Joke

	if jokeStatus := responseObject.Status; jokeStatus == 200 {
		// fmt.Fprint(w, "GET successful!\n")
		// fmt.Fprint(w, "Dad joke: "+dadJoke)
		fmt.Println("GET successful!")
		loadMessageInIndexHTML(w, "Dad joke: "+dadJoke)

	} else {
		fmt.Println("Failed! Status: " + fmt.Sprint(jokeStatus))
		loadMessageInIndexHTML(w, "Failed to get a random joke!")
	}

}

// finding a joke by putting joke id into url
func FindJokeByIDHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	jokeID := params["id"]

	// but i guess we cant even get here if it's blank in the url (from handlefunc)
	if jokeID = strings.TrimSpace(jokeID); jokeID == "" {
		http.Error(w, "Provided jokeID is blank or whitespace!", http.StatusNotFound)
		return
	}

	fmt.Println("Looking for joke in database...")

	responseObject, err := Broker.BrokerConstructor().FindJokeByID(jokeID)

	if err != nil {
		fmt.Println(err.Error())
		loadMessageInIndexHTML(w, "Failed to find joke!")
		// fmt.Print(w, "Failed to find joke!")
		return
	}

	dadJoke := responseObject.Joke

	// fmt.Fprint(w, "DB find by ID successful!\n")
	// fmt.Fprint(w, "Joke ID (for debugging): "+jokeID+"\n")
	// fmt.Fprint(w, "Dad joke: "+dadJoke)

	fmt.Println("DB find by ID successful!")
	fmt.Println("Joke ID (for debugging): " + jokeID)
	loadMessageInIndexHTML(w, "Dad joke: "+dadJoke)

}

// finding a joke by putting joke id into form
func FindJokeHandler(w http.ResponseWriter, r *http.Request) {

	// parsing input submitted from form
	r.ParseForm()
	// comes as a string array, need to convert to a string
	jokeIDArr := r.Form["jokeID"]
	jokeID := strings.Join(jokeIDArr, " ")

	if jokeID = strings.TrimSpace(jokeID); jokeID == "" {
		http.Error(w, "Provided jokeID is blank or whitespace!", http.StatusNotFound)
		return
	}

	fmt.Println("Looking for joke in database...")

	responseObject, err := Broker.BrokerConstructor().FindJokeByID(jokeID)

	if err != nil {
		fmt.Println(err.Error())
		loadMessageInIndexHTML(w, "Failed to find joke!")
		// fmt.Print(w, "Failed to find joke!")
		return
	}

	dadJoke := responseObject.Joke

	// fmt.Fprint(w, "DB find by ID successful!\n")
	// fmt.Fprint(w, "Joke ID (for debugging): "+jokeID+"\n")
	// fmt.Fprint(w, "Dad joke: "+dadJoke)

	fmt.Println("DB find by ID successful!")
	fmt.Println("Joke ID (for debugging): " + jokeID)
	loadMessageInIndexHTML(w, "Dad joke: "+dadJoke)

}

// func UpdateJokeByIDHandler(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	jokeID := params["id"]

// 	// but i guess we cant even get here if it's blank in the url (from handlefunc)
// 	if jokeID = strings.TrimSpace(jokeID); jokeID == "" {
// 		http.Error(w, "Provided jokeID is blank or whitespace!", http.StatusNotFound)
// 		return
// 	}

// 	fmt.Println("Looking for joke in database...")

// 	updatedJoke := "updated joke"

// 	err := Broker.BrokerConstructor().UpdateJokeByID(jokeID, updatedJoke)

// 	if err != nil {
// 		fmt.Println(err.Error())
// 		fmt.Fprint(w, "Failed to update joke!")
// 		return
// 	}

// 	fmt.Fprint(w, "DB update by ID successful!\n")
// 	fmt.Fprint(w, "Joke ID (for debugging): "+jokeID+"\n")

// }

func UpdateJokeHandler(w http.ResponseWriter, r *http.Request) {

	// parsing input submitted from form
	r.ParseForm()
	// comes as a string array, need to convert to a string
	jokeIDArr := r.Form["jokeID"]
	jokeID := strings.Join(jokeIDArr, " ")

	updatedJokeArr := r.Form["updatedJoke"]
	updatedJoke := strings.Join(updatedJokeArr, " ")

	if jokeID = strings.TrimSpace(jokeID); jokeID == "" {
		http.Error(w, "Provided jokeID is blank or whitespace!", http.StatusNotFound)
		return
	}

	if updatedJoke = strings.TrimSpace(updatedJoke); updatedJoke == "" {
		http.Error(w, "Provided updated joke is blank or whitespace!", http.StatusNotFound)
		return
	}

	fmt.Println("Looking for joke in database...")

	err := Broker.BrokerConstructor().UpdateJokeByID(jokeID, updatedJoke)

	if err != nil {
		fmt.Println(err.Error())
		loadMessageInIndexHTML(w, "Failed to update joke!")
		// fmt.Fprint(w, "Failed to update joke!")

		return
	}

	// fmt.Fprint(w, "DB update by ID successful!\n")
	// fmt.Fprint(w, "Joke ID (for debugging): "+jokeID+"\n")

	fmt.Println("DB update by ID successful!")
	loadMessageInIndexHTML(w, "DB update by ID successful!")
	fmt.Println("Joke ID (for debugging): " + jokeID)

}

func DeleteJokeByIDHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	jokeID := params["id"]

	// but i guess we cant even get here if it's blank in the url (from handlefunc)
	if jokeID = strings.TrimSpace(jokeID); jokeID == "" {
		http.Error(w, "Provided jokeID is blank or whitespace!", http.StatusNotFound)
		return
	}

	fmt.Println("Looking for joke in database...")

	err := Broker.BrokerConstructor().DeleteJokeByID(jokeID)

	if err != nil {
		fmt.Println(err.Error())
		loadMessageInIndexHTML(w, "Failed to delete joke!")
		// fmt.Fprint(w, "Failed to delete joke!")
		return
	}

	// fmt.Fprint(w, "DB delete by ID successful!\n")
	// fmt.Fprint(w, "Joke ID (for debugging): "+jokeID+"\n")

	fmt.Println("DB delete by ID successful!")
	loadMessageInIndexHTML(w, "DB delete by ID successful!")
	fmt.Println("Joke ID (for debugging): " + jokeID)

}

func DeleteJokeHandler(w http.ResponseWriter, r *http.Request) {

	// parsing input submitted from form
	r.ParseForm()
	// comes as a string array, need to convert to a string
	jokeIDArr := r.Form["jokeID"]
	jokeID := strings.Join(jokeIDArr, " ")

	if jokeID = strings.TrimSpace(jokeID); jokeID == "" {
		http.Error(w, "Provided jokeID is blank or whitespace!", http.StatusNotFound)
		return
	}

	fmt.Println("Looking for joke in database...")

	err := Broker.BrokerConstructor().DeleteJokeByID(jokeID)

	if err != nil {
		fmt.Println(err.Error())
		loadMessageInIndexHTML(w, "Failed to delete joke!")
		// fmt.Fprint(w, "Failed to delete joke!")
		return
	}

	// fmt.Fprint(w, "DB delete by ID successful!\n")
	// fmt.Fprint(w, "Joke ID (for debugging): "+jokeID+"\n")

	fmt.Println("DB delete by ID successful!")
	loadMessageInIndexHTML(w, "DB delete by ID successful!")
	fmt.Println("Joke ID (for debugging): " + jokeID)

}

// hello handler to play around with
func HelloHandler(w http.ResponseWriter, r *http.Request) {

	// if the path isn't right
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	// if the method isn't right
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	// otherwise print hello
	fmt.Fprintf(w, "Hello!")

}
