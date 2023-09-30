package responsibilities

import (
	"github.com/jtoppings/volleyball-server/internal/common"
)

type ListQuestionsConfig struct {
	NumberOfQuestions int
	AllQuestions      bool
}

func ListQuestions(config *ListQuestionsConfig) []*common.Question {
	return nil
}
