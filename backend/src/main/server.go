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

	offerRepository repository.OfferRepository = repository.NewOfferRepository()
	offerService    service.OfferService       = service.NewOfferService(offerRepository)
	offerController controller.OfferController = controller.NewOfferController(offerService)

	commentRepository repository.CommentRepository = repository.NewCommentRepository()
	commentService    service.CommentService       = service.NewCommentService(commentRepository)
	commentController controller.CommentController = controller.NewCommentController(commentService)

	rateRepository repository.RateRepository = repository.NewRateRepository()
	rateService    service.RateService       = service.NewRateService(rateRepository)
	rateController controller.RateController = controller.NewRateController(rateService)

	httpRouter router.Router = router.NewMuxRouter()
)

func runServer() {
	const port string = ":8000"

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(w, "Up and runing")
	})

	// Vehicles
	httpRouter.GET("/vehicles", vehicleController.GetAll)
	httpRouter.POST("/vehicles", vehicleController.Save)
	httpRouter.DELETE("/vehicles/delete-all", vehicleController.DeleteAll)

	// Offers
	httpRouter.GET("/offers", offerController.GetAll)

	// Comments
	httpRouter.POST("/comments", commentController.Save)

	// Rates
	httpRouter.POST("/rates", rateController.Save)

	httpRouter.SERVE(port)
}
