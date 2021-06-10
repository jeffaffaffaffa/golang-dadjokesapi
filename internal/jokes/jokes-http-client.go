package jokes

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	// the mongo agent
	mongoagent "gitlab.centene.com/agt/mongo"

	"github.com/gorilla/mux"
)

type Broker struct {
	http *http.Client

	// new connecting to mongo stuff
	mongo mongoagent.Agent

	// gorilla mux
	router *mux.Router
}

type JokeResponse struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

// for mongodb stuff; uses bson instead of json
// adding a time added field for the struct to use with the db
type JokeModel struct {
	ID        string    `bson:"id"`
	Joke      string    `bson:"joke"`
	Status    int       `bson:"status"`
	TimeAdded time.Time `bson:"time_added"`
}

// todo next: save jokes to database i.e. start making queries

// object constructor
func BrokerConstructor() Broker {

	agent, err := mongoagent.New(mongoagent.ConnectionInfo{
		// db name
		Database: "jokesDB",
		// collection name
		DefaultCollection: "jokesCollection",
		URL:               "mongodb://localhost:27017",
	})

	if err != nil {
		fmt.Print(err.Error())
		// return an empty broker
		return Broker{}
	}

	// the broker has ability to interact with mongodb, given that the config is set up properly
	return Broker{
		http:   http.DefaultClient,
		mongo:  agent,
		router: mux.NewRouter(),
	}
}

// receiver function
// so receiver functions are basically functions unique to the "class"
// ex. someJokeResInstance.ToModel()
func (jokeRes JokeResponse) ToModel() JokeModel {
	return JokeModel{
		ID:        jokeRes.ID,
		Joke:      jokeRes.Joke,
		Status:    jokeRes.Status,
		TimeAdded: time.Now().UTC(),
	}
}

// receiver function
// gets a joke from the API and stores it in the db if it isn't in the db yet
func (bkr Broker) RetrieveJoke() (JokeResponse, error) {
	// create request
	// send request
	// handle response

	request, err := http.NewRequest(http.MethodGet, "https://icanhazdadjoke.com/", nil)

	// if there is an error from creating http request
	if err != nil {

		// fmt.Errorf returns an error type
		return JokeResponse{}, fmt.Errorf("api failed to retrieve a joke, failed to create request; %w", err)
	}

	// adding appropriate headers for json
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/json")

	// sending http request
	// Do is within the bkr object
	response, err := bkr.http.Do(request)

	// if there is an error from request
	if err != nil {

		return JokeResponse{}, fmt.Errorf("api failed to retrieve a joke, failed to send request; %w", err)
	}

	// close response body at the very end
	defer response.Body.Close()

	var responseObject JokeResponse

	// Decode decodes the response body into the JokeResponse struct structure.
	// Decode also deserializes for us.
	if err := json.NewDecoder(response.Body).Decode(&responseObject); err != nil {

		return JokeResponse{}, fmt.Errorf("api failed to retrieve a joke, failed to decode the request body; %w", err)
	}

	// db queries
	// inserting a joke into the db if it isn't already in the db

	jokeID := responseObject.ID
	_, err = bkr.FindJokeByID(jokeID)

	if err != nil {
		// then check if the error contains the new sentinel error
		if errors.Is(err, ErrDocumentDoesntExist) {
			_, err := bkr.InsertJoke(responseObject.ToModel())
			fmt.Println(responseObject.ToModel())

			if err != nil {
				return JokeResponse{}, fmt.Errorf("api failed to insert joke in db; %w", err)
			}
		} else {
			return JokeResponse{}, fmt.Errorf("api failed to find joke in db; %w", err)
		}
	}

	// if all good, no errors, so error is nil and return the response object
	return responseObject, nil

}
