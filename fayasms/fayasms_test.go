package fayasms

import (
	"reflect"
	"testing"
)

func TestSetBody(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{input: "somemessage", want: "somemessage"},
		{input: "a very long text message", want: "a very long text message"},
	}

	f := New("", "", "")

	for _, ts := range tests {
		f.SetBody(ts.input)
		got := f.payload.Get("Message")
		if !reflect.DeepEqual(ts.want, got) {
			t.Errorf("test failed: expected %s but got %s", ts.want, got)
		}

		if len(got) > AllowedMsgLen {
			t.Errorf("test failed: message length exceeds limit")
		}
	}
}

func TestSetRecipient(t *testing.T) {
	sms := New("appkey", "appsecret", "senderid")
	recipient := "23326XXXXXXX"
	sms.SetRecipient(recipient)

	to := sms.payload.Get("To")
	rs := sms.payload.Get("Recipients")

	if to != recipient {
		t.Errorf("test failed: expected %v got %v", recipient, to)
	}

	if rs != recipient {
		t.Errorf("test failed: expected %v got %v", recipient, rs)
	}
}

func TestSetBulkRecipients(t *testing.T) {
	tests := []struct {
		input []string
		want  string
	}{
		{input: []string{"23326XXXXXXX", "23324XXXXXXX"}, want: "23326XXXXXXX,23324XXXXXXX"},
		{input: []string{"+233261111111", "+233541111111"}, want: "+233261111111,+233541111111"},
	}

	f := New("", "", "")

	for _, tb := range tests {
		f.SetBulkRecipients(tb.input)
		to := f.payload.Get("To")
		rcs := f.payload.Get("Recipients")

		if !reflect.DeepEqual(tb.want, to) {
			t.Errorf("test failed: expected %v but got %v", tb.want, to)
		}

		if !reflect.DeepEqual(tb.want, rcs) {
			t.Errorf("test failed: expected %v but got %v", tb.want, to)
		}
	}
}
