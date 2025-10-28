package jsonrpc

import "os"

func NewStdioPort() Port {
	return NewPort(
		os.Stdout,
		os.Stdin,
	)
}
