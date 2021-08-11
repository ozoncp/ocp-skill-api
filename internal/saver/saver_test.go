package saver_test

import (
	"errors"
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
		skills 		 []models.Skill
		capacity     uint
		timeout      time.Duration
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockFlusher = mocks.NewMockFlusher(ctrl)

		skills = []models.Skill{{Id: 1, UserId: 1, Name: "Initial"},
			{Id: 2, UserId: 1, Name: "Basic"},
			{Id: 3, UserId: 1, Name: "Advanced"},
		}

		capacity = 2
		timeout = 100 * time.Millisecond
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	JustBeforeEach(func() {
		s = saver.NewSaver(capacity, mockFlusher, timeout)
	})
	Context("Save in repository", func() {
		Context("without error", func() {
			BeforeEach(func() {
				mockFlusher.EXPECT().Flush([]models.Skill{{Id: 2, UserId: 1, Name: "Basic"},
					{Id:1, UserId: 1, Name: "Initial"}}).Times(1)
			})
			It("should be ok", func() {
				_ = s.Save(skills[1])
				outError := s.Save(skills[0])
				time.Sleep(time.Millisecond * 200)
				Expect(outError).Should(BeNil())
			})
		})
		Context("with error", func() {
			It("channel is closed", func() {
				s.Close()
				time.Sleep(time.Millisecond * 100)
				outError := s.Save(skills[1])
				Expect(outError).Should(BeEquivalentTo(errors.New("channel is closed")))
			})
		})
	})
	Context("Close saver", func() {
		Context("first time closed", func() {
			It("without error", func() {
				outError := s.Close()
				Expect(outError).Should(BeNil())
			})
		})
		Context("second time closed", func() {
			It("with error", func(){
				_ = s.Close()
				time.Sleep(time.Millisecond * 100)
				outError := s.Close()
				Expect(outError).Should(BeEquivalentTo(errors.New("channels already closed")))
			})
		})
	})
})

