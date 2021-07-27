package main_test

import (
	"github.com/ozoncp/ocp-skill-api/cmd/ocp-skill-api"
	"reflect"
	"testing"
)


func TestReadConfig(t *testing.T) {
	tests := []struct{
		name string
		input []string
		output []string
		expectedError bool
	}{
		{
			"Empty config list",
			[]string{},
			nil,
			true,
		},
		{
			"Correct work",
			[]string{"./assets/tests/config.txt", "./assets/tests/config.yaml"},
			[]string{"host: 8.8.8.8", "host: 9.9.9.9"},
			false,
		},
		{
			"When all fail absent",
			[]string{"./assets/tests/config.cfg"},
			[]string{},
			false,
		},
		{
			"When one file absent",
			[]string{"./assets/tests/config.cfg", "./assets/tests/config.yaml"},
			[]string{"host: 9.9.9.9"},
			false,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			output, error := main.ReadConfig(testCase.input)
			if (error != nil) != testCase.expectedError  {
				t.Errorf("Error was expected but not received")
			}
			if !reflect.DeepEqual(output, testCase.output) {
				t.Errorf("Expected output: %v got output: %v", testCase.output, output)
			}
		})
	}
}
