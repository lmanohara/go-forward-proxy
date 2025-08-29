package main

type HttpRequest struct {
	Method  string
	Path    string
	Version string
	Headers map[string]string
}
