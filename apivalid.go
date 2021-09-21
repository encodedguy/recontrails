package securitytrails

import(
	"encoding/json"
	"fmt"
	"net/http"
)

type APIInfo struct{
	Success bool `json:"success"`
}

func (s *Client) APIInfo() (*APIInfo, error){
	res, err := http.Get(fmt.Sprintf("%s/ping?apikey=%s", BaseURL, s.apiKey))

	if err != nil{
		return nil, err
	}

	defer res.Body.Close()

	var ret APIInfo

	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil{
		return nil, err
	}

	return &ret, nil
}
