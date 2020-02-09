package model

import "time"

type Skill struct {
	SkillID string
	Created time.Time
	Email   string
	Tag     string
	Level   string
}
