package ORM

import (
	"github.com/MrAmperage/GoWebStruct/WebCore/Modules/ORMModule"
	"github.com/gofrs/uuid"
)

type TransportType struct {
	Id      uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Caption string    `gorm:"not null;caption"`
}
type TransportTypesORM struct {
	ORMModule.ORM
}

func (TransportTypesORM *TransportTypesORM) GetTransportTypes() (TransportTypes []TransportType, Error error) {
	return TransportTypes, TransportTypesORM.ConnectionLink.Find(&TransportTypes).Error
}

func (TransportTypesORM *TransportTypesORM) GetTransportType(UUID uuid.UUID) (TransportType TransportType, Error error) {
	TransportType.Id = UUID
	return TransportType, TransportTypesORM.ConnectionLink.Take(&TransportType).Error
}

func (TransportTypesORM *TransportTypesORM) EditTransportType(NewTransportType TransportType) (TransportType, error) {
	return NewTransportType, TransportTypesORM.ConnectionLink.Save(&NewTransportType).Error
}

func (TransportTypesORM *TransportTypesORM) DeleteTransportType(TransportTypeId uuid.UUID) (Error error) {
	return TransportTypesORM.ConnectionLink.Delete(&TransportType{Id: TransportTypeId}).Error
}

func (TransportTypesORM *TransportTypesORM) AddTransportType(NewTransportType TransportType) (TransportType, error) {
	return NewTransportType, TransportTypesORM.ConnectionLink.Create(&NewTransportType).Error
}
