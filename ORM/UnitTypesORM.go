package ORM

import (
	"github.com/MrAmperage/GoWebStruct/WebCore/Modules/ORMModule"
	"github.com/gofrs/uuid"
)

type UnitType struct {
	Id      uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Caption string    `gorm:"not null;caption"`
}
type UnitTypesORM struct {
	ORMModule.ORM
}

func (UnitTypesORM *UnitTypesORM) GetUnitTypes() (UnitTypes []UnitType, Error error) {
	return UnitTypes, UnitTypesORM.ConnectionLink.Find(&UnitTypes).Error
}
func (UnitTypesORM *UnitTypesORM) GetUnitType(UUID uuid.UUID) (UnitType UnitType, Error error) {
	UnitType.Id = UUID
	return UnitType, UnitTypesORM.ConnectionLink.Take(&UnitType).Error
}
func (UnitTypesORM *UnitTypesORM) DeleteUnitType(UnitTypeId uuid.UUID) (Error error) {
	return UnitTypesORM.ConnectionLink.Delete(&UnitType{Id: UnitTypeId}).Error
}

func (UnitTypesORM *UnitTypesORM) AddUnitType(NewUnitType UnitType) (Error error) {
	return UnitTypesORM.ConnectionLink.Create(&NewUnitType).Error
}

func (UnitTypesORM *UnitTypesORM) EditUnitType(NewUnitType UnitType) (Error error) {
	return UnitTypesORM.ConnectionLink.Updates(&NewUnitType).Error
}
