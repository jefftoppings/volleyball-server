package api

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Question struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

type ListResponsibilitiesQuestionsResponse struct {
	Questions []*Question `json:"questions"`
}

func ListResponsibilitiesQuestionsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Parse the 'numberOfQuestions' query parameter from the URL
	numberOfQuestionsStr := r.URL.Query().Get("numberOfQuestions")
	numberOfQuestions, err := strconv.Atoi(numberOfQuestionsStr)
	if err != nil || numberOfQuestions == 0 {
		http.Error(w, "Invalid 'numberOfQuestions' parameter", http.StatusBadRequest)
		return
	}

	response := &ListResponsibilitiesQuestionsResponse{
		Questions: []*Question{{
			Question: "Who performs the coin toss with the team captains at the start of the match?",
			Answer:   "First Referee",
		}},
	}

	json.NewEncoder(w).Encode(response)
}
