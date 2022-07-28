package ORM

import (
	"github.com/MrAmperage/GoWebStruct/WebCore/Modules/ORMModule"
	"github.com/gofrs/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Role struct {
	Id                  uuid.UUID      `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Rolename            string         `gorm:"not null;rolename"`
	MenusAccess         pq.StringArray `gorm:"not null;type:varchar[];default:array[]::varchar[];menus_access"`
	OrganizationsAccess pq.StringArray `gorm:"not null;type:varchar[];default:array[]::varchar[];organizations_access"`
}
type RolesORM struct {
	ORMModule.ORM
}

func (RolesORM *RolesORM) GetRoles() (Roles []Role, Error error) {

	return Roles, RolesORM.ConnectionLink.Find(&Roles).Error
}
func (RolesORM *RolesORM) AddRole(NewRole Role) (SQLResult *gorm.DB) {

	return RolesORM.ConnectionLink.Create(&NewRole)
}
func (RolesORM *RolesORM) GetRole(UUID uuid.UUID) (Role Role, Error error) {
	Role.Id = UUID
	return Role, RolesORM.ConnectionLink.Take(&Role).Error
}

func (RolesORM *RolesORM) EditRole(NewRole Role) (Role, error) {
	return NewRole, RolesORM.ConnectionLink.Save(&NewRole).Error
}

func (RolesORM *RolesORM) DeleteRole(Uuid uuid.UUID) (SQLResult *gorm.DB) {

	return RolesORM.ConnectionLink.Delete(&Role{Id: Uuid})

}
