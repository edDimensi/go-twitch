package twitch

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	clientID = os.Getenv("CLIENT_ID")
}

func TestGetChannel(t *testing.T) {
	req := &GetChannelInputType{
		Channel: "Nightblue3",
	}
	session, err := NewSession(DefaultURL, APIV3Header, clientID)
	resp, err := session.GetChannel(req)
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.NotEqual(t, resp.Views, 0)
	}
}

func TestGetChannelTeams(t *testing.T) {
	req := &GetChannelTeamsInputType{
		Channel: "Nightblue3",
	}
	session, err := NewSession(DefaultURL, APIV3Header, clientID)
	resp, err := session.GetChannelTeams(req)
	assert.Nil(t, err)
	if assert.NotNil(t, resp) {
		assert.NotEqual(t, len(resp.Teams), 0)
	}
}
