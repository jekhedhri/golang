package application

import (
	"fmt"
	"os"

	"github.com/jekhedhri/tuto006/application/ui/console"
	configuration "github.com/jekhedhri/tuto006/configurations"
	router "github.com/jekhedhri/tuto006/configurations/routes"
)

//Run application
func Run() {
	configuration.Init()

	if os.Getenv("ui") == "cli" {
		fmt.Println("-------------------------------")
		fmt.Println("App is running on console!!")
		fmt.Println("-------------------------------")
		console.Run()
	} else {
		fmt.Println("-------------------------------")
		fmt.Println("App is running on port : ", os.Getenv("PORT"))
		fmt.Println("-------------------------------")
		router.Init()

	}

}
