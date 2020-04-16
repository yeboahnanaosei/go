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

func TestSchedule(t *testing.T) {
	tests := []struct {
		date     string
		wantDate string
		time     string
		wantTime string
	}{
		{date: "2020-08-02", wantDate: "2020-08-02", time: "12:00:00", wantTime: "12:00:00"},
		{date: "2020-04-02", wantDate: "2020-04-02", time: "15:30:42", wantTime: "15:30:42"},
	}

	f := New("", "", "")

	for _, ts := range tests {
		f.Schedule(ts.date, ts.time)
		gotDate := f.payload.Get("ScheduleDate")
		gotTime := f.payload.Get("ScheduleTime")

		if !reflect.DeepEqual(ts.wantDate, gotDate) {
			t.Errorf("test failed: expected %v but got %v", ts.wantDate, gotDate)
		}

		if !reflect.DeepEqual(ts.wantTime, gotTime) {
			t.Errorf("test failed: expected %v but got %v", ts.wantTime, gotTime)
		}
	}
}
