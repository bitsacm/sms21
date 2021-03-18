package util

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func GetTokenFromCode(code string) string {
	url := "https://oauth2.googleapis.com/token"

	redirect_uri := "http://localhost:8080/auth/google/redirect"
	client_secret := "uj-pqT0g_HNkFkJjy_3_skza"
	client_id := "214753833156-tufte7ehqjvesud51t5ta3uo59bo6ol4.apps.googleusercontent.com"

	var jsonStr = []byte(`{client_id: "` + /*os.Getenv("GOOGLE_CLIENT_ID")*/ client_id +
		`",client_secret: "` + /*os.Getenv("GOOGLE_CLIENT_SECRET")*/ client_secret +
		`",redirect_uri: "` + /*os.Getenv("GOOGLE_REDIRECT_URI")*/ redirect_uri +
		`",grant_type: 'authorization_code',code: "` + code +
		`"}`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	access_token := struct {
		A string `json:"access_token"`
	}{}

	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal([]byte(body), &access_token)

	if err != nil {
		panic(err)
	}

	return access_token.A
}
