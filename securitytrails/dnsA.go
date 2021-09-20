package securitytrails

import(
	"fmt"
	"encoding/json"
	"net/http"
)

type DnsA struct{
	Ip		[]int64		`json:"ip"`
	Organizations	[]string	`json:"organizations"`
	IpCount		[]int		`json:"ip_count"`
	LastSeen	[]string	`json:"last_seen"`
	FirstSeen	[]string	`json:"first_seen"`
}

func (s *Client) DnsA(domain string) (*DnsA, error){
	res, err := http.Get(fmt.Sprintf("%s/history/%s/dns/a?apikey=%s", BaseURL, domain, s.apiKey))

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var ret DnsA

	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil{
		return nil, err
	}

	return &ret, nil
}
