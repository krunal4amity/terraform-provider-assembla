package provider

import(
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/krunal4amity/terraform-provider-assembla/assembla"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"apiKey":{
				Type:schema.TypeString,
				Required:true,
				DefaultFunc:schema.EnvDefaultFunc("ASSEMBLA_KEY",""),
			},
			"apiSecret":{
				Type:schema.TypeString,
				Required:true,
				DefaultFunc:schema.EnvDefaultFunc("ASSEMBLA_SECRET_KEY",""),
			},
		},
		ResourcesMap:map[string]*schema.Resource{
			"webhookbasic":resourceWebhookBasic(),
			"webhookoauth1a":resourceWebhookOauth(),
		},
		ConfigureFunc:providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{},error){
	apiKey := d.Get("apiKey").(string)
	apiSecret := d.Get("apiSecret").(string)
	return assembla.NewClient(nil,apiKey,apiSecret),nil
}
