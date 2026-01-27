package internal

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/coreos/go-systemd/v22/unit"
	"github.com/thilobillerbeck/podlet2nix/pkgs/struct2nix"
)

func stringToEnv(s string) map[string]string {
	res := make(map[string]string)
	opts := strings.Split(s, " ")

	for opt := range opts {
		splitted := strings.Split(opts[opt], "=")
		res[splitted[0]] = splitted[1]
	}

	return res
}

func mapToContainer(m map[string]string) QuadletContainer {
	container := QuadletContainer{}
	container.Image = m["Image"]
	container.PublishPorts = strings.Split(m["PublishPort"], " ")
	container.Volumes = strings.Split(m["Volume"], " ")
	container.Environment = stringToEnv(m["Environment"])
	return container
}

func ParseReader(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	text := strings.Join(lines, "\n")

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
			options[opt.Name] = opt.Value
		}

		switch nameType[1] {
		case "container":
			container := mapToContainer(options)

			nix, err := struct2nix.Marshal(container, 0)
			if err != nil {
				fmt.Println("Error:", err)
			}
			fmt.Println(string(nix))
			quadlet.containers[nameType[0]] = options
		case "network":
			quadlet.networks[nameType[0]] = options
		case "pod":
			quadlet.pods[nameType[0]] = options
		}

		fmt.Println("-----")

		prettyJSON, err := json.MarshalIndent(quadlet.containers, "", "	")
		if err != nil {
			fmt.Println("Error:", err)
		}
		fmt.Println(string(prettyJSON))
	}
}
