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

func (ds InMemoryDataStore) GetAllSkills() (map[string]model.Skill, error) {
	return ds.skills, nil
}

func (ds InMemoryDataStore) GetSkill(id string) (model.Skill, error) {
	if value, exist := ds.skills[id]; exist {
		return value, nil
	} else {
		return value, &persistence.NotFoundError{Message: id}
	}
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
