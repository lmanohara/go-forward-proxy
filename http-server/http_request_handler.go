package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

func Handle(iputStream []byte) []byte {

	HttpRequest, error := parsedRequest(iputStream)

	if error != nil {
		fmt.Println(error)
	}

	var buf bytes.Buffer
	validatedPath := validatePath(HttpRequest.Path)
	if !validatedPath {
		writeResponseLine(&buf, 404)
		writeResponseHeaders(&buf, 0)
	} else if HttpRequest.Method == "POST" {
		writeResponseLine(&buf, 403)
		writeResponseHeaders(&buf, 0)
	} else if HttpRequest.Method != "GET" {
		writeResponseLine(&buf, 405)
		writeResponseHeaders(&buf, 0)
	}

	if buf.Len() > 0 {
		responseBytes := buf.Bytes()
		fmt.Println("Response as string:\n", string(responseBytes))
		return responseBytes
	}

	return handleGet(&buf)
}

func handleGet(buf *bytes.Buffer) []byte {

	body, _ := os.ReadFile("index.html")

	writeResponseLine(buf, 200)

	writeResponseHeaders(buf, len(body))

	buf.WriteString(string(body))

	responseBytes := buf.Bytes()

	fmt.Println("Response as string:\n", string(responseBytes))

	return responseBytes
}

func writeResponseLine(buff *bytes.Buffer, statusCode int) {
	statusLine := fmt.Sprintf("HTTP/1.1 %d %s\r\n", statusCode, getStatusText(statusCode))
	buff.WriteString(statusLine)
}

func writeResponseHeaders(buff *bytes.Buffer, contentLength int) {

	headers := map[string]string{
		"Content-Type":   "text/html",
		"Content-Length": strconv.Itoa(contentLength),
		"Connection":     "close",
	}

	for k, v := range headers {
		buff.WriteString(fmt.Sprintf("%s: %s\r\n", k, v))
	}

	buff.WriteString("\r\n")
}

func getStatusText(statusCode int) string {
	switch statusCode {
	case 200:
		return "OK"
	case 403:
		return "Forbidden"
	case 404:
		return "Not Found"
	case 405:
		return "Method Not Allowed"
	default:
		return ""
	}
}

func validatePath(contextPath string) bool {
	extractedPath := path.Dir(contextPath)
	fmt.Println("request path: ", contextPath)
	info, error := os.Stat(extractedPath)

	if os.IsNotExist(error) {
		return false
	}

	if error != nil {
		return false
	}

	if info.IsDir() {
		filepath.Join(contextPath, "index.html")

		return true
	}

	return false
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
