package main

import (
	"fmt"
	"os"

	"github.com/MrAmperage/GoWebStruct/ApplicationCore"
	"github.com/MrAmperage/ReportBoxAdministartionService/Controllers"
	"github.com/MrAmperage/ReportBoxAdministartionService/ORM"
)

func main() {

	AuthenticationService := &ApplicationCore.ApplicationCore{}
	ErrorInitService := AuthenticationService.Init()
	if ErrorInitService != nil {
		fmt.Println(ErrorInitService)
		os.Exit(0)
	}
	OrganizationsORM := &ORM.OrganizationsORM{}
	OrganizationsORM.SetName("OrganizationsORM")
	ManufacturersORM := &ORM.ManufacturersORM{}
	ManufacturersORM.SetName("ManufacturersORM")
	SchemeORM := &ORM.SchemeORM{}
	SchemeORM.SetName("SchemeORM")
	UnitStatesORM := &ORM.UnitStatesORM{}
	UnitStatesORM.SetName("UnitStatesORM")
	UserORM := &ORM.UserORM{}
	UserORM.SetName("UserORM")
	UnitTypesORM := &ORM.UnitTypesORM{}
	UnitTypesORM.SetName("UnitTypesORM")
	TopMenuORM := &ORM.TopMenuORM{}
	TopMenuORM.SetName("TopMenuORM")
	TransportTypeORM := &ORM.TransportTypesORM{}
	TransportTypeORM.SetName("TransportTypesORM")
	CargoTypesORM := &ORM.CargoTypesORM{}
	CargoTypesORM.SetName("CargoTypesORM")

	ReportBoxDatabase, Error := AuthenticationService.WebCore.PostgreSQL.FindByName("ReportBoxDatabase")
	if Error != nil {
		fmt.Println(Error)
	}
	ReportBoxDatabase.ORMs.Add(CargoTypesORM)
	ReportBoxDatabase.ORMs.Add(ManufacturersORM)
	ReportBoxDatabase.ORMs.Add(TransportTypeORM)
	ReportBoxDatabase.ORMs.Add(OrganizationsORM)
	ReportBoxDatabase.ORMs.Add(TopMenuORM)
	ReportBoxDatabase.ORMs.Add(UnitTypesORM)
	ReportBoxDatabase.ORMs.Add(UserORM)
	ReportBoxDatabase.ORMs.Add(SchemeORM)
	ReportBoxDatabase.ORMs.Add(UnitStatesORM)
	ErrorDatabaseConnection := AuthenticationService.WebCore.PostgreSQL.StartDatabaseConnections()
	if ErrorDatabaseConnection != nil {

		fmt.Println(ErrorDatabaseConnection)
		os.Exit(0)
	}

	ErrorRabbitMQ := AuthenticationService.StartRabbitMQ()
	if ErrorRabbitMQ != nil {

		fmt.Println(ErrorRabbitMQ)
		os.Exit(0)
	}

	Subscribe, Error := AuthenticationService.WebCore.RabbitMQ.RabbitMQChanel.GetSubscribeByQueueName("AdministartionQueue")
	if Error != nil {
		fmt.Println(Error)
	}
	//Производители
	Subscribe.MessageEmmiter.Handler("Manufacturers", Controllers.AddManufacturer).Method("POST")
	Subscribe.MessageEmmiter.Handler("Manufacturers", Controllers.GetManufacturers).Method("GET")
	Subscribe.MessageEmmiter.Handler("Manufacturers", Controllers.DeleteManufacturer).Method("DELETE")
	Subscribe.MessageEmmiter.Handler("Manufacturers", Controllers.EditManufacturer).Method("PATCH")
	//Типы транспорта
	Subscribe.MessageEmmiter.Handler("TransportTypes", Controllers.GetTransportTypes).Method("GET")
	Subscribe.MessageEmmiter.Handler("TransportTypes", Controllers.EditTransportType).Method("PATCH")
	Subscribe.MessageEmmiter.Handler("TransportTypes", Controllers.DeleteTransportType).Method("DELETE")
	Subscribe.MessageEmmiter.Handler("TransportTypes", Controllers.AddTransportType).Method("POST")
	//Состояние агрегатов
	Subscribe.MessageEmmiter.Handler("UnitStates", Controllers.GetUnitStates).Method("GET")
	Subscribe.MessageEmmiter.Handler("UnitStates", Controllers.DeleteUnitState).Method("DELETE")
	Subscribe.MessageEmmiter.Handler("UnitStates", Controllers.EditUnitState).Method("PATCH")
	Subscribe.MessageEmmiter.Handler("UnitStates", Controllers.AddUnitState).Method("POST")
	//Типы агрегатов
	Subscribe.MessageEmmiter.Handler("UnitTypes", Controllers.GetUnitTypes).Method("GET")
	Subscribe.MessageEmmiter.Handler("UnitTypes", Controllers.DeleteUnitType).Method("DELETE")
	Subscribe.MessageEmmiter.Handler("UnitTypes", Controllers.AddUnitType).Method("POST")
	Subscribe.MessageEmmiter.Handler("UnitTypes", Controllers.EditUnitType).Method("PATCH")
	//Пользователи
	Subscribe.MessageEmmiter.Handler("Users", Controllers.GetUsers).Method("GET")
	Subscribe.MessageEmmiter.Handler("Users", Controllers.AddUser).Method("POST")
	Subscribe.MessageEmmiter.Handler("Users", Controllers.DeleteUser).Method("DELETE")
	Subscribe.MessageEmmiter.Handler("Users", Controllers.EditUser).Method("PATCH")
	//Организации
	Subscribe.MessageEmmiter.Handler("Organizations", Controllers.GetOrganizations).Method("GET")
	Subscribe.MessageEmmiter.Handler("Organizations", Controllers.DeleteOrganization).Method("DELETE")
	Subscribe.MessageEmmiter.Handler("Organizations", Controllers.EditOrganization).Method("PATCH")
	Subscribe.MessageEmmiter.Handler("Organizations", Controllers.AddOrganization).Method("POST")
	//Конфигурация
	Subscribe.MessageEmmiter.Handler("Configurations", Controllers.Configuration).Method("GET")
	//Тиы грузов
	Subscribe.MessageEmmiter.Handler("CargoTypes", Controllers.GetCargoTypes).Method("GET")
	Subscribe.MessageEmmiter.Handler("CargoTypes", Controllers.DeleteCargoType).Method("DELETE")
	Subscribe.MessageEmmiter.Handler("CargoTypes", Controllers.EditCargoType).Method("PATCH")
	Subscribe.MessageEmmiter.Handler("CargoTypes", Controllers.AddCargoType).Method("POST")

	Subscribe.MessageProcessing(&ReportBoxDatabase.ORMs)

	ErrorWebServer := AuthenticationService.StartWebServer()
	if ErrorInitService != nil {

		fmt.Println(ErrorWebServer)
		os.Exit(0)
	}
}
