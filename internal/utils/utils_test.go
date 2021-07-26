package utils_test

import (
	"github.com/ozoncp/ocp-skill-api/internal/models"
	"github.com/ozoncp/ocp-skill-api/internal/utils"
	"reflect"
	"testing"
)

func TestFilterSlice(t *testing.T) {
	tests := []struct{
		name string
		input []int
		output []int
	}{
		{
			"empty slice",
			[]int{},
			[]int{},
		},
		{
			"no matches",
			[]int{22, 44, 66},
			[]int{22, 44, 66},
		},
		{
			"multiple matches",
			[]int{2, 2, 3, 33, 4, 6, 8, 8, 4, 3},
			[]int{3, 33, 6, 3},
		},
		{
			"all matched",
			[]int{2, 2, 4, 8},
			[]int{},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			output := utils.FilterSlice(testCase.input)
			if !reflect.DeepEqual(output, testCase.output) {
				t.Errorf("expected: %v, got: %v", testCase.output, output)
			}
		})
	}
}

func TestInvertMap(t *testing.T) {
	tests := []struct{
		name string
		input map[string]string
		output map[string]string
	} {
		{
			"Empty map",
			map[string]string{},
			map[string]string{},
		},
		{
			"Not empty map",
			map[string]string{"a":"z", "b":"x"},
			map[string]string{"z":"a", "x":"b"},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			output := utils.InvertMap(testCase.input)
			if !reflect.DeepEqual(output, testCase.output) {
				t.Errorf("expected: %v, got: %v", testCase.output, output)
			}
		})
	}
}

func TestSkillsToBatches(t *testing.T) {
	type args struct {
		input []models.Skill
		size int
	}

	tests := []struct{
		name string
		args args
		output [][]models.Skill
		expectedError bool
	}{
		{
			"zero batches",
			args{[]models.Skill{{1, 1, "BBB"}, {2, 2, "AAA"}}, 0},
			nil,
			true,
		},
		{
			"too many batches",
			args{[]models.Skill{{1, 1, "BBB"}, {2, 2, "AAA"}}, 3},
			nil,
			true,
		},
		{
			"chank size less than 0",
			args{[]models.Skill{{1, 1, "BBB"}, {2, 2, "AAA"}}, -3},
			nil,
			true,
		},
		{
			"chank size 1",
			args{[]models.Skill{{1, 1, "BBB"}, {2, 2, "AAA"}}, 1},
			[][]models.Skill{{{1, 1, "BBB"}}, {{2, 2, "AAA"}}},
			false,
		},
		{
			"chank size 3",
			args{[]models.Skill{{1, 1, "BBB"}, {2, 2, "AAA"},  {3, 3, "XXX"}},
				3},

			[][]models.Skill{{{1, 1, "BBB"}, {2, 2, "AAA"}, {3, 3, "XXX"}}},
			false,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			output, error := utils.SkillsToBatches(testCase.args.input, testCase.args.size)
			if (error != nil) != testCase.expectedError  {
				t.Errorf("Error was expected but not received")
			}
			if !reflect.DeepEqual(output, testCase.output) {
				t.Errorf("Expected output: %v got output: %v", testCase.output, output)
			}
		})
	}
}

func TestSliceToBatches(t *testing.T) {
	type args struct {
		input []string
		size int
	}

	tests := []struct{
		name string
		args args
		output [][]string
		expectedError bool
	}{
		{
			"zero batches",
			args{
				[]string{"one", "two", "three"},
				0,
			},
			nil,
			true,
		},
		{
			"too many batches",
			args{
				[]string{"one", "two"},
				4,
			},
			nil,
			true,
		},
		{
			"chank size less than 0",
			args{
				[]string{"one", "two"},
				-1,
			},
			nil,
			true,
		},
		{
			"chank size 1",
			args{
				[]string{"one", "two", "three", "four", "five",
								"six", "seven", "eight", "nine", "ten"},
				1,
			},
			[][]string{{"one"},{"two"}, {"three"}, {"four"}, {"five"}, {"six"}, {"seven"}, {"eight"}, {"nine"}, {"ten"}},
			false,
		},
		{
			"chank size 3",
			args{
				[]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"},
				3,
			},
			[][]string{{"one", "two", "three"}, {"four", "five", "six"}, {"seven", "eight", "nine"}, {"ten"}},
			false,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			output, error := utils.SliceToBatches(testCase.args.input, testCase.args.size)
			if (error != nil) != testCase.expectedError  {
				t.Errorf("Error was expected but not received")
			}
			if !reflect.DeepEqual(output, testCase.output) {
				t.Errorf("Expected output: %v got output: %v", testCase.output, output)
			}
		})
	}
}

func TestSkillsToMap(t *testing.T) {
	tests := []struct{
		name string
		input []models.Skill
		output map[uint64]models.Skill
		expectedError bool
	}{
		{
			"empty input",
			[]models.Skill{},
			nil,
			true,
		},
		{
			"correct work",
			[]models.Skill{{1, 2, "initial"}, {2, 2, "Best developer"}},
			map[uint64]models.Skill{1:{1, 2, "initial"}, 2:{2, 2, "Best developer"}},
			false,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			output, error := utils.SkillsToMap(testCase.input)
			if (error != nil) != testCase.expectedError  {
				t.Errorf("Error was expected but not received")
			}
			if !reflect.DeepEqual(output, testCase.output) {
				t.Errorf("Expected output: %v got output: %v", testCase.output, output)
			}
		})
	}
}
