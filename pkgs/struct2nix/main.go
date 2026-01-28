package struct2nix

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
)

func Marshal(data any, depth int) ([]byte, error) {
	switch typed := data.(type) {
	case bool:
		return bool2nix(typed)
	case string:
		return string2nix(typed)
	case int:
		return int2nix(typed)
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

func Unmarshal(data []byte, v any) error {
	return nil
}

func genIndent(depth int) string {
	return strings.Repeat("\t", depth)
}

func bool2nix(b bool) ([]byte, error) {
	if b {
		return []byte("true"), nil
	}
	return []byte("false"), nil
}

func string2nix(s string) ([]byte, error) {
	return []byte(`"` + s + `"`), nil
}

func int2nix(i int) ([]byte, error) {
	return []byte(string(i)), nil
}

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

func map2nix(m map[string]any, depth int) ([]byte, error) {
	var res strings.Builder
	res.WriteString("{\n")
	for k, v := range m {
		// TODO: make this configurable
		if v == nil {
			continue
		}
		mv, err := Marshal(v, depth)
		if err != nil {
			return nil, err
		}
		res.WriteString(genIndent(depth) + k + " = " + string(mv) + ";\n")
	}
	res.WriteString(genIndent(depth-1) + "}")

	return []byte(res.String()), nil
}

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
