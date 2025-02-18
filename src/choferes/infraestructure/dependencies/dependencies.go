package dependencies

import (
	"bus-project/src/choferes/application"
	"bus-project/src/choferes/infraestructure"
	"bus-project/src/choferes/infraestructure/http/controllers"
	"bus-project/src/core"
	"fmt"
)

var (
	mySQL infraestructure.MySQL
)

func Init() {
	db, err := core.InitMySQL()
	if err != nil {
		fmt.Println("Error de servidor")
		return
	}

	//defer db.Close()
	mySQL = *infraestructure.NewMySQL(db)
}

func AddChoferController() *controllers.AddChoferController {
	ucAddChofer := application.NewAddChoferUseCase(&mySQL)

	return controllers.NewAddChoferController(ucAddChofer)
}

func GetAllChoferesController() *controllers.GetAllChoferesController {
	ucGetAllChoferes := application.NewGetAllChoferesUseCase(&mySQL)

	return controllers.NewGetAllChoferesController(ucGetAllChoferes)
}

func UpdateChoferController() *controllers.UpdateByIDChoferController {
	ucUpdateChofer := application.NewUpdateByIDChoferUseCase(&mySQL)

	return controllers.NewUpdateByIDChoferController(ucUpdateChofer)
}

func DeleteChoferController() *controllers.DeleteByIDChoferController {
	ucDeleteChofer := application.NewDeleteByIDChoferUseCase(&mySQL)

	return controllers.NewDeleteByIDChoferController(ucDeleteChofer)
}

func GetChoferByidController() *controllers.GetChoferByIdController {
	ucGetChoferByID := application.NewGetChoferByIDUseCase(&mySQL)

	return controllers.NewGetChoferByIDController(ucGetChoferByID)
}