package internal

type Quadlet struct {
	Containers map[string]QuadletContainer  `json:"containers"`
	Networks   map[string]map[string]string `json:"networks"`
	Pods       map[string]map[string]string `json:"pods"`
}

func NewQuadlet() *Quadlet {
	return &Quadlet{
		Containers: make(map[string]QuadletContainer),
		Networks:   make(map[string]map[string]string),
		Pods:       make(map[string]map[string]string),
	}
}

type QuadletContainer struct {
	Environment  map[string]string `json:"environment"`
	Command      []string          `json:"command"`
	Image        string            `json:"image"`
	PublishPorts []string          `json:"publishPorts"`
	Volumes      []string          `json:"volumes"`
}
