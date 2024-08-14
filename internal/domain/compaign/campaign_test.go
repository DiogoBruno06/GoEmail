package compaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_NewCampaign_CreateCampaign(t *testing.T) {
	assert := assert.New(t)
	name := "Campaign X"
	content := "Body"
	contatcs := []string{"email1@e.com", "email2@e.com"}

	campaign := NewCampaign(name, content, contatcs)

	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contatcs))
}

func Test_NewCampaignID_IDisNotNull(t *testing.T) {
	assert := assert.New(t)
	name := "Campaign X"
	content := "Body"
	contatcs := []string{"email1@e.com", "email2@e.com"}

	campaign := NewCampaign(name, content, contatcs)

	assert.NotNil(campaign.ID)
}

func Test_NewCampaignID_CreatedOnisNotNull(t *testing.T) {
	assert := assert.New(t)
	name := "Campaign X"
	content := "Body"
	contatcs := []string{"email1@e.com", "email2@e.com"}
	now := time.Now().Add(-time.Minute)

	campaign := NewCampaign(name, content, contatcs)

	assert.Greater(campaign.CreatedOn, now)
}
