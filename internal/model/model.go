package model

import "time"

type Skill struct {
	SkillID string    `json:"skill_id"`
	Created time.Time `json:"created"`
	Email   string    `json:"email"`
	Tag     string    `json:"tag"`
	Level   string    `json:"level"`
}
