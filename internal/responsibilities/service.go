package responsibilities

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jtoppings/volleyball-server/internal/common"
)

type ListQuestionsConfig struct {
	NumberOfQuestions int
	AllQuestions      bool
}

func readResponsibilitiesFromFile() (map[string][]string, error) {
	// Open the JSON file for reading
	file, err := os.Open("../../assets/responsibilities.json")
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

func ListQuestions(config *ListQuestionsConfig) ([]*common.Question, error) {
	responsibilities, err := readResponsibilitiesFromFile();
	if err != nil {
		return nil, err
	}
	fmt.Printf("%+v", responsibilities)
	return []*common.Question{}, nil
}
