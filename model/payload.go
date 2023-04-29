package model

import "strings"

type Payload struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
	Error   any    `json:"error"`
}

func ToPayloadMsg(s string) string {
	return strings.ReplaceAll(strings.ToUpper(s), " ", "_")
}
