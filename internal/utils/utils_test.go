package utils

import (
	"reflect"
	"testing"
)

func TestFilterSlice(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8}
	expectedOutput := []int{1, 3, 5, 6, 7}

	output := FilterSlice(input)
	if !reflect.DeepEqual(output, expectedOutput) {
		t.Errorf("expected: %v, got: %v", expectedOutput, output)
	}
}

func TestInvertMap(t *testing.T) {
	input := map[string]string{"a":"z", "b":"x"}
	expectedOutput := map[string]string{"z":"a", "x":"b"}

	output := InvertMap(input)
	if !reflect.DeepEqual(output, expectedOutput) {
		t.Errorf("expected: %v, got: %v", expectedOutput, output)
	}
}

func TestSliceToZeroBatches(t *testing.T) {
	input := []string{"one", "two", "three"}

	output, error := SliceToBatches(input, 0)
	if error == nil {
		t.Errorf("Expected error, received %v", output)
	}
}

func TestSliceToBatches(t *testing.T) {
	input := []string{"one", "two", "three"}
	expectedOutput := [][]string{{"one", "three"}, {"two"}}

	output, error := SliceToBatches(input, 2)
	if error != nil {
		t.Errorf("Got error: %v", output)
	}
	if !reflect.DeepEqual(output, expectedOutput) {
		t.Errorf("expected: %v, got: %v", expectedOutput, output)
	}
}
