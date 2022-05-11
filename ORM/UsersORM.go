package ORM

import (
	"github.com/MrAmperage/GoWebStruct/WebCore/Modules/ORMModule"
	"gorm.io/gorm"
)

type User struct {
	Username string `gorm:"primaryKey;not null;username"`
	Password string `gorm:"not null;password"`
	Enabled  bool   `gorm:"not null;enabled"`
	Rolename string `gorm:"not null;rolename"`
}

type UserORM struct {
	ORMModule.ORM
}

func (UserORM *UserORM) DeleteUser(Username string) (SQLResult *gorm.DB) {

	return UserORM.ConnectionLink.Delete(&User{Username: Username})

}
func (UserORM *UserORM) AddUser(NewUser User) (SQLResult *gorm.DB) {

	return UserORM.ConnectionLink.Create(NewUser)
}
