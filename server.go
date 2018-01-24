package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
)

func serve() error {
	listener, err := net.Listen("tcp", "0.0.0.0:2525")
	if err != nil {
		return err
	}

	fmt.Printf("smtp server listening on 0.0.0.0:2525\n")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
			continue
		}

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	reader := newReader(
		conn.RemoteAddr().String(),
		bufio.NewReader(conn),
		bufio.NewWriter(conn),
	)

	if err := sendEvent(reader.readEvent()); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
	}
}

func sendEvent(event *event) error {
	payload, err := json.Marshal(event)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", "http://localhost:5000/email", bytes.NewReader(payload))
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unepxected status code %d", resp.StatusCode)
	}

	return nil
}
