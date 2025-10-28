package jsonrpc

import "encoding/json"

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

const (
	ErrCodeParseError           int = -32700
	ErrCodeInvalidRequest       int = -32600
	ErrCodeMethodNotFound       int = -32601
	ErrCodeInvalidParams        int = -32602
	ErrCodeInternalError        int = -32603
	ErrCodeServerNotInitialized int = -32002
	ErrCodeUnknownErrorCode     int = -32001
	ErrCodeRequestFailed        int = -32803
	ErrCodeServerCancelled      int = -32802
	ErrCodeContentModified      int = -32801
	ErrCodeRequestCancelled     int = -32800
)
