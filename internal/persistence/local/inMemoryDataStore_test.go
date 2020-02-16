package local_test

import (
	"github.com/codetaming/skillsmapper/internal/model"
	"github.com/codetaming/skillsmapper/internal/persistence/local"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
	"time"
)

var logger = log.New(os.Stdout, "local_test ", log.LstdFlags|log.Lshortfile)

func TestInMemoryStore(t *testing.T) {
	dataStore, err := local.NewInMemoryDataStore(logger)
	assert.Nil(t, err)
	skillID := uuid.Must(uuid.NewUUID()).String()
	skill := model.Skill{
		Email:   "owner@example.com",
		SkillID: skillID,
		Created: time.Now(),
	}
	dataStore.PersistSkill(skill)
	retrievedSkill, err := dataStore.GetSkill(skillID)
	retrievedAllSkills, err := dataStore.GetAllSkills()
	assert.Nil(t, err)
	assert.Equal(t, skillID, retrievedSkill.SkillID)
	assert.Equal(t, skill, retrievedSkill)
	assert.Equal(t, 1, len(retrievedAllSkills))
}

func TestNotInStoreErrors(t *testing.T) {
	dataStore, _ := local.NewInMemoryDataStore(logger)
	_, err := dataStore.GetSkill("invalid")
	assert.NotNil(t, err)
}
