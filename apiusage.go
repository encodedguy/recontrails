package securitytrails

import(
	"encoding/json"
	"fmt"
	"net/http"
)

type APIUsage struct{
	AllowedMonthlyUsage int `json:"allowed_monthly_usage"`
	CurrentMonthlyUsage int `json:"current_monthly_usage"`
}

func (s *Client) APIUsage() (*APIUsage, error){
	res, err := http.Get(fmt.Sprintf("%s/account/usage?apikey=%s", BaseURL, s.apiKey))

	if err != nil{
		return nil, err
	}

	defer res.Body.Close()

	var ret APIUsage

	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil{
		return nil, err
	}

	return &ret, nil
}
