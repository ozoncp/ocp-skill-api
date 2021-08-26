package saver

import (
	"context"
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
	flusher  flusher.Flusher
	close   chan bool
	closed  bool
	skillsChan chan models.Skill
	mutex sync.Mutex
	ticker *time.Ticker
}

func NewSaver(capacity uint,  flusher flusher.Flusher, timeout time.Duration) Saver  {
	skillSaver := saver{
		capacity: capacity,
		flusher: flusher,
		ticker: time.NewTicker(timeout),
		skillsChan: make(chan models.Skill),
		close: make(chan bool),
	}

	go func() {
		defer func() {
			skillSaver.ticker.Stop()
			close(skillSaver.skillsChan)
		}()

		for {
			select {
			case skill := <- skillSaver.skillsChan:
				skillSaver.addSkill(skill)
				if uint(len(skillSaver.skills)) > skillSaver.capacity {
					skillSaver.flush()
				}
			case <- skillSaver.ticker.C:
				if len(skillSaver.skills) > 0 {
					skillSaver.flush()
				}
			case <- skillSaver.close:
				skillSaver.closed = true
				return
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
		return errors.New("channel is closed")
	}
	s.skillsChan <- skill
	return nil
}

func (s *saver) Close() error {
	if s.closed {
		return errors.New("channels already closed")
	}
	close(s.close)
	return nil
}

func (s *saver) flush() error{
	context := context.Background()

	s.mutex.Lock()
	defer s.mutex.Unlock()

	if len(s.skills) == 0 {
		return errors.New("nothing to send")
	}

	notAdded, err := s.flusher.Flush(context, s.skills)

	if err != nil {
		for _, skill := range notAdded {
			s.skillsChan <- skill
		}
	}
	s.skills = s.skills[:0]
	return nil
}
