package model

import "strings"

type Payload struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
	Error   any    `json:"error"`
}

func IsPayloadMsg(s string) bool {
	for _, ch := range s {
		if (ch >= 'a' && ch <= 'z') || ch == ' ' {
			return false
		}
	}
	return true
}

func ToPayloadMsg(s string) string {
	return strings.ReplaceAll(strings.ToUpper(s), " ", "_")
}
