package assembla

import (
	"github.com/tj/assert"
	"testing"
)


func Test_User(t *testing.T){
	c := NewClient(nil,"06aa938f884f26db76e9","4c9a0ff184ead06bb5837fdac9357a5c88d5a5c8")
	user,resp,err := c.Users.ListAuthenticatedUser()
	t.Log(resp.Body)
	assert.NotNil(t,err)
	t.Log(user)
}