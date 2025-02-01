package cli

import (
	"ez/internal/command"
	"log"

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

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Templates").
				Options(
					huh.NewOption("Ah", "ah"),
					huh.NewOption("Ah-uth", "ahuth"),
					huh.NewOption("Ha", "ha"),
				).
				Value(&cmd),
		),
	)

	if err := form.Run(); err != nil {
		log.Fatal(err)
	}

	var err error
	switch cmd {
	case "ah":
		err = command.Ah()
	case "ah-uth":
		err = command.Ahuth()
	case "ha":
		err = command.Ha()
	default:
		log.Fatalf("unknown command %s", cmd)
	}
	if err != nil {
		log.Fatal(err)
	}
}
