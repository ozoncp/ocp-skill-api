package flusher_test

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/ozoncp/ocp-skill-api/internal/mocks"
	"github.com/ozoncp/ocp-skill-api/internal/models"
	"github.com/ozoncp/ocp-skill-api/internal/flusher"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Flusher", func() {
	var (
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
		result, outError = f.Flush(skills)
	})
	Context("Save in repository", func() {
		Context("without exception", func() {
			BeforeEach(func() {
				mockRepo.EXPECT().AddEntities(gomock.Any()).Return(nil).Times(2)
			})
			It("", func() {
				Expect(result).ShouldNot(BeNil())
				Expect(outError).Should(BeNil())
			})
		})
		Context("with exception", func() {
			BeforeEach(func() {
				mockRepo.EXPECT().AddEntities([]models.Skill{{Id: 1, UserId: 1, Name: "Initial"},
					{Id: 2, UserId: 1, Name: "Basic"}}).Return(errors.New("error")).Times(1)
			})
			It("with exception", func() {
				Expect(result).Should(BeEquivalentTo([]models.Skill{}))
				Expect(outError).ShouldNot(BeNil())
			})
		})
		Context("when chunkSize less than 1", func() {
			BeforeEach(func() {
				chunkSize = 0
			})
			It("", func() {
				Expect(result).Should(BeEquivalentTo([]models.Skill{}))
				Expect(outError).ShouldNot(BeNil())
			})
		})
	})
})
