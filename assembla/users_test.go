package assembla

import (
	"github.com/tj/assert"
	"os"
	"testing"
)

var (
	apiKey, apiKeyExists      = os.LookupEnv("ASSEMBLA_KEY")
	apiSecret, apiSecretExits = os.LookupEnv("ASSEMBLA_SECRET_KEY")
)

func Test_User(t *testing.T) {
	if !apiKeyExists || !apiSecretExits {
		t.SkipNow()
	}
	c := NewClient(nil, apiKey, apiSecret)
	_, _, err := c.Users.ListAuthenticatedUser()
	t.Log(err)
	assert.Nil(t, err)
}
