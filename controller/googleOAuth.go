package controller

import (
	"net/http"
	"github.com/labstack/echo" 
	"../constants"
	"strings"
	"net/url"
	"io/ioutil"
	"bytes"
	"encoding/json"
	"github.com/jmoiron/jsonq"
)

var GOOGLE_ID string = constants.ParseConstants()["GOOGLE_ID"]
var GOOGLE_SECRET string = constants.ParseConstants()["GOOGLE_SECRET"]
var GOOGLE_REDIRECT_URI string  = constants.ParseConstants()["GOOGLE_REDIRECT_URI"]
var GOOGLE_SCOPE string  = constants.ParseConstants()["GOOGLE_SCOPE"]

//STEP - 1 : REDIRECT TO GOOGLE'S OAUTH ENDPOINT
func RedirectToGoogleOAuth(context echo.Context) error{
	return context.Redirect(301, "https://accounts.google.com/o/oauth2/auth?client_id="+GOOGLE_ID+"&redirect_uri="+GOOGLE_REDIRECT_URI+"&scope="+GOOGLE_SCOPE+"&response_type=code")
}

//STEP - 2 : READ THE CALLBACK FROM GOOGLE ENDPOINT
func ReadGoogleOAuthData(context echo.Context) error{
	var code string = context.QueryParam("code")
	return context.String(http.StatusOK, RetriveUserData(RetrieveAccessToken(code)))
}

//STEP - 3 : GET ACCESS TOKEN IN RETURN OF AUTHORIZATION CODE
func RetrieveAccessToken(code string) string{
	request_url := "https://www.googleapis.com/oauth2/v4/token"
	form := url.Values{
		"code": {code},
		"client_id":  {GOOGLE_ID},
		"client_secret":  {GOOGLE_SECRET},
		"redirect_uri":     {GOOGLE_REDIRECT_URI},
		"grant_type":     {"authorization_code"},
	}
	body := bytes.NewBufferString(form.Encode())
	rsp, err := http.Post(request_url, "application/x-www-form-urlencoded", body)
	if err != nil {
		panic(err)
	}
	defer rsp.Body.Close()
	body_byte, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		panic(err)
	}
	data := map[string]interface{}{}
	dec := json.NewDecoder(strings.NewReader(string(body_byte)))
	dec.Decode(&data)
	jq := jsonq.NewQuery(data)
	access_token,_ := jq.String("access_token")
	return access_token
}

//STEP 4 : READ USER DATA USING ACCESS TOKEN
func RetriveUserData(access_token string) string{
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://www.googleapis.com/oauth2/v2/userinfo", nil)
	req.Header.Set("Authorization", "Bearer " + access_token)	
	res, _ := client.Do(req)
	body_byte, _ := ioutil.ReadAll(res.Body)
	return string(body_byte)
}