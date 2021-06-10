package jokes

import (
	"errors"
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

// a cool sentinel error
var ErrDocumentDoesntExist error = errors.New("document not found")

// inserting a joke into db
func (bkr Broker) InsertJoke(joke JokeModel) (string, error) {

	if joke.Joke = strings.TrimSpace(joke.Joke); joke.Joke == "" {
		// errors.New() => custom error
		return "", errors.New("db error on insert: provided joke is invalid, joke cannot be empty or whitespace")
	}
	if joke.ID = strings.TrimSpace(joke.ID); joke.ID == "" {
		// errors.New() => custom error
		return "", errors.New("db error on insert: provided joke id is invalid, joke id cannot be empty or whitespace")
	}

	// payload has to be a pointer
	// first thing is the insert id, but we don't care for it.
	id, err := bkr.mongo.Insert("jokesCollection", &joke)

	if err != nil {
		return "", errors.New("db error on insert: failed to insert joke")
	}
	return id, nil
}

// finding a joke bu its ID
func (bkr Broker) FindJokeByID(jokeID string) (JokeModel, error) {
	// looking up if the joke (based on its ID) is already in the db

	// when id is empty
	if jokeID = strings.TrimSpace(jokeID); jokeID == "" {
		// errors.New() => custom error
		return JokeModel{}, errors.New("db error on find: provided joke id is invalid, joke id cannot be empty or whitespace")
	}

	filter := bson.M{
		"id": jokeID,
	}
	var foundJoke JokeModel

	if err := bkr.mongo.FindOne("jokesCollection", filter, &foundJoke); err != nil {
		// .Error() turns error into a string type to match the string formatting
		return JokeModel{}, fmt.Errorf("db error on find: failed to find joke by joke id; %s; %w", err.Error(), ErrDocumentDoesntExist)
	}

	return foundJoke, nil
}

// TODO: error i dont understand
// update info for a joke in the db
func (bkr Broker) UpdateJokeByID(jokeID string, newJoke string) error {

	// when id is empty
	if jokeID = strings.TrimSpace(jokeID); jokeID == "" {
		// errors.New() => custom error
		return errors.New("db error on update: provided joke id is invalid, joke id cannot be empty or whitespace")
	}

	// when new joke is empty
	if newJoke = strings.TrimSpace(newJoke); newJoke == "" {
		// errors.New() => custom error
		return errors.New("db error on update: provided joke update is invalid, update cannot be empty or whitespace")
	}

	filter := bson.M{
		"id": jokeID,
	}

	// missing a keyword: $set tells mongo to set this document on the match of the joke id in the filter
	update := bson.M{
		// if using a joke model instead of bson.M, would write over existing joke with the entire new model
		// bson is just changing that one specified field, joke in this case.
		"$set": bson.M{
			"joke": newJoke,
		},
	}

	// update the joke: find it by id, and update it with the updated joke
	if err := bkr.mongo.UpdateOne("jokesCollection", "", filter, update); err != nil {
		return fmt.Errorf("db error on update: failed to update joke by joke id; %s", err.Error())
	}

	return nil

}

// delete a joke by its ID
func (bkr Broker) DeleteJokeByID(jokeID string) error {

	// when id is empty
	if jokeID = strings.TrimSpace(jokeID); jokeID == "" {
		// errors.New() => custom error
		return errors.New("db error on update: provided joke id is invalid, joke id cannot be empty or whitespace")
	}

	filter := bson.M{
		"id": jokeID,
	}

	numDocs, err := bkr.mongo.DeleteOne("jokesCollection", filter)

	if err != nil {
		return fmt.Errorf("db error on delete: failed to delete joke by id; %s", err.Error())
	}

	if numDocs != 1 {
		if numDocs == 0 {
			return errors.New("0 docs were deleted")
		} else if numDocs > 1 {
			return fmt.Errorf("%d documents were deleted, something is terribly wrong", numDocs)
		}
	}

	return nil

}
