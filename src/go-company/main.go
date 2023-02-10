package main

import (
	"fmt"
	kafkaconfig "go-company/kafkaConfig"
	"go-company/models"
	"go-company/rest"
)

func main() {
	fmt.Println("hello world")

		// Run setup.
		setup()

		// Wait for shutdown.
		//system.WaitShutdown(context.TODO())
	
		// Run cleanup.
		//cleanup()

}


func setup(){

	//connect to db
	models.SetupDatabase()

	//connect to  the kafka producer
	kafkaconfig.SetupKafkaProducer()

	//start web-server
	server := rest.NewApiServer()
	if err  := server.ListenAndServe(); err != nil{
		panic(fmt.Sprintf("failed to start web  server %s",err.Error()))
	}

}