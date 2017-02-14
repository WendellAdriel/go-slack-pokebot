package commands

import "bytes"

func ExecHelpCommand() string {
	var responseBuffer bytes.Buffer

	responseBuffer.WriteString("*List of commands:*\n")
	responseBuffer.WriteString("`pokemon:` gets info about the given *pokemon number or name*.\n")
	responseBuffer.WriteString("```pokemon charmander\npokemon 50```")

	return responseBuffer.String()
}
