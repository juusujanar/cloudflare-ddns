package cloudflare

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func UpdateARecord(ip string)  {
	client := &http.Client{}
	url := "https://api.cloudflare.com/client/v4/zones/" + Config.ZoneIdentifier + "/dns_records/" + Config.DNSRecord

	// TODO: Multiple domain support
	data := Request{
		Type:    "A",
		Name:    Config.Domain,
		Content: ip,
		Ttl:     Config.TTL,
		Proxied: Config.Proxied,
	}
	body, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(body))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Request body: " + string(body))

	req.Header.Set("X-Auth-Email", Config.Email)
	req.Header.Set("X-Auth-Key", Config.ApiToken)
	req.Header.Set("Content-Type", "application/json")

	response, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("The calculated length is:", len(string(body)), "for the url:", url)
	fmt.Println("   ", response.StatusCode)
	hdr := response.Header
	for key, value := range hdr {
		fmt.Println("   ", key, ":", value)
	}
	fmt.Println(string(body))
}