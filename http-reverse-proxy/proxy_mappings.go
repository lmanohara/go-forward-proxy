package main

import (
	"fmt"
	"strings"
)

type proxyMappings map[string]string

func (m *proxyMappings) String() string {
	parts := []string{}

	for k, v := range *m {
		parts = append(parts, fmt.Sprintf("%s=%s", k, v))
	}

	return strings.Join(parts, ", ")
}

func (m *proxyMappings) Set(value string) error {

	entries := strings.Split(value, ",")

	for _, entry := range entries {
		parts := strings.SplitN(strings.TrimSpace(entry), "=", 2)
		if len(parts) != 2 {
			return fmt.Errorf("Invalid mappings, expected path=target but got %q", entry)
		}
		(*m)[parts[0]] = parts[1]
	}

	return nil
}
