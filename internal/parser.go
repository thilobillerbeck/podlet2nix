package internal

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/coreos/go-systemd/v22/unit"
	"github.com/thilobillerbeck/podlet2nix/pkgs/struct2nix"
)

func stringToEnv(s string) map[string]string {
	if s == "" {
		return nil
	}
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

func handleInterface(s string, n string) interface{} {
	// TODO: implement rest of interfaces
	if s == "" {
		return nil
	}

	switch n {
	case "Entrypoint":
		if strings.HasPrefix(s, "[") {
			s = strings.ReplaceAll(s, "\"", "")
			s = s[1 : len(s)-1]
			return strings.Split(s, ",")
		}
		return s[1 : len(s)-1]
	}

	return s
}

func splitOrNil(s string, sep string) []string {
	if s == "" {
		return nil
	}
	return strings.Split(s, sep)
}

func FillStruct(data map[string]string, result any) {
	v := reflect.ValueOf(result).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		value := data[t.Field(i).Name]
		vlueType := t.Field(i).Type.Kind()

		switch vlueType {
		case reflect.Int:
			// TODO: handle error
			_, err := strconv.Atoi(value)
			if err != nil {
				continue
			}
		case reflect.Bool:
			v.Field(i).SetBool(value == "true")
		case reflect.Slice:
			v.Field(i).Set(reflect.ValueOf(splitOrNil(value, ",")))
		case reflect.String:
			v.Field(i).SetString(value)
		case reflect.Map:
			v.Field(i).Set(reflect.ValueOf(stringToEnv(value)))
		case reflect.Interface:
			res := handleInterface(value, t.Field(i).Name)
			if res != nil {
				v.Field(i).Set(reflect.ValueOf(res))
			}
		default:
			fmt.Println("unsupported type: ", t.Field(i).Type.Kind())
		}
	}
}

func mapToContainer(m map[string]string) ContainerOptions {
	container := ContainerOptions{}

	var containerConfig *ContainerConfig = &ContainerConfig{}
	FillStruct(m, containerConfig)
	container.ContainerConfig = containerConfig

	return container
}

func mapToBuild(m map[string]string) BuildOptions {
	build := BuildOptions{}

	var buildConfig *BuildConfig = &BuildConfig{}
	FillStruct(m, buildConfig)
	build.BuildConfig = buildConfig

	return build
}

func mapToImage(m map[string]string) ImageOptions {
	image := ImageOptions{}

	var imageConfig *ImageConfig = &ImageConfig{}
	FillStruct(m, imageConfig)
	image.ImageConfig = imageConfig

	return image
}

func mapToNetwork(m map[string]string) NetworkOptions {
	network := NetworkOptions{}

	var networkConfig *NetworkConfig = &NetworkConfig{}
	FillStruct(m, networkConfig)
	network.NetworkConfig = networkConfig

	return network
}

func mapToVolume(m map[string]string) VolumeOptions {
	volume := VolumeOptions{}

	var volumeConfig *VolumeConfig = &VolumeConfig{}
	FillStruct(m, volumeConfig)
	volume.VolumeConfig = volumeConfig

	return volume
}

func mapToPod(m map[string]string) PodOptions {
	pod := PodOptions{}

	var podConfig *PodConfig = &PodConfig{}
	FillStruct(m, podConfig)
	pod.PodConfig = podConfig

	return pod
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
		case "network":
			quadlet.Networks[nameType[0]] = mapToNetwork(options)
		case "pod":
			quadlet.Pods[nameType[0]] = mapToPod(options)
		case "volume":
			quadlet.Volumes[nameType[0]] = mapToVolume(options)
		case "image":
			quadlet.Images[nameType[0]] = mapToImage(options)
		case "build":
			quadlet.Builds[nameType[0]] = mapToBuild(options)
		}
	}

	nix, err := struct2nix.Marshal(quadlet, 0)
	if err != nil {
		os.Stderr.WriteString("Error: " + err.Error() + "\n")
		os.Exit(1)
	}

	os.Stdout.WriteString(string(nix) + "\n")
}
