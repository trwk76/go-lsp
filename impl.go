package golsp

import "github.com/trwk76/go-lsp/proto"

type (
	Impl interface {
		Initialize(params proto.InitializeParams) (proto.InitializeResult, error)
		Initialized()
		Shutdown()
		Exit()
	}
)
