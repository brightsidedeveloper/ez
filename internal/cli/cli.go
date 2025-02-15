package cli

import (
	"ez/internal/command"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/briandowns/spinner"
	"github.com/charmbracelet/huh"
)

func Run() {
	var cmd string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("What's good?").
				Options(
					huh.NewOption("Templates", "templates"),
				).
				Value(&cmd),
		),
	)

	if err := form.Run(); err != nil {
		log.Fatal(err)
	}

	switch cmd {
	case "templates":
		Templates()
	default:
		log.Fatalf("unknown command %s", cmd)
	}
}

func Templates() {
	var cmd string
	var name string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Templates").
				Options(
					huh.NewOption("Ah", "ah"),
					huh.NewOption("Ah-uth", "ahuth"),
					huh.NewOption("Ha", "ha"),
					huh.NewOption("Haai", "haai"),
				).
				Value(&cmd),
		),
		huh.NewGroup(
			huh.NewInput().
				Title("What’s your name?").
				Value(&name),
		),
	)

	if err := form.Run(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Hey, %s! Give me a second to cook...\n\n", name)
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Start()

	var err error
	switch cmd {
	case "ah":
		err = command.Ah(name)
	case "ahuth":
		err = command.Ahuth(name)
	case "ha":
		err = command.Ha(name)
	case "haai":
		err = command.Haai(name)
	default:
		log.Fatalf("unknown command %s", cmd)
	}
	if err != nil {
		log.Fatal(err)
	}
	s.Stop()
	fmt.Println("\n\nHave Fun!")
	baseDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	openCmd := exec.Command("code", ".")
	openCmd.Dir = filepath.Join(baseDir, name)
	if err := openCmd.Run(); err != nil {
		log.Fatal(err)
	}
}
