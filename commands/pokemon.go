package commands

import (
	"bytes"
	"fmt"

	"github.com/WendellAdriel/go-slack-pokebot/api"
)

func ExecPokemonCommand(searchValue string) string {
	var responseBuffer bytes.Buffer

	responseText := api.GetPokemonInfo(searchValue)
	responseBuffer.WriteString(fmt.Sprintf("*Info about Pokemon: %s*\n", searchValue))
	responseBuffer.WriteString("```")
	responseBuffer.WriteString(responseText)
	responseBuffer.WriteString("```")

	return responseBuffer.String()
}
