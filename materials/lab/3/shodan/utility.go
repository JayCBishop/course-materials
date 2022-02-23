package shodan

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *Client) Utility() (*string, error) {
	res, err := http.Get(fmt.Sprintf("%s/tools/myip?key=%s", BaseURL, s.apiKey))

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var ret string
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}

	return &ret, nil
}
