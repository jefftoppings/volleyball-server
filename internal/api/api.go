package api

import (
	"encoding/json"
	"github.com/jtoppings/volleyball-server/internal/common"
	"github.com/jtoppings/volleyball-server/internal/responsibilities"
	"net/http"
	"strconv"
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

	questionsToReturn, err := responsibilities.ListQuestions(numberOfQuestions)
	if err != nil {
        errDescription := "Error listing questions: " + err.Error() 
		http.Error(w, errDescription, http.StatusBadRequest)
	}

	response := &ListResponsibilitiesQuestionsResponse{
		Questions: questionsToReturn,
	}

	json.NewEncoder(w).Encode(response)
}
