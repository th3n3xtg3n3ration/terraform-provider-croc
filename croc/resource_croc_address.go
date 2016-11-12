package croc

import (
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func resourceCrocAddress() *schema.Resource {
	return &schema.Resource{
		Create: resourceCrocAddressCreate,
		Read:   resourceCrocAddressRead,
		Update: resourceCrocAddressUpdate,
		Delete: resourceCrocAddressDelete,
		Exists: resourceCrocAddressExists,
		Schema: map[string]*schema.Schema{
			"ipaddress": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"instanceid" : &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
			},
			"privateaddressid" : &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
			},
			"pivateaddress" : &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceCrocAddressExists(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	log.Println("resourceCrocAddressExists")
	return true, nil
}

func resourceCrocAddressCreate(d *schema.ResourceData, meta interface{}) error {
	// http://docs.aws.amazon.com/general/latest/gr/signature-version-2.html
	config := meta.(*Config)
	log.Println("resourceCrocAddressCreate Start")
	param := map[string]string{"Action" : "AllocateAddress"}
	config.sendRequest(param)
	log.Println("resourceCrocAddressCreate Stop")
	return nil
}

func resourceCrocAddressRead(d *schema.ResourceData, meta interface{}) error {
	log.Println("resourceCrocAddressRead")
	return nil
}

func resourceCrocAddressUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Println("resourceCrocAddressUpdate")
	return nil
}

func resourceCrocAddressDelete(d *schema.ResourceData, meta interface{}) error {
	log.Println("resourceCrocAddressDelete")
	return nil
}
