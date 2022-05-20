package main

import (
	"cademo/caserver"
	"cademo/config"
	"cademo/controller"
	"cademo/log"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

var logger = log.GetLogger("info")

func main() {
	logger.Info("Fabric Ca Demo.")

	caServer := caserver.GetServer()
	if err := caServer.Init(false); err != nil {
		panic(err)
	}
	if err := caServer.Start(); err != nil {
		panic(err)
	}
	logger.Info("Ca server started.")

	if err := caserver.EnrollAdmin(); err != nil {
		panic(err)
	}
	logger.Info("Ca admin Enrolled")

	controller := controller.NewController(caServer)

	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/cainfo", controller.CaInfo)
		// identity api
		v1.POST("/id/register", controller.Register)
		v1.GET("/id/all", controller.AllIdentities)
		v1.POST("/id/del", controller.DeleteIdentity)

		// certificate api
		v1.POST("/cert/enroll", controller.Enroll)
		v1.POST("/cert/reenroll", controller.ReEnroll)
		v1.POST("/cert/enrolltls", controller.EnrollTls)
		v1.POST("/cert/list", controller.CertificateList)

		// affiliation api
		v1.GET("/affi/all", controller.AllAffiliations)
	}
	port := config.C.GetInt("server.port")
	router.Run(fmt.Sprintf(":%d", port))
	logger.Info("Logic server started.")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-stop
	logger.Info("Demo shutdown.")
}
