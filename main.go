package main

import (
	"t3/utils"
)

var (
	projectName string
	language    string
	database    string
	orm         string
	styling     string
)

func main() {
	utils.ShowWelcomeMessage()

	utils.NameProject(&projectName)

	utils.SelectLanguage(&language)

	utils.SelectDatabase(&database, &orm, &styling, &language, &projectName)
}
