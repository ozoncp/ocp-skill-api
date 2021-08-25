package flusher_test

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/ozoncp/ocp-skill-api/internal/flusher"
	"github.com/ozoncp/ocp-skill-api/internal/mocks"
	"github.com/ozoncp/ocp-skill-api/internal/models"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Flusher", func() {
	var (
		ctx 		 context.Context
		ctrl         *gomock.Controller
		mockRepo     *mocks.MockRepo
		skills 		 []models.Skill
		result       []models.Skill
		f            flusher.Flusher
		chunkSize    int
		outError     error
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockRepo(ctrl)
		ctx = context.Background()

		skills = []models.Skill{{Id: 1, UserId: 1, Name: "Initial"},
			{Id: 2, UserId: 1, Name: "Basic"},
			{Id: 3, UserId: 1, Name: "Advanced"},
		}

		chunkSize = 2
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	JustBeforeEach(func() {
		f = flusher.NewFlusher(chunkSize, mockRepo)
		result, outError = f.Flush(ctx, skills)
	})
	Context("Save in repository", func() {
		Context("without exception", func() {
			BeforeEach(func() {
				mockRepo.EXPECT().AddEntities(ctx, gomock.Any()).Return(nil).Times(2)
			})
			It("should be ok", func() {
				Expect(result).Should(BeEquivalentTo([]models.Skill{}))
				Expect(outError).Should(BeNil())
			})
		})
		Context("with exception", func() {
			BeforeEach(func() {
				mockRepo.EXPECT().AddEntities(ctx, []models.Skill{{Id: 1, UserId: 1, Name: "Initial"},
					{Id: 2, UserId: 1, Name: "Basic"}}).Return(errors.New("error")).Times(1)
			})
			It("with exception", func() {
				Expect(result).Should(BeEquivalentTo(skills))
				Expect(outError).ShouldNot(BeNil())
			})
		})
		Context("when chunkSize less than 1", func() {
			BeforeEach(func() {
				chunkSize = 0
			})
			It("should be not ok", func() {
				Expect(result).Should(BeEquivalentTo([]models.Skill{}))
				Expect(outError).Should(BeEquivalentTo(errors.New("chunk size should be greater than 0")))
			})
		})
	})
})
