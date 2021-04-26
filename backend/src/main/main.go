package main

import (
	"controller"
	"fmt"
	"net/http"
	"repository"
	"service"

	"router"
)

var (
	vehicleRepository repository.VehicleRepository = repository.NewVehicleRepository()
	vehicleService    service.VehicleService       = service.NewVehicleService(vehicleRepository)
	vehicleController controller.VehicleController = controller.NewVehicleController(vehicleService)
	httpRouter        router.Router                = router.NewMuxRouter()
)

func main() {
	const port string = ":8000"

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(w, "Up and runing")
	})

	httpRouter.GET("/vehicles", vehicleController.GetAll)
	httpRouter.POST("/vehicles", vehicleController.Save)
	httpRouter.DELETE("/vehicles/delete-all", vehicleController.DeleteAll)

	httpRouter.SERVE(port)
}
