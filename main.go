package main

import (
	"github.com/hashicorp/terraform/plugin"

	"github.com/poseidon/terraform-provider-matchbox/matchbox"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: matchbox.Provider,
	})
}
