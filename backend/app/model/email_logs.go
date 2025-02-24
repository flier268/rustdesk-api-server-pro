package model

import "time"

type EmailLogs struct {
	Id        int       `xorm:"'id' int notnull pk autoincr"`
	UserId    int       `xorm:"'user_id' int"`
	TplId     int       `xorm:"'tpl_id' int"`
	From      string    `xorm:"'from' varchar(255)"`
	To        string    `xorm:"'to' varchar(255)"`
	Subject   string    `xorm:"'subject' varchar(255)"`
	Contents  string    `xorm:"'contents' mediumtext"`
	Status    int       `xorm:"'status' tinyint"` // 1=ok,2=err
	CreatedAt time.Time `xorm:"'created_at' datetime created"`
	UpdatedAt time.Time `xorm:"'updated_at' datetime updated"`
}

func (m *EmailLogs) TableName() string {
	return "email_logs"
}
