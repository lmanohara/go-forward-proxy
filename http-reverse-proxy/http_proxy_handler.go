package main

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/url"
	"strings"
)

func Handle(buff []byte, mapping proxyMappings) []byte {
	HttpRequest, error := parsedRequest(buff)
	if error != nil {
		fmt.Println(error)
	}
	path := HttpRequest.Path
	parsedUrl, error := url.Parse(path)
	if error != nil {
		fmt.Println(error)
	}

	contextPath := parsedUrl.Path

	sourceAddress := mapping[contextPath]

	fmt.Println("source address: ", sourceAddress)
	// establish connection to the http server host and port
	conn, error := net.Dial("tcp", sourceAddress)

	if error != nil {
		fmt.Println(error)
		return nil
	}

	defer conn.Close()

	forwardRequest("/", HttpRequest, conn)

	buffer := make([]byte, 4096)
	n, error := conn.Read(buffer)
	if error != nil {
		fmt.Println(error)
	}

	return buffer[:n]
}

func forwardRequest(contextPath string, HttpRequest HttpRequest, conn net.Conn) {
	var buf bytes.Buffer
	responseLine := fmt.Sprintf("GET %s %s\r\n", contextPath, HttpRequest.Version)
	buf.WriteString(responseLine)

	for k, v := range HttpRequest.Headers {
		buf.WriteString(fmt.Sprintf("%s: %s\r\n", k, v))
	}

	requestBytes := buf.Bytes()

	fmt.Println("Response as string:\n", string(requestBytes))

	// forward request headers to the http server
	conn.Write(requestBytes)
}

func parsedRequest(inputStream []byte) (HttpRequest, error) {
	data := string(inputStream)
	lines := strings.Split(data, "\r\n")
	if len(lines) == 0 {
		return HttpRequest{}, errors.New("empty request")
	}

	requestLine := lines[0]
	requestLineParts := strings.SplitAfterN(requestLine, " ", 3)
	command := strings.TrimSpace(requestLineParts[0])
	path := strings.TrimSpace(requestLineParts[1])
	version := strings.TrimSpace(requestLineParts[2])

	fmt.Println("Request line: ", requestLine)
	headers := make(map[string]string)

	for _, line := range lines[1:] {
		if line == "" {
			break
		}

		keyValue := strings.SplitN(line, ":", 2)

		if len(keyValue) == 2 {
			key := strings.TrimSpace(keyValue[0])
			value := strings.TrimSpace(keyValue[1])
			headers[key] = value
		}
	}

	for key, val := range headers {
		fmt.Printf("%s: %s\n", key, val)
	}

	req := HttpRequest{
		Method:  command,
		Path:    path,
		Version: version,
		Headers: headers,
	}

	return req, nil
}
