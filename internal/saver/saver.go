package saver

import (
	"errors"
	"github.com/ozoncp/ocp-skill-api/internal/flusher"
	"github.com/ozoncp/ocp-skill-api/internal/models"
	"sync"
	"time"
)

type Saver interface {
	Save(skill models.Skill) error
	Close() error
}

type saver struct {
	skills   []models.Skill
	capacity uint
	timeout  time.Duration
	flusher  flusher.Flusher
	close   chan bool
	closed  bool
	skillsChan chan models.Skill
	mutex sync.Mutex
}

func NewSaver(capacity uint,  flusher flusher.Flusher, timeout time.Duration) Saver  {
	skillSaver := saver{
		capacity: capacity,
		flusher: flusher,
		timeout: timeout,
	}

	go func() {
		ticker := time.NewTicker(skillSaver.timeout)
		defer skillSaver.flush()

		for {
			select {
			case skill := <- skillSaver.skillsChan:
				skillSaver.addSkill(skill)
				if uint(len(skillSaver.skills)) == skillSaver.capacity {
					_ = skillSaver.flush()
				}
			case <- ticker.C:
				skillSaver.flush()
			case <- skillSaver.close:
				close(skillSaver.skillsChan)
				close(skillSaver.close)
				skillSaver.closed = true
			}
		}
	}()

	return &skillSaver
}

func (s *saver) addSkill(skill models.Skill) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.skills = append(s.skills, skill)
}
func (s *saver) Save(skill models.Skill) error {
	if s.closed {
		return errors.New("channels closed")
	}
	s.skillsChan <- skill
	return nil
}

func (s *saver) Close() error {
	if s.closed {
		return errors.New("channels already closed")
	}
	s.close <- true
	return nil
}

func (s *saver) flush() error{
	s.mutex.Lock()
	defer s.mutex.Unlock()

	notAdded, err := s.flusher.Flush(s.skills)
	if err != nil {
		for _, skill := range notAdded {
			s.skillsChan <- skill
		}
	}
	s.skills = nil
	return nil
}
