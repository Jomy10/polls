package util

import (
	"cloud.google.com/go/firestore"
	"context"
)

type Poll struct {
	Title string `json:"title"`
	// Amount of votes mapped to the options
	Votes map[string]int `json:"votes"`
}

// Get a poll from the database
func GetPoll(db *firestore.Client, pollId string) (Poll, error) {
	doc := db.Collection("polls").Doc(pollId)
	dsnap, err := doc.Get(context.Background())

	var pollData Poll
	dsnap.DataTo(&pollData)

	return pollData, err
}
