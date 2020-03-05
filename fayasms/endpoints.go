package fayasms

import (
	"errors"
	"io/ioutil"
	"net/http"
)

var endpoints = map[string]string{
	"send":     "https://devapi.fayasms.com/send",
	"messages": "https://devapi.fayasms.com/messages",
	"balance":  "https://devapi.fayasms.com/balance",
	"estimate": "https://devapi.fayasms.com/estimate",
	"senders":  "https://devapi.fayasms.com/senders",
	"new_id":   "https://devapi.fayasms.com/senders/new",
}

// exec executes the actual http request by fetching the endpoint
// to make the request to
func (f *FayaSMS) exec(endpoint string) (response string, err error) {
	e, ok := endpoints[endpoint]
	if !ok {
		return response, errors.New("fayasms: unknown endpoint targetted")
	}

	res, err := http.PostForm(e, f.payload)
	if err != nil {
		return response, err
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return response, err
	}
	defer res.Body.Close()

	return string(data), nil
}

// Send sends the message to the recipient or recipients you've set
func (f *FayaSMS) Send() (response string, err error) {
	return f.exec("send")
}

// GetEstimate lets you know the number of units it will cost you to send the message.
// This is based on the size or length of your message body and the number of recipients.
func (f *FayaSMS) GetEstimate() (response string, err error) {
	return f.exec("estimate")
}

// GetBalance returns your current balance on FayaSMS
func (f *FayaSMS) GetBalance() (response string, err error) {
	return f.exec("balance")
}
