package assembla

import (
	"testing"
	"github.com/tj/assert"
)

func Test_CreateWebhook(t *testing.T){
	if !apiKeyExists || !apiSecretExits{
		t.SkipNow()
	}

	c := NewClient(nil,apiKey,apiSecret)
	createWb := Webhook{
		Hook{
			Title:"webhook6",
			Authentication_Type:0,
			Http_Method:0,
			External_Url:"https://example2.com/index?param1=value2",
			Filter:Filter{
				Scrum_Reports:"1",
			},
		},
	}
	wh,_,err := c.Webhooks.CreateWebhook("terraformprovider",createWb)
	t.Logf("%+v",wh)
	assert.Nil(t,err)
}


func Test_UpdateWebhook(t *testing.T){
	if !apiKeyExists || !apiSecretExits{
		t.SkipNow()
	}

	c := NewClient(nil,apiKey,apiSecret)
	updateWb:=Webhook{
		Hook{
			Title:"webhook2",
			Authentication_Type:0,
			Http_Method:0,
			External_Url:"https://example3.com/index?param1=value2",
		},
	}
	wh,_,err := c.Webhooks.UpdateWebhook("terraformprovider","webhook2", updateWb)
	t.Log(wh)
	assert.Nil(t,err)
}

func Test_DeleteWebhook(t *testing.T){
	if !apiKeyExists || !apiSecretExits{
		t.SkipNow()
	}

	c := NewClient(nil,apiKey,apiSecret)
	_,err := c.Webhooks.DeleteWebhook("terraformprovider","webhook2")
	assert.Nil(t,err)
}
