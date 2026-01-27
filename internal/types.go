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
