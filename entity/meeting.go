package entity

import (
	"errors"
	"strings"
)

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
var deleteMeetings = `DELETE FROM meeting WHERE sponsor = ? OR participators LIKE ?;`
var deleteMeetingByTitle = `DELETE FROM meeting WHERE title = ?;`
var queryMeetingByTitle = `SELECT * FROM meeting WHERE meeting.title = ?;`
var queryMeetings = `SELECT * FROM meeting AS m WHERE (m.sponsor = ? OR m.participators LIKE ?) AND (m.stime >= ? AND m.etime <= ?);`
var addParticipator = `UPDATE meeting SET participators = participators || ? WHERE title = ?;`
var updateParticipators = `UPDATE meeting SET participators = ? WHERE title = ?;`

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

// 删除用户举办或参与的所有会议
func DeleteMeetings(username string) error {
	var err error
	_, err = db.Exec(deleteMeetings, username, "%"+username+"%")
	checkErr(err)
	return err
}

func QueryMeetings(username, stime, etime string) ([]Meeting, error) {
	var err error
	rows, err := db.Query(queryMeetings, username, "%"+username+"%", stime, etime)
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

func AddParticipator(username, title, participator string) error {
	var err error

	// TODO: 判断 username 是否是举办者
	// TODO: 判断 participant 是否已经是参与者

	_, err = db.Exec(addParticipator, participator+",", title)
	checkErr(err)
	return err
}

func DeleteParticipator(username, title, participator string) error {
	var err error
	meeting := Meeting{}
	row := db.QueryRow(queryMeetingByTitle, username)
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

// 如果是发起者，则删除会议
// 如果是参与者，则退出会议
func DeleteMeeting(username, title string) error {
	var err error
	meeting := Meeting{}
	row := db.QueryRow(queryMeetingByTitle, username)
	row.Scan(&meeting.Title, &meeting.Sponsor, &meeting.Participators, &meeting.Stime, &meeting.Etime)
	if meeting.Title == "" {
		return errors.New("该会议不存在")
	}
	if meeting.Sponsor == username {
		_, err = db.Exec(deleteMeetingByTitle, title)
		checkErr(err)
		return err
	} else if pos := strings.Index(meeting.Participators, username); pos != -1 {
		newParticipators := meeting.Participators[0:pos] + meeting.Participators[pos+len(username)+1:]
		_, err = db.Exec(updateParticipators, newParticipators, title)
		checkErr(err)
		return err
	} else {
		return errors.New("与该会议无关系")
	}
}
