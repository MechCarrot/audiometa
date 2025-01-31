package models

import (
	"bytes"
	"encoding/json"
)

type Audio struct {
	Id       string
	Path     string
	Metadata Metadata
	Status   string
	Error    []error
}

func (audio *Audio) JSON() (string, error) {
	audioJSON, err := json.Marshal(audio)
	if err != nil {
		return "", err
	}
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(audioJSON), "", "    "); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}
