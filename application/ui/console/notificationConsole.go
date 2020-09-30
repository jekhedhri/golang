package console

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jekhedhri/tuto006/application/infra/product"
	"github.com/manifoldco/promptui"
)

// Run prompt
func Run() {
	var result string = ""
	for {
		menu(result)
		if result == "exit" {
			break
		}
	}
}

func menu(result string) {

	var err error

	actions := []string{"exit", "Find", "Create"}

	prompt := promptui.Select{
		Label: "Choose an action :",
		Items: actions,
	}

	_, result, err = prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	if result == "Find" {
		notificationService, _ := product.InjectProductService()

		prompt := promptui.Prompt{
			Label: "Id",
		}

		id, _ := prompt.Run()

		idInt, _ := strconv.Atoi(id)

		result, _ := notificationService.Find(idInt)

		fmt.Printf("Found %q\n", *result)

	}

	if result == "exit" {
		fmt.Printf("Shutdown")
		os.Exit(0)
	}

}
