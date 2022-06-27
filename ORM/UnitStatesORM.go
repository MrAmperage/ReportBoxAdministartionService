package ORM

import (
	"github.com/MrAmperage/GoWebStruct/WebCore/Modules/ORMModule"
	"github.com/gofrs/uuid"
)

type UnitState struct {
	Id      uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Caption string    `gorm:"not null;caption"`
}
type UnitStatesORM struct {
	ORMModule.ORM
}

func (UnitStatesORM *UnitStatesORM) GetUnitStates() (UnitStates []UnitState, Error error) {
	return UnitStates, UnitStatesORM.ConnectionLink.Find(&UnitStates).Error
}

func (UnitStatesORM *UnitStatesORM) GetUnitState(UUID uuid.UUID) (UnitState UnitState, Error error) {
	UnitState.Id = UUID
	return UnitState, UnitStatesORM.ConnectionLink.Take(&UnitState).Error
}
