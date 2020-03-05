package fayasms

import (
	"io/ioutil"
	"net/http"
)

var endPoints = map[string]string{
	"send":     "https://devapi.fayasms.com/send",
	"messages": "https://devapi.fayasms.com/messages",
	"balance":  "https://devapi.fayasms.com/balance",
	"estimate": "https://devapi.fayasms.com/estimate",
	"senders":  "https://devapi.fayasms.com/senders",
	"new_id":   "https://devapi.fayasms.com/senders/new",
}

// Send sends the message you've sent to the recipient
// or recipients you've set
func (f *FayaSMS) Send() (response string, err error) {
	endpoint, _ := endPoints["send"]
	res, err := http.PostForm(endpoint, f.payload)

	if err != nil {
		return response, err
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return response, err
	}
	defer res.Body.Close()

	return string(data), err
}

// GetEstimate lets you know the number of units it
// will cost you to send the message you intend to
// send. This is based on the size or length of your
// message body and the number of recipients.
func (f *FayaSMS) GetEstimate() (response string, err error) {
	endpoint, _ := endPoints["send"]
	res, err := http.PostForm(endpoint, f.payload)

	if err != nil {
		return response, err
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return response, err
	}
	defer res.Body.Close()

	return string(data), err
}
