package fayasms

import (
	"fmt"
	"net/url"
	"strings"
)

// FayaSMS holds the entire data to be sent
type FayaSMS struct {
	payload url.Values
}

// BodyCharLimit is the limit on the number
// of allowed characters in the SMS body
const BodyCharLimit = 3200

var payload = map[string][]string{
	"AppKey":       {},
	"AppSecret":    {},
	"From":         {},
	"To":           {},
	"Message":      {},
	"ScheduleDate": {},
	"ScheduleTime": {},
	"MessageId":    {},
	"Recipients":   {},
	"Name":         {},
	"Description":  {},
}

// New returns a new FayaSMS instance
func New(appKey, appSecret, senderID string) *FayaSMS {
	f := new(FayaSMS)
	f.payload = url.Values{}
	f.payload.Set("AppKey", appKey)
	f.payload.Set("AppSecret", appSecret)
	f.payload.Set("From", senderID)
	return f
}

// SetBody sets the body of the text message to be sent.
// The body must not be more than 3200 characters. Must
// contain on UTF-8 characters
func (f *FayaSMS) SetBody(body string) error {
	bodyLength := len(body)
	if len(body) > BodyCharLimit {
		return fmt.Errorf("fayasms: sms body cannot be more than %d characters. you currently have %d", BodyCharLimit, bodyLength)
	}

	f.payload.Set("Message", body)
	return nil
}

// SetRecipient sets the recipient of the message.
// It must comply with the telephone rules. Meaning
// the number must be in international telephone format
// (e.g: 23326XXXXXXX,23324XXXXXXX,23320XXXXXXX...)
// If you are sending a message to multiple recipients
// use SetBulkRecipients instead
func (f *FayaSMS) SetRecipient(r string) {
	f.payload.Set("To", r)
	f.payload.Set("Recipients", r)
}

// SetBulkRecipients sets all phone numbers in r
// as recipients of the text message. Each phone
// number in r must comply with international
// telephone rules. Meaning the number must be in
// international format. eg e.g: 23326XXXXXXX,23324XXXXXXX
func (f *FayaSMS) SetBulkRecipients(r []string) {
	recipients := strings.Join(r, ",")
	f.payload.Set("To", recipients)
	f.payload.Set("Recipients", recipients)
}

