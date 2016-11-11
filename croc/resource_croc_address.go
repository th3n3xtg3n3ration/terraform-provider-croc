package croc

import (
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"net/url"
	"io/ioutil"
	"net/http"
	"bytes"
	"time"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
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
	var buffer bytes.Buffer
	u, _ := url.Parse(config.api_url)

	timestamp := time.Now().UTC().Format(time.RFC3339)
	buffer.WriteString("GET\n") 
	buffer.WriteString(u.Path + "\n")
	buffer.WriteString("AWSAccessKeyId=" + url.QueryEscape(config.project + ":" + config.access_key) + "\n")
	buffer.WriteString("Action=AllocateAddress\n")
	buffer.WriteString("SignatureMethod=HmacSHA256\n")
	buffer.WriteString("SignatureVersion=2\n")
	buffer.WriteString("Timestamp=" + timestamp + "\n")
	buffer.WriteString("Version=2013-02-01\n")
	log.Println(buffer.String())
	mac := hmac.New(sha256.New, []byte(config.secret_key))
	mac.Write([]byte(buffer.String()))
	log.Println(base64.StdEncoding.EncodeToString(mac.Sum(nil)))
	req, err := http.NewRequest("GET", config.api_url + "/?AWSAccessKeyId=" + url.QueryEscape(config.project + ":" + config.access_key) +
		"&Action=AllocateAddress&SignatureMethod=HmacSHA256&SignatureVersion=2&Timestamp=" + timestamp + "&Version=2013-02-01&Signature=" +
		base64.StdEncoding.EncodeToString(mac.Sum(nil)), nil)
	req.Header.Add("User-Agent", "Terraform croc plugin")
	resp, err := config.client.Do(req)
	if err != nil {
		log.Println("Error")
		// handle error
	} else {
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		log.Printf("%v",body)
	}
	log.Println("resourceCrocAddressCreate")
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
