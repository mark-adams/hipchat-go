package hipchat

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// NewAccessToken represents a newly created Hipchat OAuth access token
type NewAccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   uint32 `json:"expires_in"`
	GroupID     uint32 `json:"group_id"`
	GroupName   string `json:"group_name"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

// GetAccessToken returns back an access token for a given client ID and client secrets
func GetAccessToken(clientID string, clientSecret string, scopes []string) (token *NewAccessToken, err error) {

	data := url.Values{"grant_type": {"client_credentials"}, "scopes": {strings.Join(scopes, " ")}}
	req, err := http.NewRequest("POST", "https://api.hipchat.com/v2/oauth/token", strings.NewReader(data.Encode()))
	req.SetBasicAuth(clientID, clientSecret)
	req.Header.Set("Content-type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		content, readerr := ioutil.ReadAll(resp.Body)

		if readerr != nil {
			content = []byte("Unknown error")
		}

		return nil, fmt.Errorf("Couldn't retrieve access token: %s", content)
	}

	content, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal(content, &token)

	return token, nil

}
