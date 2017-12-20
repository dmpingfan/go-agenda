package entities

import (
	"errors"
	"strings"
)

// Meeting .
type Meeting struct {
	Title         string
	Sponsor       string
	Participators string
	Stime         string
	Etime         string
}

var createMeetingTable = `
CREATE TABLE IF NOT EXISTS meeting (
	title VARCHAR(32) NOT NULL,
	sponsor VARCHAR(32) NOT NULL,
	participators TEXT NOT NULL,
	stime CHAR(16) NOT NULL,
	etime CHAR(16) NOT NULL
);
`
var createMeeting = `INSERT INTO meeting (title, sponsor, participators, stime, etime) VALUES (?, ?, ?, ?, ?);`
var deleteMeetingByUsername = `DELETE FROM meeting WHERE sponsor = ?;`
var deleteMeetingByTitle = `DELETE FROM meeting WHERE title = ?;`
var queryMeetingByTitle = `SELECT * FROM meeting WHERE meeting.title = ?;`
var queryMeetingsByUsername = `SELECT * FROM meeting AS m WHERE m.sponsor = ? OR m.participators LIKE ?;`
var queryMeetingsByTime = `SELECT * FROM meeting AS m WHERE (m.sponsor = ? OR m.participators LIKE ?) AND (m.stime >= ? AND m.etime <= ?);`
var addParticipator = `UPDATE meeting SET participators = participators || ? WHERE title = ?;`
var updateParticipators = `UPDATE meeting SET participators = ? WHERE title = ?;`

// CreateMeeting .
func CreateMeeting(username, title, participators, stime, etime string) error {
	var err error
	// 检查会议的标题是否已存在
	meeting := Meeting{}
	row := db.QueryRow(queryMeetingByTitle, title)
	row.Scan(&meeting.Title, &meeting.Sponsor, &meeting.Participators, &meeting.Stime, &meeting.Etime)
	if meeting.Title != "" {
		return errors.New("会议标题已存在")
	}
	// 创建会议
	_, err = db.Exec(createMeeting, title, username, participators, stime, etime)
	checkErr(err)
	return err
}

// AddParticipator .
func AddParticipator(username, title, participator string) error {
	var err error

	// TODO: 判断 username 是否是举办者
	// TODO: 判断 participant 是否已经是参与者

	_, err = db.Exec(addParticipator, participator+",", title)
	checkErr(err)
	return err
}

// DeleteParticipator .
func DeleteParticipator(username, title, participator string) error {
	var err error
	meeting := Meeting{}
	row := db.QueryRow(queryMeetingByTitle, title)
	row.Scan(&meeting.Title, &meeting.Sponsor, &meeting.Participators, &meeting.Stime, &meeting.Etime)
	if meeting.Title == "" {
		return errors.New("该会议不存在")
	}
	if meeting.Sponsor != username {
		return errors.New("只有该会议的举办者才能删除参与者")
	}
	pos := strings.Index(meeting.Participators, participator)
	if pos == -1 {
		return errors.New("该会议没有该参与者")
	}
	newParticipators := meeting.Participators[0:pos] + meeting.Participators[pos+len(participator)+1:]
	_, err = db.Exec(updateParticipators, newParticipators, title)
	checkErr(err)
	return err
}

// QueryMeetingsByTime .
func QueryMeetingsByTime(username, stime, etime string) ([]Meeting, error) {
	var err error
	rows, err := db.Query(queryMeetingsByTime, username, "%"+username+"%", stime, etime)
	checkErr(err)
	defer rows.Close()
	meetings := make([]Meeting, 0, 0)
	for rows.Next() {
		meeting := Meeting{}
		err = rows.Scan(&meeting.Title, &meeting.Sponsor, &meeting.Participators, &meeting.Stime, &meeting.Etime)
		checkErr(err)
		meetings = append(meetings, meeting)
	}
	return meetings, err
}

// DeleteMeeting .
// 如果是发起者，则删除会议
// 如果是参与者，则退出会议
func DeleteMeeting(username, title string) error {
	var err error
	meeting := Meeting{}
	row := db.QueryRow(queryMeetingByTitle, title)
	row.Scan(&meeting.Title, &meeting.Sponsor, &meeting.Participators, &meeting.Stime, &meeting.Etime)
	if meeting.Title == "" {
		return errors.New("该会议不存在")
	}
	if meeting.Sponsor == username {
		_, err = db.Exec(deleteMeetingByUsername, username)
		checkErr(err)
	} else if pos := strings.Index(meeting.Participators, username); pos != -1 {
		newParticipators := meeting.Participators[0:pos] + meeting.Participators[pos+len(username)+1:]
		_, err = db.Exec(updateParticipators, newParticipators, title)
		checkErr(err)
	} else {
		err = errors.New("用户并没有参与该会议")
	}
	return err
}

// DeleteMeetings .
// 如果是发起者，则删除会议
// 如果是参与者，则退出会议
func DeleteMeetings(username string) error {
	var err error
	rows, err := db.Query(queryMeetingsByUsername, username, "%"+username+"%")
	checkErr(err)
	defer rows.Close()
	meetings := make([]Meeting, 0, 0)
	for rows.Next() {
		meeting := Meeting{}
		err = rows.Scan(&meeting.Title, &meeting.Sponsor, &meeting.Participators, &meeting.Stime, &meeting.Etime)
		checkErr(err)
		meetings = append(meetings, meeting)
	}
	for _, meeting := range meetings {
		if meeting.Sponsor == username {
			_, err = db.Exec(deleteMeetingByUsername, username)
			checkErr(err)
		} else {
			pos := strings.Index(meeting.Participators, username)
			newParticipators := meeting.Participators[0:pos] + meeting.Participators[pos+len(username)+1:]
			_, err = db.Exec(updateParticipators, newParticipators, meeting.Title)
			checkErr(err)
		}
	}
	return err
}
