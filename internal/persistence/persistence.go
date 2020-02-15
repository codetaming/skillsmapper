package persistence

import "github.com/codetaming/skillsmapper/internal/model"

type DataStore interface {
	SkillPersister
	SkillGetter
}

type NotFoundError struct {
	Message string
}

func (e *NotFoundError) Error() string { return e.Message }

type SkillPersister interface {
	PersistSkill(skill model.Skill) (err error)
}

type SkillGetter interface {
	GetAllSkills() ([]model.Skill, error)
	GetSkill(skillID string) (model.Skill, error)
}
