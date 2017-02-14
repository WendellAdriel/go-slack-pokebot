package commands

import "fmt"

func ExecHelpCommand() string {
	response := fmt.Sprintf("*List of commands:*\n`pokemon:` gets info about the given *pokemon number or name*.\n```pokemon charmander\npokemon 50```")
	return response
}
