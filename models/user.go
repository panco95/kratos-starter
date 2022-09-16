package models

import "time"

type UserStatus string

var (
	UserStatusNormal UserStatus = "normal"
	UserStatusLock   UserStatus = "lock"
)

type User struct {
	Model
	Username      string     `gorm:"column:username;not null;default:'';type:varchar(50);index:username" json:"username" binding:"max=50"`
	Password      string     `gorm:"column:password;not null;default:'';type:varchar(200)" json:"password"`
	PasswordSalt  string     `gorm:"column:password_salt;not null;default:'';type:varchar(200)" json:"passwordSalt"`
	Mobile        string     `gorm:"column:mobile;not null;default:'';type:varchar(50)" json:"mobile"`
	Remark        string     `gorm:"column:remark;not null;default:'';type:varchar(1000)" json:"remark"`
	Status        UserStatus `gorm:"column:status;not null;default:'normal';type:varchar(50)" json:"status"`
	LastLoginTime *time.Time `gorm:"column:last_login_time;" json:"lastLoginTime"`
	LastLoginIp   string     `gorm:"column:last_login_ip;not null;default:'';type:varchar(20)" json:"lastLoginIp"`
	LoginTimes    uint       `gorm:"column:login_times;not null;default:0;type:int(10)" json:"loginTimes"`
}
