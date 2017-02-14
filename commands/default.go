package commands

import "bytes"

func ExecDefaultCommand() string {
	var responseBuffer bytes.Buffer

	responseBuffer.WriteString("*Command not found!!!*\n")
	responseBuffer.WriteString("Use command `help` to get a list of commands available")

	return responseBuffer.String()
}
