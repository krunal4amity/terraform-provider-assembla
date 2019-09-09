package provider

import(
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/krunal4amity/terraform-provider-assembla/assembla"
	"strconv"
)

func resourceWebhookBasic() *schema.Resource{
	return &schema.Resource{
		Schema:map[string]*schema.Schema{
			"space_name":{
				Type:schema.TypeString,
				Required:true,
				Description:"name of the space",
			},
			"title":{
				Type:schema.TypeString,
				Required:true,
				Description:"name of the webhook",
			},
			"auth_type":{
				Type:schema.TypeInt,
				Description:"the type of auth",
				Default:1,
			},
			"http_method":{
				Type:schema.TypeInt,
				Description:"http method- 1 for  GET or 0 for POST",
				Required:true,
			},
			"external_url":{
				Type:schema.TypeString,
				Description:"external url",
				Required:true,
			},
			"content_type":{
				Type:schema.TypeString,
				Description:"content type e.g. application/json etc",
				Optional:true,
			},
			"content":{
				Type:schema.TypeString,
				Description:"POST body",
				Optional:true,
			},
			"post_about_code":{
				Type:schema.TypeString,
				Description:"Post updates about code. Value '1' or '0' in string",
				Optional:true,
			},
			"post_about_code_reviews":{
				Type:schema.TypeString,
				Description:"Post updates about code reviews. Values '1' or '0' in string ",
				Optional:true,
			},
			"post_about_build_tool":{
				Type:schema.TypeString,
				Description:"Post updates about build tool. Values '1' or '0' in string ",
				Optional:true,
			},
			"post_about_file":{
				Type:schema.TypeString,
				Description:"Post updates about file. Values '1' or '0' in string",
				Optional:true,
			},
			"post_about_flow":{
				Type:schema.TypeString,
				Description:"Post updates about flow. Values '1' or '0' in string",
				Optional:true,
			},
			"post_about_ssh_tool":{
				Type:schema.TypeString,
				Description:"Post updates about ssh tool, Values '1' or '0' in string",
				Optional:true,
			},
			"post_about_scrum_reports":{
				Type:schema.TypeString,
				Description:"Post udpates about scrum reports. Values '1' or '0' in string",
				Optional:true,
			},
			"post_about_stream":{
				Type:schema.TypeString,
				Description:"Post updates about stream. Values '1' or '0' in string",
				Optional:true,
			},
			"post_about_team":{
				Type:schema.TypeString,
				Description:"Post updates about team. Values 1 or 0 in string",
				Optional:true,
			},
			"post_about_trac":{
				Type:schema.TypeString,
				Description:"Post updates about trac. Values 1 or 0 in string",
				Optional:true,
			},
			"post_about_wiki":{
				Type:schema.TypeString,
				Description:"Post updates about wiki. Values 1 or 0 in string",
				Optional:true,
			},
		},
		Create:resourceCreateWebhook,
		Update:resourceUpdateWebhook,
		Delete:resourceDeleteWebhook,
	}
}

func resourceCreateWebhook(d *schema.ResourceData,m interface{}) error{
	apiClient := m.(*assembla.Client)
	createWh := assembla.Webhook{
		Webhook :assembla.Hook{
			Title : d.Get("title").(string),
			Authentication_Type : d.Get("auth_type").(int),
			Http_Method: d.Get("http_method").(int),
			External_Url: d.Get("external_url").(string),
			Content : d.Get("content").(string),
			Content_Type : d.Get("content_type").(string),
			Filter : assembla.Filter{
				Code : d.Get("post_about_code").(string),
				Code_Reviews: d.Get("post_about_code_reviews").(string),
				Build_Tool: d.Get("post_about_build_tool").(string),
				File: d.Get("post_about_file").(string),
				Flow: d.Get("post_about_flow").(string),
				Ssh_Tool : d.Get("post_about_ssh_tool").(string),
				Scrum_Reports : d.Get("post_about_scrum_reports").(string),
				Stream:d.Get("post_about_stream").(string),
				Team: d.Get("post_about_team").(string),
				Trac: d.Get("post_about_trac").(string),
				Wiki: d.Get("post_about_wiki").(string),
			},
		},
	}
	wh,_,err := apiClient.Webhooks.CreateWebhook(d.Get("space_name").(string),createWh)
	if err != nil{
		return err
	}
	d.SetId(strconv.Itoa(wh.Id))
	return nil
}

func resourceUpdateWebhook(d *schema.ResourceData,m interface{}) error{
	apiClient := m.(*assembla.Client)
	whId := d.Id()
	createWh := assembla.Webhook{
		Webhook :assembla.Hook{
			Title : d.Get("title").(string),
			Authentication_Type : d.Get("auth_type").(int),
			Http_Method: d.Get("http_method").(int),
			External_Url: d.Get("external_url").(string),
			Content : d.Get("content").(string),
			Content_Type : d.Get("content_type").(string),
			Filter : assembla.Filter {
				Code : d.Get("post_about_code").(string),
				Code_Reviews: d.Get("post_about_code_reviews").(string),
				Build_Tool: d.Get("post_about_build_tool").(string),
				File: d.Get("post_about_file").(string),
				Flow: d.Get("post_about_flow").(string),
				Ssh_Tool : d.Get("post_about_ssh_tool").(string),
				Scrum_Reports : d.Get("post_about_scrum_reports").(string),
				Stream:d.Get("post_about_stream").(string),
				Team: d.Get("post_about_team").(string),
				Trac: d.Get("post_about_trac").(string),
				Wiki: d.Get("post_about_wiki").(string),
			},
		},
	}
	_,_,err := apiClient.Webhooks.UpdateWebhook(d.Get("space_name").(string),whId,createWh)
	if err != nil{
		return err
	}
	return nil
}

func resourceDeleteWebhook(d *schema.ResourceData,m interface{}) error{
	apiClient := m.(*assembla.Client)
	whId := d.Id()
	_,err := apiClient.Webhooks.DeleteWebhook(d.Get("space_name").(string),whId)
	if err != nil{
		return err
	}
	d.SetId("")
	return nil
}
