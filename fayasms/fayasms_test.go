package fayasms

import (
	"testing"
	"strings"
)

func TestSetBody(t *testing.T) {
	sms := New("appkey", "appsecret", "senderid")
	sms.SetBody("somebody");
	
	msg := sms.payload.Get("Message")

	if msg != "somebody" {
		t.Errorf("Test failed: expected %v but got %v", "somebody", msg)
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
	sms := New("appkey", "appsecret", "senderid")
	recipients := []string{"23326XXXXXXX", "23324XXXXXXX"}
	sms.SetBulkRecipients(recipients)

	expected := strings.Join(recipients, ",")

	to := sms.payload.Get("To")
	rs := sms.payload.Get("Recipients")

	if to != expected {
		t.Errorf("test failed: expected %v but got %v", expected, to)
	}

	if rs != expected {
		t.Errorf("test failed: expected %v but got %v", expected, rs)
	}
}