package wetsuit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

func GetAuthorizeURL(origin string, cid string, state string) (string, error) {
	u, err := url.ParseQuery(fmt.Sprintf("client_id=%s&response_type=code", cid))
	if err != nil {
		return "", err
	}
	u.Set("state", state)
	return fmt.Sprintf("%s/oauth/authorize?%s", origin, u.Encode()), nil
}

func GetToken(origin string, cid string, cs string, state string, code string) (string, error) {
	type TokenJSON struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
	}
	rb := url.Values{
		"client_id":     {cid},
		"client_secret": {cs},
		"code":          {code},
		"state":         {state},
		"grant_type":    {"authorization_code"},
	}

	resp, err := http.PostForm(fmt.Sprintf("%s/oauth/token", origin), rb)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", errors.New(fmt.Sprintf("%s: %s", resp.Status, b))
	}

	tj := &TokenJSON{}
	if err := json.Unmarshal(b, tj); err != nil {
		return "", err
	}

	return tj.AccessToken, nil
}
