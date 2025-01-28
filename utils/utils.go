package utils

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
	"github.com/charmbracelet/lipgloss"
)

var (
	theme = huh.ThemeCatppuccin()
	style = lipgloss.NewStyle().
		Bold(true).Foreground(lipgloss.Color("#b3ecf1")).PaddingTop(1).PaddingBottom(1)
)

func ShowWelcomeMessage() {
	fmt.Println(`
   ___                           , __                                            
  / (_)                         /|/  \                                           
 |      ,_    _   __, _|_  _     |___/               __,  __,     __,    _    _  
 |     /  |  |/  /  |  |  |/     | \   |   | |   |  /  | /  |    /  |  |/ \_|/ \_
  \___/   |_/|__/\_/|_/|_/|__/   |  \_/ \_/|/ \_/|_/\_/|/\_/|_/  \_/|_/|__/ |__/ 
                                          /|          /|              /|   /|    
                                          \|          \|              \|   \|
`)
}

func NameProject(projectName *string) {
	huh.NewInput().
		Title("Name your project:").
		Value(projectName).Run()
}

func SelectLanguage(language *string) {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Javascript or Typescript?").
				Options(
					huh.NewOption("○ Typescript", "typescript"),
					huh.NewOption("○ Javascript", "javascript"),
				).
				Value(language),
		),
	).WithTheme(theme)
	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func SelectDatabase(database *string, orm *string, styling *string, language *string, projectName *string) {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Choose your database:").
				Options(
					huh.NewOption("○ MongoDB", "mongodb"),
					huh.NewOption("○ MySQL", "mysql"),
					huh.NewOption("○ PostgreSQL", "postgresql"),
					huh.NewOption("○ SQLite", "sqlite"),
				).
				Value(database),
		),
	).WithTheme(theme)
	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}
	if *database == "mongodb" {
		MongoDB(orm, styling, language, database, projectName)
	} else if *database == "mysql" {
		Mysql(orm, styling, language, database, projectName)
	} else if *database == "postgresql" {
		Postgresql(orm, styling, language, database, projectName)
	} else if *database == "sqlite" {
		Sqlite(orm, styling, language, database, projectName)
	} else {
		fmt.Println("invalid choice")
	}
}

func MongoDB(orm *string, styling *string, language *string, database *string, projectName *string) {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Would you like to use mongoose?").
				Options(
					huh.NewOption("○ Yes", "yes"),
					huh.NewOption("○ No", "no"),
				).
				Value(orm),
		),
	).WithTheme(theme)
	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}
	Tailwind(styling, orm, database, language, projectName)
}

func Mysql(orm *string, styling *string, language *string, database *string, projectName *string) {}

func Postgresql(orm *string, styling *string, language *string, database *string, projectName *string) {
}

func Sqlite(orm *string, styling *string, language *string, database *string, projectName *string) {}

func Tailwind(styling *string, orm *string, database *string, language *string, projectName *string) {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Would you like to use Tailwind?").
				Options(
					huh.NewOption("○ Yes", "yes"),
					huh.NewOption("○ No", "no"),
				).
				Value(styling),
		),
	).WithTheme(theme)
	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	if *language == "javascript" && *database == "mongodb" && *orm == "yes" && *styling == "no" {
		// clone the repo
		cmd := exec.Command("git", "clone", "https://github.com/Shivam915201/mern-js")
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		// rename the directory
		err = os.Rename("mern-js", *projectName)
		if err != nil {
			fmt.Println("Error renaming directory:", err)
			return
		}

		// remove the .git folder
		cmd2 := exec.Command("sh", "-c", fmt.Sprintf("cd %s && git remote remove origin && rm -rf .git", *projectName))
		err = cmd2.Run()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		Spinner()
	}
	*styling = ""
	*orm = ""
	*database = ""
	*language = ""
}

func Spinner() {
	action := func() {
		time.Sleep(2 * time.Second)
	}
	if err := spinner.New().Title("Intialising Project").Action(action).Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(style.Render("Project Intialised! "))
	fmt.Println(style.Render("Run the backend server with cd backend && npm run server"))
	fmt.Println(style.Render("Run the frontend server with cd frontend && npm run dev"))
	fmt.Println(style.Render("Deploy on render or railway or vercel"))
}
