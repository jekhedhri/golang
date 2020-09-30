package configuration

import (
	"flag"
	"fmt"
	"os"

	"github.com/tkanos/gonfig"
)

//Configuration is global configuration
type Configuration struct {
	Name        string `env:"name"`
	Description string `env:"description"`
	Version     string `env:"version"`
	Author      string `env:"author"`
	Company     string `env:"company"`
	Env         string `env:"env"`
	UI          string `env:"ui"`
	Infra       string `env:"infra"`
	Port        string `env:"port"`
	DbDriver    string `env:"dbdriver"`
	DbUser      string `env:"dbuser"`
	DbPassword  string `env:"dbpassword"`
	DbPort      string `env:"db.port"`
	DbHost      string `env:"dbhost"`
	DbName      string `env:"dbname"`
	DbMode      string `env:"dbmode"`
	MockBaseURL string `env:"mockBaseURL"`
	MockPort    string `env:"mockPort"`
}

var (
	env   = flag.String("env", "local", "environnement")
	ui    = flag.String("ui", "api", "ui type")
	infra = flag.String("infra", "infra", "infra type")
)

// Init configurations
func Init() Configuration {

	configuration := Configuration{}

	configuration = getConfigsFromFile()

	setGlobalVariableEnv(configuration)

	printConfigsOnConsole(configuration)

	return configuration

}

func getConfigFileName() string {
	var filePath = "configurations/profiles/local.json"
	if *env == "local" {
		filePath = "configurations/profiles/local.json"
	} else if *env == "staging" {
		filePath = "configurations/profiles/staging.json"
	} else if *env == "prod" {
		filePath = "configurations/profiles/prod.json"
	}
	return filePath
}

func getConfigsFromFile() Configuration {

	flag.Parse()

	configuration := Configuration{}

	err := gonfig.GetConf(getConfigFileName(), &configuration)
	if err != nil {
		fmt.Println(err)
		os.Exit(500)
	}

	configuration.UI = *ui
	configuration.Infra = *infra
	configuration.Env = *env

	return configuration

}

func setGlobalVariableEnv(configuration Configuration) {
	os.Setenv("ui", configuration.UI)
	os.Setenv("Infra", configuration.Infra)
	os.Setenv("PORT", configuration.Port)

	os.Setenv("ENV_name", configuration.Name)
	os.Setenv("ENV_description", configuration.Description)
	os.Setenv("ENV_version", configuration.Version)
	os.Setenv("ENV_author", configuration.Author)

	os.Setenv("ENV_company", configuration.Company)
	os.Setenv("ENV_environnement", configuration.Env)
	os.Setenv("ENV_port", configuration.Port)

	os.Setenv("ENV_dbdriver", configuration.DbDriver)
	os.Setenv("ENV_dbuser", configuration.DbUser)
	os.Setenv("ENV_dbpassword", configuration.DbPassword)

	os.Setenv("ENV_dbport", configuration.DbPort)
	os.Setenv("ENV_dbhost", configuration.DbHost)
	os.Setenv("ENV_dbname", configuration.DbName)

	os.Setenv("ENV_mockBaseUrl", configuration.MockBaseURL)
	os.Setenv("ENV_mockPort", configuration.MockPort)

}

func printConfigsOnConsole(configuration Configuration) {
	fmt.Println("=======================================")
	fmt.Println("Name 		:", configuration.Name)
	fmt.Println("Description 	:", configuration.Description)
	fmt.Println("Version 	:", configuration.Version)
	fmt.Println("---------------------------------------")
	fmt.Println("Author 		:", configuration.Author)
	fmt.Println("Company 	:", configuration.Company)
	fmt.Println("---------------------------------------")
	fmt.Println("Port 		:", configuration.Port)
	fmt.Println("Env 		:", configuration.Env)
	fmt.Println("UI 		:", configuration.UI)
	fmt.Println("Infra 		:", configuration.Infra)
	fmt.Println("=======================================")

}
