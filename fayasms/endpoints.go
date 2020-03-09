package fayasms

import (
	"errors"
	"fmt"
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

// mandatoryFields are required by FayaSMS to be present in every request
var mandatoryFields = []string{
	"AppKey", "AppSecret",
}

// checkMandatoryFields checks to ensure that the mandatory fields are set
func (f *FayaSMS) checkMandatoryFields(mandatoryFields []string) error {
	for _, field := range mandatoryFields {
		if f.payload.Get(field) == "" {
			return fmt.Errorf("fayasms: a mandatory field has not been set. please supply all mandatory fields which are: %v", mandatoryFields)
		}
	}

	return nil
}

// contingentFields are only required based on the endpoint being hit.
// This map shows the endpoints and the fields they require
var contingentFields = map[string][]map[string]string{
	"send": {
		{"name": "From", "errMsg": "no sender id has been set"},
		{"name": "Message", "errMsg": "no message body has been set"},
		{"name": "To", "errMsg": "no recipient has been set"},
	},
	"estimate": {
		{"name": "Recipients", "errMsg": "no recipient has been set"},
		{"name": "Message", "errMsg": "no message body set"},
	},
}

var extraContingentFields = map[string][]map[string]string{
	"send": {
		{"name": "ScheduleDate", "errMsg": "no ScheduleDate supplied"},
		{"name": "ScheduleTime", "errMsg": "no ScheduleTime supplied"},
	},
	"messages": {
		{"name": "MessageId", "errMsg": "no message id supplied"},
	},
}

// checkContingentFields checks that all contingent fields required by endpoint are set
func (f *FayaSMS) checkContingentFields(endpoint string, contingentFields map[string][]map[string]string) error {
	fields, ok := contingentFields[endpoint]

	// Some endpoints do not have any contingent fields
	if !ok {
		return nil
	}

	for _, field := range fields {
		if f.payload.Get(field["name"]) == "" {
			return fmt.Errorf("fayasms: %v", field["errMsg"])
		}
	}

	if f.extra {
		f.extra = false
		err := f.checkContingentFields(endpoint, extraContingentFields)
		if err != nil {
			return err
		}
	}

	return nil
}

// exec executes the actual http request by fetching the endpoint
// to make the request to
func (f *FayaSMS) exec(endpoint string) (response string, err error) {
	endpnt, ok := endpoints[endpoint]
	if !ok {
		return response, errors.New("fayasms: unknown endpoint targetted")
	}

	if err = f.checkMandatoryFields(mandatoryFields); err != nil {
		return response, err
	}

	if err = f.checkContingentFields(endpoint, contingentFields); err != nil {
		return response, err
	}

	res, err := http.PostForm(endpnt, f.payload)
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

// RequestSenderID makes a request to FayaSMS for a new sender id
// senderID is the sender id you are requesting for.
// desc is a description of the sender id. What will you use the id for.
// The description is used in the approval process
func (f *FayaSMS) RequestSenderID(senderID, desc string) (response string, err error) {
	f.payload.Set("Name", senderID)
	f.payload.Set("Description", desc)

	return f.exec("new_id")
}


// RetrieveMessages returns all the messages you've sent using your AppKey and AppSecret
func (f *FayaSMS) RetrieveMessages() (response string, err error) {
	return f.exec("messages")
}


// RetrieveMessage retrieves a particular message you've sent whose id is messageID
func (f *FayaSMS) RetrieveMessage(messageID string) (response string, err error) {
	f.payload.Set("MessageId", messageID)
	f.extra = true
	return f.exec("messages")
}
