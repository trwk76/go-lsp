package jsonrpc

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/textproto"
	"strconv"
)

func NewPort(rd io.Reader, wr io.Writer) Port {
	br, ok := rd.(*bufio.Reader)
	if !ok {
		br = bufio.NewReader(rd)
	}

	bw, ok := wr.(*bufio.Writer)
	if !ok {
		bw = bufio.NewWriter(wr)
	}

	return Port{
		rd: rd,
		wr: wr,
		br: br,
		bw: bw,
	}
}

func (p *Port) Receive() (textproto.MIMEHeader, Message, error) {
	var (
		clen int
		msg  Message
	)

	rd := textproto.NewReader(p.br)

	hdr, err := rd.ReadMIMEHeader()
	if err != nil {
		return hdr, msg, err
	}

	if val := hdr.Get(ContentLengthHeader); val != "" {
		clen, err = strconv.Atoi(val)
		if err != nil || clen < 0 {
			return hdr, msg, ErrInvalidContentLength
		}
	} else {
		return hdr, msg, ErrInvalidContentLength
	}

	raw := make([]byte, clen)
	if _, err := io.ReadFull(p.br, raw); err != nil {
		return hdr, msg, err
	}

	if err = json.Unmarshal(raw, &msg); err != nil {
		return hdr, msg, err
	}

	return hdr, msg, nil
}

func (p *Port) Send(hdr textproto.MIMEHeader, msg Message) error {
	if hdr == nil {
		hdr = make(textproto.MIMEHeader)
	}

	raw, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	hdr.Set(ContentLengthHeader, strconv.Itoa(len(raw)))

	if _, err = io.Copy(p.bw, bytes.NewReader(raw)); err != nil {
		return err
	}

	if err = p.bw.Flush(); err != nil {
		return err
	}

	return nil
}

func (p *Port) Close() error {
	if err := p.bw.Flush(); err != nil {
		return err
	}

	if cl, ok := p.wr.(io.Closer); ok {
		if err := cl.Close(); err != nil {
			return err
		}
	}

	if cl, ok := p.rd.(io.Closer); ok {
		if err := cl.Close(); err != nil {
			return err
		}
	}

	return nil
}

type (
	Port struct {
		rd io.Reader
		wr io.Writer
		br *bufio.Reader
		bw *bufio.Writer
	}
)

var (
	ErrInvalidContentLength = errors.New("invalid content length")
)

const (
	ContentLengthHeader = "Content-Length"
)
