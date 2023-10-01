package responsibilities

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jtoppings/volleyball-server/internal/common"
)

func readResponsibilitiesFromFile() (map[string][]string, error) {
	file, err := os.Open("assets/responsibilities.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read the file content
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	// Define a struct to unmarshal the JSON data into
	var data map[string]struct {
		Responsibilities []string `json:"responsibilities"`
	}

	// Unmarshal the JSON data into the struct
	if err := json.Unmarshal(content, &data); err != nil {
		return nil, err
	}

	// Extract responsibilities for each referee type
	responsibilities := make(map[string][]string)
	for refereeType, refData := range data {
		responsibilities[refereeType] = refData.Responsibilities
	}

	return responsibilities, nil
}

func ListQuestions(numberOfQuestions int) ([]*common.Question, error) {
	responsibilities, err := readResponsibilitiesFromFile();
	if err != nil {
		return nil, err
	}

	var allQuestions []*common.Question
	for refereeType, duties := range responsibilities {
		for _, duty := range duties {
			allQuestions = append(allQuestions, &common.Question{
				Question: duty,
				Answer:   refereeType,
			})
		}
	}

	// Shuffle questions
	common.ShuffleArray(allQuestions)

	if numberOfQuestions != 0 {
		allQuestions = allQuestions[:numberOfQuestions]
	}

	return allQuestions, nil
}
