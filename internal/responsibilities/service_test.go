package responsibilities

import (
	"reflect"
	"testing"

	"github.com/jtoppings/volleyball-server/internal/common"
)

func TestListQuestions(t *testing.T) {
	type args struct {
		config *ListQuestionsConfig
	}
	tests := []struct {
		name    string
		args    args
		want    []*common.Question
		wantErr bool
	}{
		{
			name:    "should list questions",
			args:    args{
				config: &ListQuestionsConfig{
					NumberOfQuestions: 25,
					AllQuestions:      false,
				},
			},
			want:    []*common.Question{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ListQuestions(tt.args.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListQuestions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListQuestions() = %v, want %v", got, tt.want)
			}
		})
	}
}
