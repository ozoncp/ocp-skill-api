package utils

import (
	"errors"
)

func SliceToBatches(slice []string, size int) ([][]string, error) {
	if size <= 0 {
		return nil, errors.New("chunk size should be greater than 0")
	}

	chunksCount := len(slice) / size + 1
	output := make([][]string, chunksCount)
	for i := 0; i < len(slice); i++ {
		sliceIndex := i % size
		output[sliceIndex] = append(output[sliceIndex], slice[i])
	}

	return output, nil
}

func InvertMap(input map[string]string) map[string]string {
	output := make(map[string]string, len(input))
	for k, v := range input {
		output[v] = k
	}

	return output
}

func FilterSlice(input []int) []int  {
	skipList := map[int]bool{2: true, 4: true, 8: true}
	output := make([]int, 0)
	sliceSize := len(input)
	for i := 0; i < sliceSize; i++ {
		value := input[i]
		if _, found := skipList[value]; !found {
			output = append(output, value)
		}
	}
	return output
}