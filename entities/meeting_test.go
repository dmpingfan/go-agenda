package entities_test

import (
	"testing"

	"github.com/painterdrown/go-agenda/entities"
)

func TestMeeting(t *testing.T) {
	var err error
	entities.Register("A", "A")
	entities.Register("B", "B")
	entities.Register("C", "C")
	entities.Register("D", "D")
	if err = entities.CreateMeeting("A", "test", "B,C,", "2017-12-01/14:00:00", "2017-12-01/16:00:00"); err != nil {
		t.Error("CreateMeeting 错误")
	}
	if err = entities.AddParticipator("A", "test", "D"); err != nil {
		t.Error("AddParticipator 错误")
	}
	if err = entities.DeleteParticipator("A", "test", "D"); err != nil {
		t.Error("DeleteParticipator 错误")
	}
	if _, err = entities.QueryMeetingsByTime("A", "2017-12-01/00:00:00", "2017-12-31/23:59:59"); err != nil {
		t.Error("QueryMeetingsByTime 错误")
	}
	if err = entities.DeleteMeeting("B", "test"); err != nil {
		t.Error("DeleteMeeting 错误")
	}
	if err = entities.DeleteMeetings("A"); err != nil {
		t.Error("DeleteMeetings 错误")
	}
}
