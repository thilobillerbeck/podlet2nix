package internal

type Quadlet struct {
	containers map[string]map[string]string
	networks   map[string]map[string]string
	pods       map[string]map[string]string
}

func NewQuadlet() *Quadlet {
	return &Quadlet{
		containers: make(map[string]map[string]string),
		networks:   make(map[string]map[string]string),
		pods:       make(map[string]map[string]string),
	}
}

type QuadletContainer struct {
	Environment  map[string]string `json:"environment"`
	Command      []string          `json:"command"`
	Image        string            `json:"image"`
	PublishPorts []string          `json:"publishPorts"`
	Volumes      []string          `json:"volumes"`
}
