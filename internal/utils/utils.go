package utils

import (
	"errors"
	"github.com/ozoncp/ocp-skill-api/internal/models"
)

func SliceToBatches(slice []string, size int) ([][]string, error) {
	if size <= 0 {
		return nil, errors.New("chunk size should be greater than 0")
	}

	if size > len(slice) {
		return nil, errors.New("chunk size should be less than input slice")
	}

	chunksCount := len(slice) / size
	if len(slice) % size != 0 {
		chunksCount = chunksCount + 1
	}

	output := make([][]string, 0)
	for i := 0; i < chunksCount; i ++ {
		from := i * size
		to := (i+1) * size
		if to > len(slice) {
			to = len(slice)
		}
		output = append(output, slice[from:to])
	}

	return output, nil
}

func SkillsToBatches(skills []models.Skill, size int) ([][]models.Skill, error) {
	if size <= 0 {
		return nil, errors.New("chunk size should be greater than 0")
	}

	if size > len(skills) {
		return nil, errors.New("chunk size should be less than input slice")
	}

	chunksCount := len(skills) / size
	if len(skills) % size != 0 {
		chunksCount = chunksCount + 1
	}

	output := make([][]models.Skill, 0)
	for i := 0; i < chunksCount; i ++ {
		from := i * size
		to := (i+1) * size
		if to > len(skills) {
			to = len(skills)
		}
		output = append(output, skills[from:to])
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

func SkillsToMap(input []models.Skill)(map[uint64]models.Skill, error){
	if len(input) < 1 {
		return nil, errors.New("empty input")
	}

	output := make(map[uint64]models.Skill, len(input))
	for _, v := range input {
		output[v.Id] = v
	}

	return output, nil
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
