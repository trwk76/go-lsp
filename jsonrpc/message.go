package jsonrpc

import (
	"encoding/json"
	"fmt"
)

func ToError(err error) Error {
	if res, ok := err.(Error); ok {
		return res
	}

	return ErrUnknown
}

type (
	MessageID = json.RawMessage

	Message struct {
		JSONRPC string          `json:"jsonrpc"`
		ID      MessageID       `json:"id,omitempty"`
		Method  string          `json:"method,omitempty"`
		Params  json.RawMessage `json:"params,omitempty"`
		Result  json.RawMessage `json:"result,omitempty"`
		Error   *Error          `json:"error,omitempty"`
	}

	Error struct {
		Code    int             `json:"code"`
		Message string          `json:"message"`
		Data    json.RawMessage `json:"data,omitempty"`
	}
)

func(e Error) Error() string {
	return fmt.Sprintf("jsonrpc error %d: %s", e.Code, e.Message)
}

var (
	ErrParseError       Error = err(-32700, "Parse error")
	ErrInvalidRequest   Error = err(-32600, "Invalid request")
	ErrMethodNotFound   Error = err(-32601, "Method not found")
	ErrInvalidParams    Error = err(-32602, "Invalid parameters")
	ErrInternalError    Error = err(-32603, "Internal error")
	ErrNotInit          Error = err(-32002, "Server not initialized")
	ErrUnknown          Error = err(-32001, "Unknown error")
	ErrRequestFailed    Error = err(-32803, "Request failed")
	ErrServerCancelled  Error = err(-32802, "Server cancelled")
	ErrContentModified  Error = err(-32801, "Content modified")
	ErrRequestCancelled Error = err(-32800, "Request cancelled")
)

func err(code int, msg string) Error {
	return Error{
		Code:    code,
		Message: msg,
	}
}

var (
	_ error = Error{}
)
