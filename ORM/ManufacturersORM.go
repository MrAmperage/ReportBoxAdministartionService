package ORM

import (
	"github.com/MrAmperage/GoWebStruct/WebCore/Modules/ORMModule"
	"github.com/gofrs/uuid"
)

type Manufacturer struct {
	Id      uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Caption string    `gorm:"not null;caption"`
}
type ManufacturersORM struct {
	ORMModule.ORM
}

func (ManufacturersORM *ManufacturersORM) GetManufacturers() (Manufacturers []Manufacturer, Error error) {
	return Manufacturers, ManufacturersORM.ConnectionLink.Find(&Manufacturers).Error
}

func (ManufacturersORM *ManufacturersORM) GetManufacturer(UUID uuid.UUID) (Manufacturer Manufacturer, Error error) {
	Manufacturer.Id = UUID
	return Manufacturer, ManufacturersORM.ConnectionLink.Take(&Manufacturer).Error
}

func (ManufacturersORM *ManufacturersORM) DeleteManufacturer(ManufacturerId uuid.UUID) (Error error) {
	return ManufacturersORM.ConnectionLink.Delete(&Manufacturer{Id: ManufacturerId}).Error
}

func (ManufacturersORM *ManufacturersORM) EditManufacturer(NewManufacturer Manufacturer) (Manufacturer, error) {
	return NewManufacturer, ManufacturersORM.ConnectionLink.Updates(&NewManufacturer).Error
}
