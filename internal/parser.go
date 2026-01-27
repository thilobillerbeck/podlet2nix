package internal

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/coreos/go-systemd/v22/unit"
)

func ParseReader(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	text := strings.Join(lines, "\n")

	if strings.Contains(text, "\n---\n\n") && strings.HasPrefix(text, "#") {
		quadlet := NewQuadlet()

		splitted := strings.SplitSeq(text, "\n---\n\n")

		for part := range splitted {
			lines := strings.Split(part, "\n")
			nameType := strings.Split(strings.TrimPrefix(lines[0], "# "), ".")
			body, err := unit.Deserialize(strings.NewReader(strings.Join(lines[1:], "\n")))
			if err != nil {
				panic(err)
			}

			var options map[string]string = make(map[string]string)

			for _, opt := range body {
				options[LowerFirst(opt.Name)] = opt.Value
			}

			switch nameType[1] {
			case "container":
				quadlet.containers[nameType[0]] = options
			case "network":
				quadlet.networks[nameType[0]] = options
			case "pod":
				quadlet.pods[nameType[0]] = options
			}
		}

		prettyJSON, err := json.MarshalIndent(quadlet.containers, "", "    ")
		if err != nil {
			fmt.Println("Error:", err)
		}
		fmt.Println(string(prettyJSON))
	}
}
