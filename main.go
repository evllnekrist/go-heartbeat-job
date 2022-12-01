package main

import (
	"fmt"
	"go-heartbeat-job/database"
	"go-heartbeat-job/routers"
	"go-heartbeat-job/helpers"
)

func main() {
	var PORT = ":8080"
	database.StartDB()
	helpers.RunCronJobs_Weather()
	routers.StartServer().Run(PORT)
	fmt.Scanln()
}
