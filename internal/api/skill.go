package api

import (
	"context"
	"github.com/rs/zerolog/log"
	desc "github.com/ozoncp/ocp-skill-api/pkg/ocp-skill-api"
)

type SkillAPI struct {
	desc.UnimplementedOcpSkillApiServer
}

func NewSkillAPI() desc.OcpSkillApiServer {
	return &SkillAPI{}
}

func (s *SkillAPI) ListSkillsV1(ctx context.Context, request *desc.ListSkillsRequestV1) (*desc.ListSkillsResponseV1, error) {
	log.Printf("Show skills with params: %v", request)
	return &desc.ListSkillsResponseV1{}, nil
}

func (s *SkillAPI) CreateSkillV1(ctx context.Context, request *desc.CreateSkillRequestV1) (*desc.CreateSkillResponseV1, error) {
	log.Printf("Create skill with params: %v", request)
	return &desc.CreateSkillResponseV1{}, nil
}

func (s *SkillAPI) DescribeSkillV1(ctx context.Context, request *desc.DescribeSkillRequestV1) (*desc.DescribeSkillResponseV1, error) {
	log.Printf("Describe skill with params: %v", request)
	return &desc.DescribeSkillResponseV1{}, nil
}

func (s *SkillAPI) RemoveSkillV1(ctx context.Context, request *desc.RemoveSkillRequestV1) (*desc.RemoveSkillResponseV1, error) {
	log.Printf("Remove skill with params: %v", request)
	return &desc.RemoveSkillResponseV1{}, nil
}
