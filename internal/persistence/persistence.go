package persistence

import "github.com/codetaming/skillsmapper/internal/model"

type DataStore interface {
	SkillPersister
	SkillGetter
}

type SkillPersister interface {
	PersistSkill(skill model.Skill) (err error)
}

type SkillGetter interface {
	GetSkill(id string) (model.Skill, error)
}
