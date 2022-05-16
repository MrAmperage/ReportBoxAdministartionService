package ORM

import (
	"github.com/MrAmperage/GoWebStruct/WebCore/Modules/ORMModule"
	"github.com/gofrs/uuid"
	"github.com/jackc/pgtype"
	"gorm.io/gorm"
)

type Scheme struct {
	Id      uuid.UUID    `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Caption string       `gorm:"not null;caption"`
	Scheme  pgtype.JSONB `gorm:"not null;scheme"`
}

type SchemeORM struct {
	ORMModule.ORM
}

func (SchemeORM *SchemeORM) AddScheme(NewScheme Scheme) (SQLResult *gorm.DB) {

	return SchemeORM.ConnectionLink.Create(&NewScheme)
}

func (SchemeORM *SchemeORM) DeleteScheme(SchemeId uuid.UUID) (SQLResult *gorm.DB) {

	return SchemeORM.ConnectionLink.Delete(&Scheme{Id: SchemeId})
}

func (SchemeORM *SchemeORM) GetShemes() (Schemes []Scheme, Error error) {
	return Schemes, SchemeORM.ConnectionLink.Find(&Schemes).Error
}
func (SchemeORM *SchemeORM) EditScheme(NewScheme Scheme) (SQLResult *gorm.DB) {

	return SchemeORM.ConnectionLink.Updates(&NewScheme)
}
