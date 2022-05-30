package core

import (
	"net/http"
	"os"
)

func (receiver *HTTP)SendSMS(msg interface{}) (resp *http.Response, err error) {

	headers := make(map[string]interface{})
	url := os.Getenv("SEND_SMS_URL")
	headers["x-api-key"] = os.Getenv("MSG_API_KEY")
	resp, err = receiver.REQUEST(url, "POST", headers, msg)
	if err != nil { return nil, err}
	return resp, nil
}