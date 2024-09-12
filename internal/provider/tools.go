package provider

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

type Locationer interface {
	IsValid() error
}

type RuleInfo struct {
	Name string `json:"name"`
	Uuid string `json:"uuid"`
}

func EncodeLocation(loc Locationer) (string, error) {
	b, err := json.Marshal(loc)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(b), nil
}

func DecodeLocation(s string, loc Locationer) error {
	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(b, loc); err != nil {
		return err
	}

	return loc.IsValid()
}

func ProviderParamDescription(desc, defaultValue, envName, jsonName string) string {
	var b strings.Builder

	b.WriteString(desc)

	if defaultValue != "" {
		b.WriteString(fmt.Sprintf(" Default: `%s`.", defaultValue))
	}

	if envName != "" {
		b.WriteString(fmt.Sprintf(" Environment variable: `%s`.", envName))
	}

	if jsonName != "" {
		b.WriteString(fmt.Sprintf(" JSON config file variable: `%s`.", jsonName))
	}

	return b.String()
}
