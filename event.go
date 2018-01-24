package main

import (
	"bufio"
	"fmt"
	"strings"
)

type (
	reader struct {
		addr   string
		reader *bufio.Reader
		writer *bufio.Writer
	}

	event struct {
		Recipient string `json:"recipient"`
		Subject   string `json:"subject"`
		Body      string `json:"body"`
	}
)

func newReader(addr string, r *bufio.Reader, w *bufio.Writer) *reader {
	return &reader{addr: addr, reader: r, writer: w}
}

func (r *reader) readEvent() *event {
	r.write("220 Welcome")

	r.read()
	r.write("250 Sender OK")

	recipientLine, _ := r.read()
	r.write("250 Recipient OK")

	r.read()
	r.write("250 OK")

	r.read()
	r.write("354 Data")

	payload := r.readPayload()
	r.write("250 OK")

	// Split into headers and body
	parts := strings.SplitN(payload, "\r\n\r\n", 2)

	r.read()
	r.write("221 Bye")

	return &event{
		Recipient: extractRecipient(recipientLine),
		Subject:   extractSubject(parts[0]),
		Body:      strings.TrimSpace(parts[1]),
	}
}

//
// Read/Write Helpers

func (r *reader) read() (string, bool) {
	reply, err := r.reader.ReadString('\n')
	fmt.Printf("[%s] >> %s", r.addr, reply)
	return reply, err == nil
}

func (r *reader) write(line string) {
	fmt.Printf("[%s] << %s\n", r.addr, line)
	r.writer.WriteString(line + "\r\n")
	r.writer.Flush()
}

func (r *reader) readPayload() string {
	payload := ""

	for {
		line, ok := r.read()

		if !ok || line == ".\r\n" {
			return payload
		}

		payload += line
	}
}

//
// Parsing Helpers

func extractRecipient(line string) string {
	return strings.Trim(strings.TrimSpace(strings.SplitN(line, ":", 2)[1]), "<>")
}

func extractSubject(headers string) string {
	for _, line := range strings.Split(headers, "\r\n") {
		if strings.HasPrefix(line, "Subject:") {
			return strings.TrimSpace(line[8:])
		}
	}

	return ""
}
