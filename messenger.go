package sccrawler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
)

// Send SMS via Aligo
func Send(data *AligoApiData) AligoResponse {
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
	comment := "key:" + data.Key
	comment += ",user_id:" + data.User_id
	comment += ",sender:" + data.Sender
	comment += ",receiver:" + data.Receiver
	comment += ",msg:" + data.Msg
	comment += ",msgType:" + data.Msg_type
	comment += ",destination:" + data.Destination
	comment += ",rdate:" + data.Rdate
	comment += ",rtime:" + data.Rtime
	comment += ",testmode:" + data.Testmode_yn

	checkError(err, "Aligo failed to post form\n["+comment+"]")

	defer req.Body.Close()

	respBody, err := ioutil.ReadAll(req.Body)
	checkError(err, "Aligo failed to read sending response")

	var resp map[string]interface{}
	result := NewResponseData()

	if err := json.Unmarshal(respBody, &resp); err == nil {
		codeType := reflect.TypeOf(resp["result_code"]).String()
		if codeType == "string" {
			result.Result_code = resp["result_code"].(string)
		} else if codeType == "float64" {
			result.Result_code = strconv.FormatFloat(resp["result_code"].(float64), 'f', -1, 64)
		}
		result.Message = resp["message"].(string)
	}

	return result
}
