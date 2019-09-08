package provider

import(
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/krunal4amity/terraform-provider-assembla/assembla"
)

func resourceWebhookBasic() *schema.Resource{
	return &schema.Resource{
		Schema:map[string]*schema.Schema{
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
	return nil
}

func resourceUpdateWebhook(d *schema.ResourceData,m interface{}) error{

}

func resourceDeleteWebhook(d *schema.ResourceData,m interface{}) error{

}