package main

import (
	"iotaVisionGo/router"
	"iotaVisionGo/services"
	"iotaVisionGo/utils"
	"log"
	"net/http"
)

func main() {
	log.Println("In Main App")

	var dbconn = utils.GetConnection()
	services.SetDB(dbconn)
	var appRouter = router.CreateRouter()
	utils.CreateTablesIfNotExist()

	log.Println("Listening on Port 8000")
	log.Fatal(http.ListenAndServe(":8000", appRouter))
}
