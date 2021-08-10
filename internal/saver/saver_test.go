package saver_test

import (
	"github.com/golang/mock/gomock"
	"github.com/ozoncp/ocp-skill-api/internal/mocks"
	"github.com/ozoncp/ocp-skill-api/internal/models"
	"github.com/ozoncp/ocp-skill-api/internal/saver"
	"time"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Saver", func() {
	var (
		ctrl         *gomock.Controller
		mockFlusher  *mocks.MockFlusher
		s 			 saver.Saver
		timeout      time.Duration
		skills       []models.Skill
		capacity     uint
		outError     error
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockFlusher = mocks.NewMockFlusher(ctrl)
		skills = []models.Skill{{Id: 1, UserId: 1, Name: "Initial"},
			{Id: 2, UserId: 1, Name: "Basic"},
			{Id: 3, UserId: 1, Name: "Advanced"},
		}
		timeout  = 1
		capacity = 2
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("#Save", func() {
		BeforeEach(func() {
			capacity = 0
			s = saver.NewSaver(capacity, mockFlusher, timeout)
		})

		BeforeEach(func() {
			mockFlusher.EXPECT().Flush(gomock.Any()).Times(10)
		})
		When("when flag closed is false", func() {
			It("should be ok", func() {
				for _, skill := range skills {
					Expect(s.Save(skill)).Should(BeNil())
				}
				s.Close()
				Expect(outError).Should(BeNil())
			})
		})
		When("when flag closed is true", func() {
			JustBeforeEach(func() {
				s.Close()
			})
			It("should be error", func() {
				Expect(outError).Should(BeEquivalentTo(1))
			})
		})
	})
})
