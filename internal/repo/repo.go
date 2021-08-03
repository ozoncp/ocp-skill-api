package repo

import "github.com/ozoncp/ocp-skill-api/internal/models"

// Repo - interface for store Skill
type Repo interface {
	AddEntities(skills []models.Skill) error
	ListEntities(limit, offset uint64) ([]models.Skill, error)
	DescribeEntity(entityId uint64) (*models.Skill, error)
}
