package fayasms

import (
	"fmt"
	"net/url"
	"strings"
)

// FayaSMS holds the entire data to be sent
type FayaSMS struct {
	payload url.Values
	extra   bool
}

// AllowedMsgLen is the limit on the number of allowed characters in the SMS body
const AllowedMsgLen = 3200

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
// The body must not be more than 3200 characters.
// Must contain only UTF-8 characters
func (f *FayaSMS) SetBody(body string) error {
	msg := strings.TrimSpace(body)
	msgLen := len(msg)
	if msgLen > AllowedMsgLen {
		return fmt.Errorf("fayasms: sms body cannot be more than %d characters. you currently have %d", AllowedMsgLen, msgLen)
	}

	f.payload.Set("Message", body)
	return nil
}

// SetRecipient sets the recipient of the message. It must comply with the
// telephone rules. Meaning the number must be in international telephone format
// e.g: 23326XXXXXXX,23324XXXXXXX,23320XXXXXXX If you are sending a message
// to multiple recipients use SetBulkRecipients instead
func (f *FayaSMS) SetRecipient(r string) {
	f.payload.Set("To", r)
	f.payload.Set("Recipients", r)
}

// SetBulkRecipients sets all phone numbers in r as recipients of the text message.
// Each phone number in r must comply with international telephone rules.
// Meaning the number must be in international format.
// eg e.g: 23326XXXXXXX,23324XXXXXXX
func (f *FayaSMS) SetBulkRecipients(r []string) {
	recipients := strings.Join(r, ",")
	f.payload.Set("To", recipients)
	f.payload.Set("Recipients", recipients)
}

// Schedule schedules a message to be sent later on the set date and time.
// date must be in the format "YYYY-MM-DD" eg. "2020-12-31" which is (2020 December 31st).
// time must be in the 24hr format "HH:ii:ss" eg "13:30:04" which is 1pm 30min 4sec
func (f *FayaSMS) Schedule(date string, time string) {
	f.payload.Set("ScheduleDate", date)
	f.payload.Set("ScheduleTime", time)
	f.extra = true
}
