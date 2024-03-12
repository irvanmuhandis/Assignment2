package main

import (
	"assignment2/database"
	_ "assignment2/docs"
	"assignment2/router"
	"fmt"
	"log"
)

// @title           Assignment 2 REST API M.Irvan Muhandis
// @version         1.0
// @description     This is a REST API created to fullfil assignment class
// @termsOfService  http://swagger.io/terms/
// @contact.name   M.Irvan Muhandis
// @contact.url    https://wa.me/6285701514915
// @contact.email  irvanmuhandis@gmail.com
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @host      localhost:8080
// @BasePath  /
// @securityDefinitions.basic  BasicAuth
// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	//Start database
	database.StartDB()
	PORT := ":8080"
	fmt.Println("Server Start at PORT : ", PORT)

	//Run webservice
	err := router.MyRouter().Run(PORT)
	if err != nil {
		log.Fatal(err)
	}
}
