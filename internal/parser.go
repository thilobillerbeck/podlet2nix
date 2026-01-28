package internal

import (
	"bufio"
	"io"
	"os"
	"strings"

	"github.com/coreos/go-systemd/v22/unit"
	"github.com/thilobillerbeck/podlet2nix/pkgs/struct2nix"
)

func stringToEnv(s string) map[string]string {
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

func splitOrNil(s string, sep string) []string {
	if s == "" {
		return nil
	}
	return strings.Split(s, sep)
}

func mapToContainer(m map[string]string) ContainerOptions {
	container := ContainerOptions{}
	container.ContainerConfig = &ContainerConfig{
		AddCapability: splitOrNil(m["AddCapability"], " "),
		Image:         m["Image"],
		PublishPort:   splitOrNil(m["PublishPort"], " "),
		Volume:        splitOrNil(m["Volume"], " "),
		Environment:   stringToEnv(m["Environment"]),
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

	quadlet := Quadlet{
		Builds:     make(map[string]BuildOptions),
		Containers: make(map[string]ContainerOptions),
		Images:     make(map[string]ImageOptions),
		Networks:   make(map[string]NetworkOptions),
		Pods:       make(map[string]PodOptions),
		Volumes:    make(map[string]VolumeOptions),
	}

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
		}
	}

	nix, err := struct2nix.Marshal(quadlet, 0)
	if err != nil {
		os.Stderr.WriteString("Error: " + err.Error() + "\n")
		os.Exit(1)
	}

	os.Stdout.WriteString(string(nix))
}
