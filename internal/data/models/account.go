package models

import "time"

type AccountStatus string

var (
	AccountStatusNormal AccountStatus = "normal"
	AccountStatusLock   AccountStatus = "lock"
)

type Account struct {
	Model
	Username      string        `gorm:"column:username;not null;default:'';type:varchar(50);index:username" json:"username" binding:"max=50"`
	Password      string        `gorm:"column:password;not null;default:'';type:varchar(200)" json:"password"`
	PasswordSalt  string        `gorm:"column:password_salt;not null;default:'';type:varchar(200)" json:"passwordSalt"`
	Remark        string        `gorm:"column:remark;not null;default:'';type:varchar(1000)" json:"remark"`
	LastLoginTime *time.Time    `gorm:"column:last_login_time;" json:"lastLoginTime"`
	LastLoginIp   string        `gorm:"column:last_login_ip;not null;default:'';type:varchar(20)" json:"lastLoginIp"`
	LoginTimes    uint          `gorm:"column:login_times;not null;default:0;type:int(10)" json:"loginTimes"`
	Status        AccountStatus `gorm:"column:status;not null;default:'normal';type:varchar(50)" json:"status"`
	Mobile        string        `gorm:"column:mobile;not null;default:'';type:varchar(50)" json:"mobile"`
}
