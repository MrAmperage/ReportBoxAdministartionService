package ORM

import (
	"time"

	"github.com/MrAmperage/GoWebStruct/WebCore/Modules/ORMModule"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type User struct {
	Id        uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Username  string    `gorm:"primaryKey;not null;username"`
	Password  string    `gorm:"not null;password"`
	Enabled   bool      `gorm:"not null;enabled"`
	RoleId    string    `gorm:"not null;rolename"`
	Role      Role      `gorm:"foreignkey:Id;references:RoleId"`
	StartDate time.Time `gorm:"not null;start_date"`
	EndDate   time.Time `gorm:"not null;end_date"`
}

type UserORM struct {
	ORMModule.ORM
}

func (UserORM *UserORM) DeleteUser(Uuid uuid.UUID) (SQLResult *gorm.DB) {

	return UserORM.ConnectionLink.Delete(&User{Id: Uuid})

}
func (UserORM *UserORM) AddUser(NewUser User) (SQLResult *gorm.DB) {

	return UserORM.ConnectionLink.Create(&NewUser)
}

func (UserORM *UserORM) EditUser(NewUser User) (SQLResult *gorm.DB) {

	return UserORM.ConnectionLink.Updates(&NewUser)
}

func (UserORM *UserORM) GetUsers() (Users []User, Error error) {

	Error = UserORM.ConnectionLink.Find(&Users).Error
	if Error != nil {

		return
	}
	for Index := range Users {
		Error = UserORM.ConnectionLink.Model(Users[Index]).Association("Role").Find(&Users[Index].Role)
		if Error != nil {
			return
		}

	}
	return Users, Error
}

func (UserORM *UserORM) GetUser(Uuid uuid.UUID) (User User, Error error) {
	User.Id = Uuid
	return User, UserORM.ConnectionLink.Take(&User).Error
}
