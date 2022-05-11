package ORM

import "github.com/MrAmperage/GoWebStruct/WebCore/Modules/ORMModule"

type User struct {
	Username string `gorm:"primaryKey;not null;username"`
	Password string `gorm:"not null;password"`
	Enabled  bool   `gorm:"not null;enabled"`
	Rolename string `gorm:"not null;rolename"`
}

type UserORM struct {
	ORMModule.ORM
}

func (UserORM *UserORM) AddUser(NewUser User) {

	UserORM.ConnectionLink.Create(NewUser)
}
