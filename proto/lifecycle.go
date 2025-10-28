package proto

import "encoding/json"

type (
	ClientCapabilities struct {

	}

	ClientInfo struct {
		Name    string `json:"name"`
		Version string `json:"version,omitempty"`
	}

	InitializeParams struct {
		ProcessID        *int32             `json:"processId"`
		ClientInfo       *ClientInfo        `json:"clientInfo,omitempty"`
		Locale           string             `json:"locale,omitempty"`
		RootPath         *string            `json:"rootPath"`
		RootURI          *DocumentURI       `json:"rootUri"`
		InitOptions      json.RawMessage    `json:"initializationOptions,omitempty"`
		Capabilities     ClientCapabilities `json:"capabilities"`
		Trace            TraceVal           `json:"trace,omitempty"`
		WorkspaceFolders []WorkspaceFolder  `json:"workspaceFolders,omitempty"`
	}

	InitializeResult struct {
		Capabilities ServerCapabilities `json:"capabilities"`
		ServerInfo   *ServerInfo        `json:"serverInfo,omitempty"`
	}

	ServerCapabilities struct {
		PositionEncoding PositionEncoding `json:"positionEncoding"`
		TextDocumentSync TextDocumentSync `json:"textDocumentSync"`
	}

	ServerInfo struct {
		Name    string `json:"name"`
		Version string `json:"version,omitempty"`
	}

	TraceVal string
)

const (
	TraceNone     TraceVal = ""
	TraceOff      TraceVal = "off"
	TraceMessages TraceVal = "messages"
)
