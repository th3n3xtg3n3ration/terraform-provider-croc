package croc

import (
	"net/http"
	"net/url"
	"time"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"strings"
	"sort"
	"bytes"
	"log"
	"io/ioutil"
)

const SignatureMethod string = "HmacSHA256"
const SignatureVersion string = "2"
const Version string = "2013-02-01"

type Config struct {
	client *http.Client
	api_url string
	project string
	access_key string
	secret_key string	
}


func (c Config) signRequest(query string) string {
	u, _ := url.Parse(c.api_url)
	var string_to_sign string = "GET\n" + strings.Split(u.Host, ":")[0] + "\n" + u.Path + "\n" + query
	mac := hmac.New(sha256.New, []byte(c.secret_key))
	mac.Write([]byte(string_to_sign))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

func (c Config) sendRequest(params map[string]string) {
	var buffer bytes.Buffer
	var i int = 0
	default_params := map[string]string{
		"AWSAccessKeyId" : c.project + ":" + c.access_key,
		"SignatureMethod" : SignatureMethod,
		"SignatureVersion" : SignatureVersion,
		"Timestamp" : time.Now().UTC().Format(time.RFC3339),
		"Version" : Version,
	}
	for k, v := range params {
		default_params[k] = v
	}
	mk := make([]string, len(default_params))
	for k, _ := range default_params {
		mk[i] = k
		i++
	}
	sort.Strings(mk)
	for _, e := range mk {
		buffer.WriteString(e + "=" + url.QueryEscape(default_params[e]) + "&")
	}
	query := buffer.String()
	hash := c.signRequest(query[0:len(query) - 1])
	query = query + "Signature=" + url.QueryEscape(hash)
	log.Println(query)
	req, err := http.NewRequest("GET", c.api_url + "?" + query, nil)
	
	req.Header.Add("User-Agent", "Boto/2.12.0 Python/2.7.9 Linux/3.16.0-4-amd64")
	resp, err := c.client.Do(req)
	if err != nil {
		log.Println("Error")
		// handle error
	} else {
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		log.Printf("%s",body)
	}

	
}

func newCrocClient(api_url, access_key, secret_key, project string) *Config{
	return &Config{
		client : &http.Client{
			Timeout: time.Second * 10,
		},
		api_url: api_url,
		project: project,
		access_key: access_key,
		secret_key: secret_key,
	}
}
