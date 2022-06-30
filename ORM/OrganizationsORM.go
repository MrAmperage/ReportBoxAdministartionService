package ORM

import (
	"github.com/MrAmperage/GoWebStruct/WebCore/Modules/ORMModule"
	"github.com/gofrs/uuid"
)

type Organization struct {
	Id       uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Caption  string    `gorm:"not null;caption"`
	Internal bool      `gorm:"not null;internal"`
}
type OrganizationsORM struct {
	ORMModule.ORM
}

func (OrganizationsORM *OrganizationsORM) GetOrganizations() (Organizations []Organization, Error error) {
	return Organizations, OrganizationsORM.ConnectionLink.Find(&Organizations).Error
}

func (OrganizationsORM *OrganizationsORM) GetOrganization(UUID uuid.UUID) (Organization Organization, Error error) {
	Organization.Id = UUID
	return Organization, OrganizationsORM.ConnectionLink.Take(&Organization).Error
}

func (OrganizationsORM *OrganizationsORM) DeleteOrganization(OrganizationId uuid.UUID) (Error error) {
	return OrganizationsORM.ConnectionLink.Delete(&Organization{Id: OrganizationId}).Error
}
