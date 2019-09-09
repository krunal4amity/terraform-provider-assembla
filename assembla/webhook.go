package assembla

import(
	"fmt"
	"net/http"
)

type WebhookService struct {
	client *Client
}


type Filter struct {
	Code string `json:"code"`
	Code_Reviews string `json:"code_reviews"`
	Build_Tool string `json:"build_tool"`
	Flow string `json:"flow"`
	File string `json:"file"`
	Ssh_Tool string `json:"ssh_tool"`
	Scrum_Reports string `json:"scrum_reports"`
	Stream string `json:"stream"`
	Team string `json:"team"`
	Trac string `json:"trac"`
	Wiki string `json:"wiki"`
}

type Hook struct {
	Id int `json:"id"`
	Title string `json:"title" validate:"required"`
	Enabled bool `json:"enabled"`
	Authentication_Type int `json:"authentication_type" validate:"required"`
	App_Api_Key string `json:"app_api_key"`
	App_Secret string `json:"app_secret"`
	App_Request_Token_Url string `json:"app_request_token_url"`
	App_Access_Token_Url string `json:"app_access_token_url"`
	App_Authorize_Url string `json:"app_authorize_url"`
	App_Authorize_Query string `json:"app_authorize_query"`
	Access_Token string `json:"access_token"`
	Access_Token_Secret string `json:"access_token_secret"`
	External_Url string `json:"external_url"`
	Http_Method int `json:"http_method" validate:"required"`
	Content_Type string `json:"content_type"`
	Content string `json:"content"`
	Filter Filter `json:"filter"`
	Created_At string `json:"created_at"`
	Updated_At string `json:"updated_at"`
}

type Webhook struct{
	Webhook Hook `json:"webhook"`
}

//POST /v1/spaces/:space_id/webhooks
//Create a Webhook
func (s *WebhookService) CreateWebhook(spaceId string,  wb Webhook) (Hook,*http.Response, error){
	//if err := validate.Struct(wb);err != nil{
	//	return nil,nil,err
	//}
	path := fmt.Sprintf("spaces/%s/webhooks",spaceId)
	req,err := s.client.NewRequest(http.MethodPost,path,&wb)
	if err != nil{
		return Hook{},nil,err
	}

	var whRes Hook
	resp,err := s.client.Do(req,&whRes)
	if err != nil{
		return Hook{},resp,err
	}
	return whRes,resp,nil
}


//PUT /v1/spaces/:space_id/webhooks/:webhook_id
//Update a webhook
func (s *WebhookService) UpdateWebhook(spaceId, webhookId string, wb Webhook) (Hook,*http.Response,error){
	path := fmt.Sprintf("spaces/%s/webhooks/%s",spaceId,webhookId)
	req,err := s.client.NewRequest(http.MethodPut,path,&wb)
	if err != nil{
		return Hook{},nil,err
	}

	var whRes Hook
	resp,err := s.client.Do(req,&whRes)
	if err != nil{
		return Hook{},resp,err
	}
	return whRes,resp,nil
}


//DELETE /v1/spaces/:space_id/webhooks/:webhook_id
//Delete a Webhook by ID. This request is only available to users with All permissions in space.
func (s *WebhookService) DeleteWebhook(spaceId, webhookId string) (*http.Response,error){
	path := fmt.Sprintf("spaces/%s/webhooks/%s",spaceId,webhookId)
	req,err := s.client.NewRequest(http.MethodDelete,path,nil)
	if err != nil{
		return nil,err
	}

	resp,err := s.client.Do(req,nil)
	if err != nil{
		return nil,err
	}
	return resp,nil
}