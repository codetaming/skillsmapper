package local

import (
	"github.com/codetaming/skillsmapper/internal/model"
	"github.com/codetaming/skillsmapper/internal/persistence"
	"log"
)

type InMemoryDataStore struct {
	logger *log.Logger
	skills map[string]model.Skill
}

func (ds InMemoryDataStore) GetAllSkills() ([]model.Skill, error) {
	panic("implement me")
}

func (ds InMemoryDataStore) GetSkill(skillID string) (model.Skill, error) {
	return ds.skills[skillID], nil
}

func (ds InMemoryDataStore) PersistSkill(skill model.Skill) (err error) {
	ds.skills[skill.SkillID] = skill
	return nil
}

func NewInMemoryDataStore(logger *log.Logger) (persistence.DataStore, error) {
	return &InMemoryDataStore{
		logger: logger,
		skills: make(map[string]model.Skill),
	}, nil
}
