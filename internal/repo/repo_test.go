package repo_test

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/ozoncp/ocp-skill-api/internal/models"
	"github.com/ozoncp/ocp-skill-api/internal/repo"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Repo", func() {
	const tableName = "skills"

	var (
		db     *sql.DB
		sqlxDB *sqlx.DB
		dbMock sqlmock.Sqlmock
		ctx    context.Context
		r      repo.Repo
		skills []models.Skill
	)

	BeforeEach(func() {
		var err error
		db, dbMock, err = sqlmock.New()
		Expect(err).Should(BeNil())
		sqlxDB = sqlx.NewDb(db, "sqlmock")
		ctx = context.Background()
		r = repo.NewRepo(sqlxDB)
		skills = []models.Skill{{UserId: 1, Name: "Initial"},
			{UserId: 1, Name: "Basic"},
			{UserId: 1, Name: "Advanced"},
		}
	})

	AfterEach(func() {
		var err error
		dbMock.ExpectClose()
		err = db.Close()
		Expect(err).Should(BeNil())
	})

	Context("AddSkills", func() {
		It("should be success", func() {
			expectedQueryArgs := make([]driver.Value, 0, len(skills)*2)

			for _, skill := range skills {
				expectedQueryArgs = append(expectedQueryArgs, skill.UserId, skill.Name)
			}

			dbMock.ExpectPrepare(
				"INSERT INTO skills (user_id,name) VALUES ($1,$2),($3,$4),($5,$6)",
			).
				ExpectQuery().
				WithArgs(expectedQueryArgs...)

			err := r.AddSkills(ctx, skills)
			Expect(err).Should(BeNil())

		})
	})
})