package croc

import (
	"log"
       "github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	)
	
 func Provider()  terraform.ResourceProvider {
  return &schema.Provider{
  Schema: map[string]*schema.Schema{
      "project" : &schema.Schema{
      Type: schema.TypeString,
	  Required: true,
      },
      "storage_url" : &schema.Schema{
      Type: schema.TypeString,
	  Optional: true,
	  Default: "https://storage.cloud.croc.ru:443",
      },
      "api_url" : &schema.Schema{
      Type: schema.TypeString,
	  Optional: true,
	  Default: "https://api.cloud.croc.ru:443/",
      },
      "monitoring_url" : &schema.Schema{
      Type: schema.TypeString,
	  Optional: true,
	  Default: "https://monitoring.cloud.croc.ru:443/",
      },
	"ec2_access_key" : &schema.Schema{
	Type: schema.TypeString,
	    Required: true,
	},
	"ec2_secret_key" : &schema.Schema{
	Type: schema.TypeString,
	    Required: true,
	    },
	"aws_access_key" : &schema.Schema{
	Type: schema.TypeString,
	    Required: true,
	    },
       "aws_secret_key" : &schema.Schema{
	Type: schema.TypeString,
	    Required: true,
	    },
    },
  ResourcesMap: map[string]*schema.Resource{},
      },
    ConfigureFunc: providerConfigure,
}

func providerConfigure(d * schema.ResourceData) (interface{}, error) {
 config := Config{}
  log.Println("[INFO], Initializing Croc client")
  return config.Client()
}
