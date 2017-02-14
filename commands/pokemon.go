package commands

import "fmt"

func ExecPokemonCommand(searchValue string) string {
	response := fmt.Sprintf("*Info about Pokemon: %s*\nINFO", searchValue)
	return response
}
