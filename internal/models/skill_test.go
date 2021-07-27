package models_test

import (
	"github.com/ozoncp/ocp-skill-api/internal/models"
	"testing"
)

func TestString(t *testing.T) {
	skill := models.Skill{Id: 1, UserId: 1, Name: "First"}
	expectedOutput := "Skill #{id: 1, user_id: 1, name: 'First'}"
	output := skill.String()
	if expectedOutput != output {
		t.Errorf("expected: %v, got: %v", expectedOutput, output)
	}
}
