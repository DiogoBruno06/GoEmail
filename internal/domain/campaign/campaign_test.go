package compaign

import "testing"

func TestNewCampaign(t *testing.T) {
	name := "Campaign X"
	content := "Body"
	contatcs := []string{"email1@e.com", "email2@e.com"}

	campaign := NewCampaign(name, content, contatcs)

	if campaign.ID != "1" {
		t.Errorf("O id não é 1")
	} else if campaign.Name != name {
		t.Errorf("Expected correct name")
	} else if campaign.Content != content {
		t.Errorf("Expected correct content: ")
	} else if len(campaign.Contacts) != len(contatcs) {
		t.Errorf("Expected correct contacts")
	}
}
