package testing

import (
	"errors"
	"testing"

	"github.com/luitel777/purkheli/utils"
)

type testCases struct {
	input          string
	expectedOutput error
}

func TestValidate(t *testing.T) {
	testcases := []testCases{
		{"hello world this is my first post on Purkheli platform", nil},
		{"", errors.New(utils.TITLE_ERROR)},
		{"      ", errors.New(utils.WHITESPACE_ERROR)},
		{"   VALID TITLE   ", nil},
		{"ZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZ", errors.New(utils.MAXLENGTH_ERROR)},
	}

	for _, tc := range testcases {
		t.Run(tc.input, func(t *testing.T) {
			result := utils.ValidateTitle(tc.input)
			if result == nil && tc.expectedOutput == nil {
			} else if errors.Is(result, tc.expectedOutput) {
				t.Errorf("Expected %s got %s", tc.expectedOutput, result)
			}
		})
	}

}
