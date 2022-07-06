package ORM

import (
	"github.com/MrAmperage/GoWebStruct/WebCore/Modules/ORMModule"
	"github.com/gofrs/uuid"
)

type CargoType struct {
	Id      uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Caption string    `gorm:"not null;caption"`
}
type CargoTypesORM struct {
	ORMModule.ORM
}

func (CargoTypesORM *CargoTypesORM) GetCargoTypes() (CapgoTypes []CargoType, Error error) {
	return CapgoTypes, CargoTypesORM.ConnectionLink.Find(&CapgoTypes).Error
}

func (CargoTypesORM *CargoTypesORM) GetCargoType(UUID uuid.UUID) (CargoType CargoType, Error error) {
	CargoType.Id = UUID
	return CargoType, CargoTypesORM.ConnectionLink.Take(&CargoType).Error
}
func (CargoTypesORM *CargoTypesORM) DeleteCargoType(CargoTypeId uuid.UUID) (Error error) {
	return CargoTypesORM.ConnectionLink.Delete(&CargoType{Id: CargoTypeId}).Error
}

func (CargoTypesORM *CargoTypesORM) EditCargoType(NewCargoType CargoType) (CargoType, error) {
	return NewCargoType, CargoTypesORM.ConnectionLink.Updates(&NewCargoType).Error
}
