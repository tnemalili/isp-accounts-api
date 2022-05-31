package core

import (
	"fmt"
	"net/http"
	"os"
)

func (receiver *HTTP) SendSMSHandler(msg interface{}) (resp *http.Response, err error) {

	headers := make(map[string]interface{})
	url := os.Getenv("SEND_SMS_URL")
	headers["x-api-key"] = os.Getenv("MSG_API_KEY")
	resp, err = receiver.REQUEST(url, "POST", headers, msg)
	if err != nil { return nil, err}
	return resp, nil
}

func (receiver *HTTP) SendEmailHandler(msg interface{}) (resp *http.Response, err error) {

	headers := make(map[string]interface{})
	url := os.Getenv("SEND_EMAIL_URL")
	headers["x-api-key"] = os.Getenv("MSG_API_KEY")
	resp, err = receiver.REQUEST(url, "POST", headers, msg)
	if err != nil { return nil, err}
	return resp, nil
}

func (receiver *HTTP) SendPushHandler(msg interface{}) (resp *http.Response, err error) {

	headers := make(map[string]interface{})
	url := os.Getenv("SEND_PUSH_URL")
	headers["x-api-key"] = os.Getenv("MSG_API_KEY")
	resp, err = receiver.REQUEST(url, "POST", headers, msg)
	if err != nil { return nil, err}
	return resp, nil
}

func (receiver *HTTP) FetchCustomerHandler(id string) (resp *http.Response, err error) {

	headers := make(map[string]interface{})
	url := fmt.Sprintf("%v/%v",  os.Getenv("CUSTOMERS_URL"), id)
	headers["x-api-key"] = os.Getenv("CUSTOMERS_API_KEY")
	resp, err = receiver.REQUEST(url, "GET", headers, nil)
	if err != nil { return nil, err}
	return resp, nil
}