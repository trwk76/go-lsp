package jsonrpc_test

import (
	"strings"
	"testing"

	"github.com/trwk76/go-lsp/jsonrpc"
)

func TestPort(t *testing.T) {
	rd := strings.NewReader("Content-Length: 52\r\n\r\n{\"jsonrpc\":\"2.0\",\"id\":1,\"method\":\"test\",\"params\":{}}")
	wr := strings.Builder{}
	prt := jsonrpc.NewPort(rd, &wr)

	hdr, msg, err := prt.Receive()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	if err := prt.Send(hdr, msg); err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	t.Log(wr.String())
}
