package models

import "fmt"

type Skill struct {
	Id uint64
	UserId uint64
	Name string
}

func (s Skill) String() string {
	return fmt.Sprintf("Skill #{id: %v, user_id: %v, name: '%v'}", s.Id, s.UserId, s.Name)
}
