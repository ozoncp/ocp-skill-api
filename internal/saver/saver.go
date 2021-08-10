package saver

import (
	"github.com/ozoncp/ocp-skill-api/internal/flusher"
	"github.com/ozoncp/ocp-skill-api/internal/models"
	"time"
)

type Saver interface {
	Save(skill models.Skill)
	Close()
}

type saver struct {
	skills   []models.Skill
	capacity uint
	timeout  time.Duration
	flusher  flusher.Flusher
	closed   chan bool
	skillsChan chan models.Skill
}

func NewSaver(capacity uint,  flusher flusher.Flusher, timeout time.Duration) Saver  {
	skillSaver := saver{
		capacity: capacity,
		flusher: flusher,
		timeout: timeout,
	}
	skillSaver.init()

	return &skillSaver
}

func (s *saver) Save(skill models.Skill)  {
	s.skillsChan <- skill
}

func (s *saver) Close() {
	s.closed <- true
}

func (s *saver) flush() {
	s.flusher.Flush(s.skills)
	s.skills = []models.Skill{{}}
}

func (s *saver) init() {
	go func() {
		ticker := time.NewTicker(s.timeout)
		defer s.flush()

		for {
			select {
			case skill := <- s.skillsChan:
				s.skills = append(s.skills, skill)
				if uint(len(s.skills)) == s.capacity {
					s.flush()
				}
			case <- ticker.C:
				s.flush()
			case <- s.closed:
				close(s.skillsChan)
				close(s.closed)
				return
			}
		}
	}()
}
