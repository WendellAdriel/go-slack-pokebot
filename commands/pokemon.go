package commands

import (
	"bytes"
	"fmt"
)

func ExecPokemonCommand(searchValue string) string {
	var responseBuffer bytes.Buffer

	responseBuffer.WriteString(fmt.Sprintf("*Info about Pokemon: %s*\n", searchValue))
	responseBuffer.WriteString("INFO")

	return responseBuffer.String()
}
