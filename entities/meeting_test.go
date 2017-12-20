package entities

import (
	"testing"
)

func TestMeeting(t *testing.T) {
	var err error
	Register("A", "A")
	Register("B", "B")
	Register("C", "C")
	Register("D", "D")
	if err = CreateMeeting("A", "test", "B,C,", "2017-12-01/14:00:00", "2017-12-01/16:00:00"); err != nil {
		t.Error("CreateMeeting 错误")
	}
	if err = AddParticipator("A", "test", "D"); err != nil {
		t.Error("AddParticipator 错误")
	}
	if err = DeleteParticipator("A", "test", "D"); err != nil {
		t.Error("DeleteParticipator 错误")
	}
	if _, err = QueryMeetingsByTime("A", "2017-12-01/00:00:00", "2017-12-31/23:59:59"); err != nil {
		t.Error("QueryMeetingsByTime 错误")
	}
	if err = DeleteMeeting("B", "test"); err != nil {
		t.Error("DeleteMeeting 错误")
	}
	if err = DeleteMeetings("A"); err != nil {
		t.Error("DeleteMeetings 错误")
	}
}
