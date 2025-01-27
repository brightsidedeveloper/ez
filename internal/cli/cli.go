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
				Title("Templates").
				Options(
					huh.NewOption("Ah", "proto"),
				).
				Value(&cmd),
		),
	)

	if err := form.Run(); err != nil {
		log.Fatal(err)
	}

	switch cmd {
	case "proto":
		err := command.Proto()
		if err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatalf("unknown command %s", cmd)
	}
}
