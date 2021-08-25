package flusher

import (
	"context"
	"github.com/ozoncp/ocp-skill-api/internal/models"
	"github.com/ozoncp/ocp-skill-api/internal/repo"
	"github.com/ozoncp/ocp-skill-api/internal/utils"
)

// Flusher - interface for store
type Flusher interface {
	Flush(context context.Context, skills []models.Skill) ([]models.Skill, error)
}

type flusher struct {
	chunkSize int
	repo  repo.Repo
}

// NewFlusher return Flusher with batches
func NewFlusher(chunkSize int, entityRepo repo.Repo) Flusher {
	return &flusher{
		chunkSize: chunkSize,
		repo:      entityRepo,
	}
}

func (f *flusher) Flush(context context.Context,skills []models.Skill) ([]models.Skill, error) {
	batches, error := utils.SkillsToBatches(skills, f.chunkSize)
	counter := 0
	added := make([]models.Skill, 0)

	if error != nil {
		return added, error
	}
	for _, batch := range batches {
		error := f.repo.AddEntities(context, batch)
		if error != nil {
			return skills[counter:], error
		}
		counter += len(batch)
	}

	return []models.Skill{}, nil
}