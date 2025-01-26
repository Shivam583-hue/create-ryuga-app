package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	//	tea "github.com/charmbracelet/bubbletea"

	huh "github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
	"github.com/charmbracelet/lipgloss"
)

var style = lipgloss.NewStyle().
	Bold(true).Foreground(lipgloss.Color("#b3ecf1")).PaddingTop(1).PaddingBottom(1)

var (
	projectName  string
	firstChoice  string
	secondChoice string
	thirdChoice  string
	fourthChoice string
)

var theme = huh.ThemeCatppuccin()

func main() {
	fmt.Println(`
   ___                           , __                                            
  / (_)                         /|/  \                                           
 |      ,_    _   __, _|_  _     |___/               __,  __,     __,    _    _  
 |     /  |  |/  /  |  |  |/     | \   |   | |   |  /  | /  |    /  |  |/ \_|/ \_
  \___/   |_/|__/\_/|_/|_/|__/   |  \_/ \_/|/ \_/|_/\_/|/\_/|_/  \_/|_/|__/ |__/ 
                                          /|          /|              /|   /|    
                                          \|          \|              \|   \|
`)
	huh.NewInput().
		Title("Name your project:").
		Value(&projectName).Run()

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Javascript or Typescript?").
				Options(
					huh.NewOption("○ Typescript", "typescript"),
					huh.NewOption("○ Javascript", "javascript"),
				).
				Value(&firstChoice),
		),
	).WithTheme(theme)
	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}
	if firstChoice == "typescript" {
		Typescript()
	} else if firstChoice == "javascript" {
		Typescript()
	} else {
		fmt.Println("invalid choice")
	}
}

//------------------------typescript functions-----------------------

func Typescript() {
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
				Value(&secondChoice),
		),
	).WithTheme(theme)
	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}
	if secondChoice == "mongodb" {
		MongoDB()
	} else if secondChoice == "mysql" {
		MongoDB()
	} else if secondChoice == "postgresql" {
		MongoDB()
	} else if secondChoice == "sqlite" {
		MongoDB()
	} else {
		fmt.Println("invalid choice")
	}
}

// ------------------------javascript functions-----------------------
func MongoDB() {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Would you like to use mongoose?").
				Options(
					huh.NewOption("○ Yes", "yes"),
					huh.NewOption("○ No", "no"),
				).
				Value(&thirdChoice),
		),
	).WithTheme(theme)
	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}
	Tailwind()
}

func Tailwind() {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Would you like to use Tailwind?").
				Options(
					huh.NewOption("○ Yes", "yes"),
					huh.NewOption("○ No", "no"),
				).
				Value(&fourthChoice),
		),
	).WithTheme(theme)
	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(fourthChoice, thirdChoice, secondChoice, firstChoice)
	if firstChoice == "javascript" && secondChoice == "mongodb" && thirdChoice == "yes" && fourthChoice == "no" {
		cmd := exec.Command("git", "clone", "https://github.com/Shivam915201/mern-js")
		// cmd.Stdout = os.Stdout
		// cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		err = os.Rename("mern-js", projectName)
		if err != nil {
			fmt.Println("Error renaming directory:", err)
			return
		}
		Spinner()
	}
	fourthChoice = ""
	thirdChoice = ""
	secondChoice = ""
	firstChoice = ""
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
