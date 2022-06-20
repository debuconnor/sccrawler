package sccrawler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

func Send(data *AligoApiData) AligoResponse{
	const URL = "https://apis.aligo.in/send/"

	payload := url.Values{}
	payload.Set("key", data.Key)
	payload.Set("user_id", data.User_id)
	payload.Set("sender", data.Sender)
	payload.Set("receiver", data.Receiver)
	payload.Set("msg", data.Msg)
	payload.Set("msg_type", data.Msg_type)
	payload.Set("destination", data.Destination)
	payload.Set("rdate", data.Rdate)
	payload.Set("rtime", data.Rtime)
	payload.Set("testmode_yn", data.Testmode_yn)
	
	c := &http.Client{}
	req, err := c.PostForm(URL, payload)
	checkError(err)

	defer req.Body.Close()

	respBody, err := ioutil.ReadAll(req.Body)
	checkError(err)

	var resp map[string]interface{}
	result := NewResponseData()
	
	if err := json.Unmarshal(respBody, &resp); err == nil{
		result.Result_code = resp["result_code"].(string)
		result.Message = resp["message"].(string)
	}

	return result
}