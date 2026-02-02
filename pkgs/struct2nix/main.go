// Package struct2nix provides functionality to marshal Go data structures
// into Nix expression format.
package struct2nix

import (
	"encoding/json"
	"errors"
	"math"
	"os"
	"strconv"
	"strings"
)

// Marshal converts a Go data structure into Nix expression format.
// It recursively processes the data structure and returns the Nix representation
// as a byte slice. The depth parameter is used for indentation control.
func Marshal(data any, depth int) ([]byte, error) {
	switch typed := data.(type) {
	case bool:
		return bool2nix(typed)
	case string:
		return string2nix(typed)
	case int:
		return int2nix(typed)
	case float64:
		return float2nix(typed)
	case []any:
		return arr2nix(typed, depth+1)
	case map[string]any:
		return map2nix(typed, depth+1)
	case interface{}:
		m, err := structToMap(typed)
		if err != nil {
			return nil, err
		}
		return map2nix(m, depth+1)
	case nil:
		return []byte("null"), nil
	default:
		os.Stderr.WriteString("unsupported type\n")
		return []byte(""), errors.New("unsupported type")
	}
}

// Helper function to generate indentation based on depth
func genIndent(depth int) string {
	return strings.Repeat("  ", depth)
}

// Helper function to convert a boolean to a Nix expression
func bool2nix(b bool) ([]byte, error) {
	if b {
		return []byte("true"), nil
	}
	return []byte("false"), nil
}

// Helper function to convert a string to a Nix expression
// This function escapes special characters in the string
func string2nix(s string) ([]byte, error) {
	escaped := strings.ReplaceAll(s, "\\", "\\\\")
	escaped = strings.ReplaceAll(escaped, "\"", "\\\"")
	escaped = strings.ReplaceAll(escaped, "${", "\\${")
	return []byte(`"` + escaped + `"`), nil
}

// Helper function to convert an integer to a Nix expression
func int2nix(i int) ([]byte, error) {
	return []byte(strconv.Itoa(i)), nil
}

// Helper function to convert a float to a Nix expression
func float2nix(f float64) ([]byte, error) {
	// Check if the float is an integer
	if f == math.Trunc(f) {
		return strconv.AppendInt(nil, int64(f), 10), nil
	}
	return strconv.AppendFloat(nil, f, 'f', -1, 64), nil
}

// Helper function to convert an array to a Nix expression
func arr2nix(arr []any, depth int) ([]byte, error) {
	var res strings.Builder
	res.WriteString("[\n")
	for _, v := range arr {
		mv, err := Marshal(v, depth)
		if err != nil {
			return nil, err
		}
		res.WriteString(genIndent(depth) + string(mv) + "\n")
	}
	res.WriteString(genIndent(depth-1) + "]")
	return []byte(res.String()), nil
}

// Helper function to convert a map to a Nix expression
func map2nix(m map[string]any, depth int) ([]byte, error) {
	var res strings.Builder
	res.WriteString("{\n")
	for k, v := range m {
		mv, err := Marshal(v, depth)
		if err != nil {
			return nil, err
		}
		res.WriteString(genIndent(depth) + k + " = " + string(mv) + ";\n")
	}
	res.WriteString(genIndent(depth-1) + "}")

	return []byte(res.String()), nil
}

// Helper function to convert a struct to a map
func structToMap(obj any) (map[string]any, error) {
	var res map[string]any

	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsonBytes, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
