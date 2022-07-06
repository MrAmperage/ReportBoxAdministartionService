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
