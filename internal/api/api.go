package api

import (
	"encoding/json"
	"net/http"
	"strconv"
    "github.com/jtoppings/volleyball-server/internal/common"
    "github.com/jtoppings/volleyball-server/internal/responsibilities"
)

type ListResponsibilitiesQuestionsResponse struct {
	Questions []*common.Question `json:"questions"`
}

func ListResponsibilitiesQuestionsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Parse the 'numberOfQuestions' query parameter from the URL
	numberOfQuestionsStr := r.URL.Query().Get("numberOfQuestions")
	numberOfQuestions, err := strconv.Atoi(numberOfQuestionsStr)
	if err != nil {
		http.Error(w, "Invalid 'numberOfQuestions' parameter", http.StatusBadRequest)
		return
	}

    allQuestionsStr := r.URL.Query().Get("allQuestions")
    allQuestions, err := strconv.ParseBool(allQuestionsStr)
    if err != nil {
        http.Error(w, "Invalid 'allQuestions' parameter", http.StatusBadRequest)
        return
    }

    if numberOfQuestions == 0 && !allQuestions {
        http.Error(w, "You must specify numberOfQuestions or allQuestions=true", http.StatusBadRequest)
        return
    }

    questionsToReturn := responsibilities.ListQuestions(&responsibilities.ListQuestionsConfig{
    	NumberOfQuestions: numberOfQuestions,
    	AllQuestions:      allQuestions,
    })
	response := &ListResponsibilitiesQuestionsResponse{
		Questions: questionsToReturn,
	}

	json.NewEncoder(w).Encode(response)
}
