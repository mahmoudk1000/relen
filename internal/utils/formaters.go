package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

func Format(data any) (string, error) {
	var sb strings.Builder

	val := reflect.ValueOf(data)
	if val.Kind() != reflect.Struct {
		return "", fmt.Errorf("provided argument is not a struct")
	}

	sb.WriteString(fmt.Sprintf("%s:\n", val.Type().Name()))

	for i := 0; i < val.NumField(); i++ {
		key := val.Type().Field(i).Name
		value := val.Field(i)

		if !value.IsZero() {
			sb.WriteString(fmt.Sprintf("    - %s: %v\n", key, value.Interface()))
		}
	}

	return sb.String(), nil
}

func FormatJSON(data any) (string, error) {
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
