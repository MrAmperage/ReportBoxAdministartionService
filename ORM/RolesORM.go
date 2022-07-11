package ORM

import (
	"github.com/MrAmperage/GoWebStruct/WebCore/Modules/ORMModule"
	"github.com/gofrs/uuid"
)

type Role struct {
	Id       uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Rolename string    `gorm:"not null;rolename"`
}
type RolesORM struct {
	ORMModule.ORM
}

func (RolesORM *RolesORM) GetRoles() (Roles []Role, Error error) {

	return Roles, RolesORM.ConnectionLink.Find(&Roles).Error
}

func (RolesORM *RolesORM) GetRole(UUID uuid.UUID) (Role Role, Error error) {
	Role.Id = UUID
	return Role, RolesORM.ConnectionLink.Take(&Role).Error
}
