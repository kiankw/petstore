package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kiankw/petstore/handler"
	"github.com/kiankw/petstore/service"
	"github.com/kiankw/petstore/service/dal"
)

func main() {
	dalClient := dal.GetDAL()
	serviceStruct := service.NewService(dalClient)
	httpHandler := handler.NewHandler(serviceStruct)
	r := gin.Default()

	r.GET("/pets", httpHandler.GetAllPets)
	r.POST("/pets", httpHandler.CreatePet)
	r.DELETE("/pets/:id", httpHandler.DeletePet)
	r.GET("/pets/:id", httpHandler.GetPetByID)
	r.PATCH("pets/:id", httpHandler.UpdatePet)

	r.Run(":8082")
}
