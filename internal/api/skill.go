package api

import (
	"context"
	"github.com/ozoncp/ocp-skill-api/internal/models"
	"github.com/rs/zerolog/log"
	"github.com/ozoncp/ocp-skill-api/internal/repo"
	desc "github.com/ozoncp/ocp-skill-api/pkg/ocp-skill-api"
)

type SkillAPI struct {
	repo repo.Repo
	desc.UnimplementedOcpSkillApiServer
}

func NewSkillAPI(repo repo.Repo) desc.OcpSkillApiServer {
	return &SkillAPI{repo: repo}
}

func (s *SkillAPI) ListSkillsV1(ctx context.Context, request *desc.ListSkillsRequestV1) (*desc.ListSkillsResponseV1, error) {
	log.Printf("Show skills with params: %v", request)
	entities, err := s.repo.ListEntities(ctx, request.Limit, request.Offset)
	if err != nil {
		return nil, err
	}

	result := make([]*desc.Skill, 0, len(entities))
	for _, entity := range entities {
		skill := &desc.Skill{
			Id:     entity.Id,
			UserId: entity.UserId,
			Name:   entity.Name,
		}

		result = append(result, skill)
	}

	return &desc.ListSkillsResponseV1{Skills: result}, nil
}

func (s *SkillAPI) CreateSkillV1(ctx context.Context, request *desc.CreateSkillRequestV1) (*desc.CreateSkillResponseV1, error) {
	log.Printf("Create skill with params: %v", request)
	skill := models.Skill{
		UserId: request.UserId,
		Name:   request.Name,
	}
	skillId, err := s.repo.AddEntity(ctx, skill)

	return &desc.CreateSkillResponseV1{Id: skillId}, err
}

func (s *SkillAPI) DescribeSkillV1(ctx context.Context, request *desc.DescribeSkillRequestV1) (*desc.DescribeSkillResponseV1, error) {
	log.Printf("Describe skill with params: %v", request)
	skill, err := s.repo.DescribeEntity(ctx, request.Id)

	if err != nil {
		return nil, err
	}

	return &desc.DescribeSkillResponseV1{Skill: &desc.Skill{Id: skill.Id, UserId: skill.UserId, Name: skill.Name}}, nil
}

func (s *SkillAPI) RemoveSkillV1(ctx context.Context, request *desc.RemoveSkillRequestV1) (*desc.RemoveSkillResponseV1, error) {
	log.Printf("Remove skill with params: %v", request)
	result, err := s.repo.RemoveEntity(ctx, request.Id)


	return &desc.RemoveSkillResponseV1{Id: result}, err
}
