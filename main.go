package main

import (
    "github.com/hashicorp/terraform/plugin"
    "github.com/AlexisSellier/terraform-provider-croc/croc"
    )

    func main() {
     plugin.Serve(&plugin.ServeOpts{
	ProviderFunc: croc.Provider,
})
}
