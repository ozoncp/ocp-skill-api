package api

import (
	"context"
	"encoding/binary"
	"github.com/ozoncp/ocp-skill-api/internal/models"
	"github.com/ozoncp/ocp-skill-api/internal/producer"
	"github.com/ozoncp/ocp-skill-api/internal/utils"
	"github.com/rs/zerolog/log"
	"github.com/ozoncp/ocp-skill-api/internal/repo"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	desc "github.com/ozoncp/ocp-skill-api/pkg/ocp-skill-api"
	opentracinglog "github.com/opentracing/opentracing-go/log"
	"time"
)

const batchSize = 2
var (
	createCounter = promauto.NewCounter(prometheus.CounterOpts{Name: "skills_created"})
	updateCounter = promauto.NewCounter(prometheus.CounterOpts{Name: "skills_updated"})
	removeCounter = promauto.NewCounter(prometheus.CounterOpts{Name: "skills_removed"})
)

type SkillAPI struct {
	repo repo.Repo
	kafka producer.Producer
	desc.UnimplementedOcpSkillApiServer
}

func NewSkillAPI(repo repo.Repo, kafka producer.Producer) desc.OcpSkillApiServer {
	return &SkillAPI{repo: repo, kafka: kafka}
}

func (s *SkillAPI) ListSkillsV1(ctx context.Context, request *desc.ListSkillsRequestV1) (*desc.ListSkillsResponseV1, error) {
	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("ListSkillsV1")
	defer span.Finish()

	log.Printf("Show skills with params: %v", request)
	skills, err := s.repo.ListSkills(ctx, request.Limit, request.Offset)
	if err != nil {
		return nil, err
	}

	result := make([]*desc.Skill, 0, len(skills))
	for _, skill := range skills {
		item := &desc.Skill{
			Id:     skill.Id,
			UserId: skill.UserId,
			Name:   skill.Name,
		}

		result = append(result, item)
	}

	return &desc.ListSkillsResponseV1{Skills: result}, nil
}

func (s *SkillAPI) CreateSkillV1(ctx context.Context, request *desc.CreateSkillRequestV1) (*desc.CreateSkillResponseV1, error) {
	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("CreateSkillV1")
	defer span.Finish()

	log.Printf("Create skill with params: %v", request)
	skill := models.Skill{
		UserId: request.UserId,
		Name:   request.Name,
	}
	skillId, err := s.repo.AddSkill(ctx, skill)
	if err != nil {
		return nil, err
	}

	message := producer.Message{
		Type: producer.Create,
		Body: map[string]interface{}{
			"id":        skillId,
			"timestamp": time.Now().Unix(),
		},
	}

	err = s.kafka.Send(message)
	if err != nil {
		return nil, err
	}
	createCounter.Inc()
	return &desc.CreateSkillResponseV1{Id: skillId}, err
}

func (s *SkillAPI) DescribeSkillV1(ctx context.Context, request *desc.DescribeSkillRequestV1) (*desc.DescribeSkillResponseV1, error) {
	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("DescribeSkillV1")
	defer span.Finish()

	log.Printf("Describe skill with params: %v", request)
	skill, err := s.repo.DescribeSkill(ctx, request.Id)

	if err != nil {
		return nil, err
	}

	return &desc.DescribeSkillResponseV1{Skill: &desc.Skill{Id: skill.Id, UserId: skill.UserId, Name: skill.Name}}, nil
}

func (s *SkillAPI) RemoveSkillV1(ctx context.Context, request *desc.RemoveSkillRequestV1) (*desc.RemoveSkillResponseV1, error) {
	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("RemoveSkillV1")
	defer span.Finish()

	log.Printf("Remove skill with params: %v", request)
	result, err := s.repo.RemoveSkill(ctx, request.Id)
	if err != nil {
		return nil, err
	}

	message := producer.Message{
		Type: producer.Remove,
		Body: map[string]interface{}{
			"id":        request.Id,
			"timestamp": time.Now().Unix(),
		},
	}

	err = s.kafka.Send(message)
	if err != nil {
		return nil, err
	}
	removeCounter.Inc()
	return &desc.RemoveSkillResponseV1{Id: result}, err
}

func (s *SkillAPI) UpdateSkillV1(ctx context.Context, request *desc.UpdateSkillRequestV1) (*desc.UpdateSkillResponseV1, error) {
	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("UpdateSkillV1")
	defer span.Finish()

	log.Printf("Update skill with params: %v", request)
	skill := models.Skill{
		Id: request.Id,
		UserId: request.UserId,
		Name:   request.Name,
	}
	skillId, err := s.repo.UpdateSkill(ctx, skill)
	if err != nil {
		return nil, err
	}

	message := producer.Message{
		Type: producer.Update,
		Body: map[string]interface{}{
			"id":        skillId,
			"timestamp": time.Now().Unix(),
		},
	}

	err = s.kafka.Send(message)
	if err != nil {
		return nil, err
	}
	updateCounter.Inc()
	return &desc.UpdateSkillResponseV1{Id: skillId}, err
}

func (s *SkillAPI) MultiCreateSkillsV1(ctx context.Context, request *desc.MultiCreateSkillRequestV1) (*desc.MultiCreateSkillResponseV1, error) {
	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("MultiCreateSkillsV1")
	defer span.Finish()

	log.Printf("Create skills with params: %v", request)
	result := make([]models.Skill, 0, len(request.Skills))
	for _, skill := range request.Skills {
		item := models.Skill{
			Id:     skill.Id,
			UserId: skill.UserId,
			Name:   skill.Name,
		}

		result = append(result, item)
	}

	skillBatches, err := utils.SkillsToBatches(result, batchSize)
	if err != nil {
		return nil, err
	}
	var added uint64 = 0
	for _, chunk := range skillBatches {
		childSpan := tracer.StartSpan("MultiCreateSkillsV1Batch", opentracing.ChildOf(span.Context()))
		childSpan.LogFields(opentracinglog.Int("size", binary.Size(chunk)))
		defer childSpan.Finish()
		err := s.repo.AddSkills(ctx, chunk)
		if err != nil {
			return nil, err
		}
		added = added + uint64(len(chunk))
	}

	return &desc.MultiCreateSkillResponseV1{Added: added}, err
}