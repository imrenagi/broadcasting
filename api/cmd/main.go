package main

import (
	"github.com/rs/zerolog/log"

	"github.com/imrenagi/broadcasting/api/cmd/commands"
)

func main() {
	err := commands.NewRootCommand().Execute()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
}
