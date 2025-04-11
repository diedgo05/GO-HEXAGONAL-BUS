package dependencies

import (
	"bus-project/src/buses/application"
	"bus-project/src/buses/application/services"
	"bus-project/src/buses/infraestructure"
	"bus-project/src/buses/infraestructure/adapters"
	"bus-project/src/buses/infraestructure/http/controllers"
	"bus-project/src/core"
)

var (
	mySQL infraestructure.MySQL
	eventService *services.Event
)

func InitBus() {
	db,err := core.InitMySQL()
	if err != nil {
		return
	}
	mySQL = *infraestructure.NewMySQL(db)
	rabbit := adapters.NewRabbit()
	eventService = services.NewEvent(rabbit)
}

func AddBusController() *controllers.AddBusController {
	ucAddBus := application.NewAddBusUseCase(&mySQL)

	return controllers.NewAddBusController(ucAddBus, eventService)
}

func GetAllBusesController() *controllers.GetAllBusesController {
	ucGetAllBuses := application.NewGetAllBusesUseCase(&mySQL)

	return controllers.NewGetAllBusesController(ucGetAllBuses)
}

func UpdateBusController() *controllers.UpdateBusByIDController {
	ucUpdateBus := application.NewUpdateBusByIDUseCase(&mySQL)

	return controllers.NewUpdateBusByIDController(ucUpdateBus)
}

func GetBusByIdChoferController() *controllers.GetBusByIdChoferController {
	ucGetBusByChofer := application.NewFindBusByIdChoferUseCase(&mySQL)

	return controllers.NewGetBusByIdChoferController(ucGetBusByChofer)
}

func DeleteBusController() *controllers.DeleteBusByIDController {
	ucDeleteBus := application.NewDeleteBusByIDUseCase(&mySQL)

	return controllers.NewDeleteBusByIDController(ucDeleteBus)
}