package main

import (
	"cademo/caserver"
	"cademo/controller"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Fabric Ca Demo.")

	caServer := caserver.GetServer()
	if err := caServer.Init(false); err != nil {
		panic(err)
	}
	if err := caServer.Start(); err != nil {
		panic(err)
	}
	fmt.Println("Ca server started.")

	controller := controller.NewController(caServer)

	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/cainfo", controller.CaInfo)
		v1.POST("/register", controller.Register)
		v1.POST("/enroll", controller.Enroll)
	}
	router.Run(":9300")
	fmt.Println("Logic server started.")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-stop
	fmt.Println("Demo shutdown.")
}
