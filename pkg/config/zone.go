package cloudflare

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetZoneIdentifier(domain string) string {
	resp, err := http.Get("https://api.cloudflare.com/client/v4/zones?name=" + domain)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println("Got Zone Information:\n", string(body))
	return string(body)
}
