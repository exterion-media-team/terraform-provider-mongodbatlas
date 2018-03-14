package main

import (
	"github.com/hashicorp/terraform/plugin"
	"terraform-provider-mongodbatlas/mongodbatlas"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: mongodbatlas.Provider})
}
