package Meeting

import (
	"github.com/painterdrown/go-agenda/entities"
)

type Meetings struct {
	Initiator    string
	Participator string
	Title        string
	STime        string
	ETime        string
}

func AddOneMeeting(futureMeeting Meetings) error {
	err := entities.CreateMeeting(futureMeeting.Initiator, futureMeeting.Title, futureMeeting.Participator, futureMeeting.STime, futureMeeting.ETime)
	return err
}

func AddParticipators(currentUser, Title string, NewParticipator string) error {
	err := entities.AddParticipator(currentUser, Title, NewParticipator)
	return err
}

func DeleteParticipators(currentUser, Title string, NewParticipator string) error {
	err := entities.DeleteParticipator(currentUser, Title, NewParticipator)
	return err
}

func DeleteMeetingByTitle(currentUser, Title string) error {
	err := entities.DeleteMeeting(currentUser, Title)
	return err
}

func QueryMeetingByTime(currentUser string, STime, ETime string) ([]entities.Meeting, error) {
	meetings, err := entities.QueryMeetings(currentUser, STime, ETime)
	return meetings, err
}

func ClearAllUserMeeting(currentUser string) error {
	err := entities.DeleteMeetings(currentUser)
	return err
}
