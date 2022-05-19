package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	pfb "github.com/jomy10/polls/firebase"
	"github.com/jomy10/polls/util"
)

// Get poll info
func PollInfoHandler(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/api/info" {
		http.Error(res, "Wrong endpoint", http.StatusTeapot)
		return
	}

	if req.Method != "GET" {
		http.Error(res, "Method is not supported", http.StatusMethodNotAllowed)
		return
	}

	res.Header().Add("Access-Control-Allow-Origin", util.ALLOW_ORIGIN)
	res.Header().Add("Access-Control-Allow-Methods", "GET")
	res.Header().Add("Access-Control-Allow-Headers", "*")

	// Firebase
	_, db, err := pfb.InitFirebase("jomy-database")
	if err != nil {
		http.Error(res, "Couldn't connct to firebase", http.StatusBadGateway)
		log.Printf("Error connecting to firebase %d\n", err)
		return
	}

	// Get Parameters
	params, err := getInfoParams(req)
	if err != nil {
		http.Error(res, "No or invalid parameters", http.StatusBadRequest)
		log.Println("No or invalid parameters", err)
		return
	}
	pollId := params.PollId

	// Poll data
	pollData, err := util.GetPoll(db, pollId)
	if err != nil {
		http.Error(res, "Couldn't read firestre document", http.StatusBadGateway)
		log.Println("Couldln't read firestore document", err)
		return
	}

	json, err := json.Marshal(pollData)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	// json.NewEncoder(res).Encode(pollData)
	res.Write(json)
}

type InfoParams struct {
	PollId string `json:"pollId"`
}

func getInfoParams(req *http.Request) (InfoParams, error) {
	pollId := req.URL.Query().Get("pollId")
	var err error
	if pollId == "" {
		err = errors.New("No poll id")
	}

	pollInfo := InfoParams{PollId: pollId}

	return pollInfo, err
}
