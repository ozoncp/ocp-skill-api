package repo

import (
	"github.com/ozoncp/ocp-skill-api/internal/models"
	"github.com/jmoiron/sqlx"
	"github.com/Masterminds/squirrel"
	"context"
)

// Repo - interface for store Skill
type Repo interface {
	AddEntities(context context.Context, skills []models.Skill) error
	AddEntity(context context.Context, skill models.Skill) (uint64, error)
	ListEntities(context context.Context, limit, offset uint64) ([]models.Skill, error)
	DescribeEntity(context context.Context, entityId uint64) (*models.Skill, error)
	RemoveEntity(context context.Context, entityId uint64) (uint64, error)
}

type repo struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) Repo {
	return &repo{
		db: db,
	}
}

func (r *repo) AddEntity(context context.Context, skill models.Skill) (uint64, error) {
	var result uint64

	query := squirrel.
		Insert("skills").
		Columns("user_id", "name").
		RunWith(r.db).
		Suffix("RETURNING \"id\"").
		PlaceholderFormat(squirrel.Dollar).
		Values(skill.UserId, skill.Name)

	rows, err := query.QueryContext(context)
	if err != nil {
		return result, err
	}
	rows.Next()
	err = rows.Scan(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (r *repo) AddEntities(context context.Context, skills []models.Skill) error {
	query := squirrel.
		Insert("skills").
		Columns("user_id", "name").
		RunWith(r.db).
		PlaceholderFormat(squirrel.Dollar)

	for _, skill := range skills {
		query = query.Values(skill.UserId, skill.Name)
	}

	_, err := query.ExecContext(context)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) ListEntities(context context.Context, limit, offset uint64) ([]models.Skill, error) {
	//var skills []models.Skill
	query := squirrel.Select("id", "user_id", "name").
		From("skills").
		RunWith(r.db).
		Limit(limit).
		Offset(offset).
		PlaceholderFormat(squirrel.Dollar)
	result, err := query.QueryContext(context)
	if err != nil {
		return nil, err
	}

	skills := make([]models.Skill, 0, limit)

	for result.Next() {
		var skill models.Skill
		err := result.Scan(&skill.Id, &skill.UserId, &skill.Name)
		if err != nil {
			return nil, err
		}
		skills = append(skills, skill)
	}

	return skills, nil
}

func (r *repo) DescribeEntity(context context.Context, entityId uint64) (*models.Skill, error){
	var skill models.Skill
	query := squirrel.Select("id", "user_id", "name").
			From("skills").
		    Where(squirrel.Eq{"id": entityId}).
			RunWith(r.db).
			PlaceholderFormat(squirrel.Dollar)

	result := query.QueryRowContext(context)

	err := result.Scan(&skill.Id, &skill.UserId, &skill.Name)
	if err != nil {
		return nil, err
	}

	return &skill, nil
}

func (r *repo) RemoveEntity(context context.Context, entityId uint64) (uint64, error) {

	query := squirrel.Delete("skills").
		Where(squirrel.Eq{"id": entityId}).
		RunWith(r.db).
		PlaceholderFormat(squirrel.Dollar)
	_, err := query.ExecContext(context)

	if err != nil {
		return entityId, err
	}

	return entityId, nil
}