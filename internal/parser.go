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
	fmt.Println(s)

	res := make(map[string]string)
	opts := strings.Split(s, " ")

	for opt := range opts {
		splitted := strings.Split(opts[opt], "=")
		if (len(splitted)) < 2 {
			res[splitted[0]] = ""
		} else {
			res[splitted[0]] = splitted[1]
		}
	}

	return res
}

func mapToContainer(m map[string]string) QuadletContainer {
	container := QuadletContainer{}
	container.Image = m["Image"]

	if m["PublishPort"] != "" {
		container.PublishPorts = strings.Split(m["PublishPort"], " ")
	}

	if m["Volume"] != "" {
		container.Volumes = strings.Split(m["Volume"], " ")
	}

	if m["Environment"] != "" {
		container.Environment = stringToEnv(m["Environment"])
	}

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
			quadlet.Containers[nameType[0]] = mapToContainer(options)
		case "network":
			quadlet.Networks[nameType[0]] = options
		case "pod":
			quadlet.Pods[nameType[0]] = options
		}

		nix, err := struct2nix.Marshal(quadlet, 0)
		if err != nil {
			fmt.Println("Error:", err)
		}
		fmt.Println(string(nix))

		fmt.Println("-----")

		prettyJSON, err := json.MarshalIndent(quadlet, "", "	")
		if err != nil {
			fmt.Println("Error:", err)
		}
		fmt.Println(string(prettyJSON))
	}
}
