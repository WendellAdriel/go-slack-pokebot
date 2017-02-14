package commands

import "fmt"

func ExecDefaultCommand() string {
	response := fmt.Sprintf("*Command not found!!!*\nUse command `help` to get a list of commands available")
	return response
}
