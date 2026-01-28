package internal

type Quadlet struct {
	Enable     bool                    `json:"enable,omitempty"`
	AutoEscape bool                    `json:"autoEscape,omitempty"`
	AutoUpdate AutoUpdateConfig        `json:"autoUpdate,omitempty"`
	Builds     map[string]BuildOptions `json:"builds,omitempty"`
}

type AutoUpdateConfig struct {
	Enable   bool   `json:"enable"`
	Calendar string `json:"calendar"`
}

type BuildOptions struct {
	AutoStart      bool                `json:"autoStart,omitempty"`
	BuildConfig    BuildConfig         `json:"buildConfig,omitempty"`
	QuadletConfig  BuildQuadletConfig  `json:"quadletConfig,omitempty"`
	RawConfig      string              `json:"rawConfig,omitempty"`
	Ref            string              `json:"ref,omitempty"`
	RootlessConfig BuildRootlessConfig `json:"rootlessConfig,omitempty"`
	ServiceCOnfig  map[string]string   `json:"serviceConfig,omitempty"`
	UnitConfig     map[string]string   `json:"unitConfig,omitempty"`
}

type BuildRootlessConfig struct {
	Uid int `json:"uid,omitempty"`
}

type BuildQuadletConfig struct {
	DefaultDependencies bool `json:"defaultDependencies,omitempty"`
}

type BuildConfig struct {
	AddGroups            []string          `json:"addGroups,omitempty"`
	Annotations          map[string]string `json:"annotations,omitempty"`
	Arch                 string            `json:"arch,omitempty"`
	AuthFile             string            `json:"authFile,omitempty"`
	BuildArgs            map[string]string `json:"buildArgs,omitempty"`
	DNS                  []string          `json:"dns,omitempty"`
	DNSOption            []string          `json:"dnsOption,omitempty"`
	DNSSeach             []string          `json:"dnsSeach,omitempty"`
	Environments         map[string]string `json:"environments,omitempty"`
	File                 string            `json:"file,omitempty"`
	ForceRM              bool              `json:"forceRm,omitempty"`
	GlobalArgs           []string          `json:"globalArgs,omitempty"`
	IgnoreFile           string            `json:"ignoreFileand,omitempty"`
	Labels               map[string]string `json:"labels,omitempty"`
	ContainersConfModule []string          `json:"modules,omitempty"`
	Network              []string          `json:"networks,omitempty"`
	PodmanArgs           []string          `json:"podmanArgs,omitempty"`
	Pull                 string            `json:"pull,omitempty"`
	Retrty               int               `json:"retry,omitempty"`
	RetryDelay           string            `json:"retryDelay,omitempty"`
	Secret               []string          `json:"secrets,omitempty"`
	Tag                  string            `json:"tag,omitempty"`
	Target               string            `json:"target,omitempty"`
	TLSVerify            bool              `json:"tlsVerify,omitempty"`
	Variant              string            `json:"variant,omitempty"`
	Volumes              []string          `json:"volumes,omitempty"`
	WorkDir              string            `json:"workDir,omitempty"`
}
