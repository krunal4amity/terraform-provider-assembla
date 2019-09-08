package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/krunal4amity/terraform-provider-assembla/provider"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc:provider.Provider,
	})
}
