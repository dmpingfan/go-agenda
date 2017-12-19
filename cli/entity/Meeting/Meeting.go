package Meeting

import (
	"github.com/painterdrown/go-agenda/entity"
)

type Meetings struct {
	Initiator    string
	Participator string
	Title        string
	STime        string
	ETime        string
}

func AddOneMeeting(futureMeeting Meetings) error {
	err := entity.CreateMeeting(futureMeeting.Initiator, futureMeeting.Title, futureMeeting.Participator, futureMeeting.STime, futureMeeting.ETime)
	return err
}

func AddParticipators(currentUser, Title string, NewParticipator string) error {
	err := entity.AddParticipator(currentUser, Title, NewParticipator)
	return err
}

func DeleteParticipators(currentUser, Title string, NewParticipator string) error {
	err := entity.DeleteParticipator(currentUser, Title, NewParticipator)
	return err
}

func DeleteMeetingByTitle(currentUser, Title string) error {
	err := entity.DeleteMeeting(currentUser, Title)
	return err
}

func QueryMeetingByTime(currentUser string, STime, ETime string) ([]entity.Meeting, error) {
	meetings, err := entity.QueryMeetings(currentUser, STime, ETime)
	return meetings, err
}

func ClearAllUserMeeting(currentUser string) error {
	err := entity.DeleteMeetings(currentUser)
	return err
}
