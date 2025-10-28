package proto

type (
	TextDocumentSync int
)

const (
	TextDocumentSyncNone        TextDocumentSync = 0
	TextDocumentSyncFull        TextDocumentSync = 1
	TextDocumentSyncIncremental TextDocumentSync = 2
)
