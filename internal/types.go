package internal

type Quadlet struct {
	Enable     bool                        `json:"enable,omitempty"`
	AutoEscape bool                        `json:"autoEscape,omitempty"`
	AutoUpdate AutoUpdateConfig            `json:"autoUpdate,omitempty"`
	Builds     map[string]BuildOptions     `json:"builds,omitempty"`
	Containers map[string]ContainerOptions `json:"containers,omitempty"`
	Images     map[string]ImageOptions     `json:"images,omitempty"`
	Networks   map[string]NetworkOptions   `json:"networks,omitempty"`
}

type AutoUpdateConfig struct {
	Enable   bool   `json:"enable"`
	Calendar string `json:"calendar"`
}

type BuildOptions struct {
	AutoStart      bool              `json:"autoStart,omitempty"`
	BuildConfig    BuildConfig       `json:"buildConfig,omitempty"`
	QuadletConfig  QuadletConfig     `json:"quadletConfig,omitempty"`
	RawConfig      string            `json:"rawConfig,omitempty"`
	Ref            string            `json:"ref,omitempty"`
	RootlessConfig RootlessConfig    `json:"rootlessConfig,omitempty"`
	ServiceCOnfig  map[string]string `json:"serviceConfig,omitempty"`
	UnitConfig     map[string]string `json:"unitConfig,omitempty"`
}

type RootlessConfig struct {
	Uid int `json:"uid,omitempty"`
}

type QuadletConfig struct {
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

type ContainerOptions struct {
	AutoStart       bool              `json:"autoStart,omitempty"`
	ContainerConfig ContainerConfig   `json:"containerConfig,omitempty"`
	QuadletConfig   QuadletConfig     `json:"quadletConfig,omitempty"`
	RawConfig       string            `json:"rawConfig,omitempty"`
	Ref             string            `json:"ref,omitempty"`
	RootlessConfig  RootlessConfig    `json:"rootlessConfig,omitempty"`
	ServiceConfig   map[string]string `json:"serviceConfig,omitempty"`
	UnitConfig      map[string]string `json:"unitConfig,omitempty"`
}

type ImageOptions struct {
	AutoStart      bool              `json:"autoStart,omitempty"`
	ImageConfig    ImageConfig       `json:"imageConfig,omitempty"`
	QuadletConfig  QuadletConfig     `json:"quadletConfig,omitempty"`
	RawConfig      string            `json:"rawConfig,omitempty"`
	Ref            string            `json:"ref,omitempty"`
	RootlessConfig RootlessConfig    `json:"rootlessConfig,omitempty"`
	ServiceConfig  map[string]string `json:"serviceConfig,omitempty"`
	UnitConfig     map[string]string `json:"unitConfig,omitempty"`
}

type ImageConfig struct {
	AllTags              bool     `json:"allTags,omitempty"`
	Arch                 string   `json:"arch,omitempty"`
	AuthFile             string   `json:"authFile,omitempty"`
	CertDir              string   `json:"certDir,omitempty"`
	Creds                string   `json:"creds,omitempty"`
	DecryptionKey        string   `json:"decryptionKey,omitempty"`
	GlobalArgs           []string `json:"globalArgs,omitempty"`
	Image                string   `json:"image,omitempty"`
	ContainersConfModule []string `json:"modules,omitempty"`
	OS                   string   `json:"os,omitempty"`
	PodmanArgs           []string `json:"podmanArgs,omitempty"`
	Policy               string   `json:"policy,omitempty"`
	Retry                int      `json:"retry,omitempty"`
	RetryDelay           string   `json:"retryDelay,omitempty"`
	ImageTag             string   `json:"tag,omitempty"`
	TLSVerify            bool     `json:"tlsVerify,omitempty"`
	Variant              string   `json:"variant,omitempty"`
}

type NetworkOptions struct {
	AutoStart      bool              `json:"autoStart,omitempty"`
	NetworkConfig  NetworkConfig     `json:"networkConfig,omitempty"`
	QuadletConfig  QuadletConfig     `json:"quadletConfig,omitempty"`
	RawConfig      string            `json:"rawConfig,omitempty"`
	Ref            string            `json:"ref,omitempty"`
	RootlessConfig RootlessConfig    `json:"rootlessConfig,omitempty"`
	ServiceConfig  map[string]string `json:"serviceConfig,omitempty"`
	UnitConfig     map[string]string `json:"unitConfig,omitempty"`
}

type NetworkConfig struct {
	DisableDNS           bool              `json:"disableDns,omitempty"`
	DNS                  []string          `json:"dns,omitempty"`
	Driver               string            `json:"driver,omitempty"`
	Gateway              []string          `json:"gateways,omitempty"`
	GlobalArgs           []string          `json:"globalArgs,omitempty"`
	InterfaceName        string            `json:"interfaceName,omitempty"`
	Internal             bool              `json:"internal,omitempty"`
	IPRange              []string          `json:"ipRanges,omitempty"`
	IPAMDriver           string            `json:"ipamDriver,omitempty"`
	IPv6                 bool              `json:"ipv6,omitempty"`
	Label                map[string]string `json:"labels,omitempty"`
	ContainersConfModule []string          `json:"modules,omitempty"`
	NetworkName          string            `json:"name,omitempty"`
	NetworkDeleteOnStop  bool              `json:"networkDeleteOnStop,omitempty"`
	Options              map[string]string `json:"options,omitempty"`
	PodmanArgs           []string          `json:"podmanArgs,omitempty"`
	Subnet               []string          `json:"subnets,omitempty"`
}

type ContainerConfig struct {
	AddCapability         []string          `json:"addCapabilities,omitempty"`
	GroupAdd              []string          `json:"addGroups,omitempty"`
	AddHost               []string          `json:"addHosts,omitempty"`
	Annotation            map[string]string `json:"annotations,omitempty"`
	AppArmor              string            `json:"appArmor,omitempty"`
	AutoUpdate            string            `json:"autoUpdate,omitempty"`
	CgroupsMode           string            `json:"cgroupsMode,omitempty"`
	AddDevice             []string          `json:"devices,omitempty"`
	DNS                   []string          `json:"dns,omitempty"`
	DNSOption             []string          `json:"dnsOption,omitempty"`
	DNSSearch             []string          `json:"dnsSearch,omitempty"`
	DropCapability        []string          `json:"dropCapabilities,omitempty"`
	Entrypoint            interface{}       `json:"entrypoint,omitempty"` // string or []string
	EnvironmentFile       []string          `json:"environmentFiles,omitempty"`
	EnvironmentHost       bool              `json:"environmentHost,omitempty"`
	Environment           map[string]string `json:"environments,omitempty"`
	Exec                  interface{}       `json:"exec,omitempty"` // string or []string
	ExposePort            []string          `json:"exposePorts,omitempty"`
	GIDMap                []string          `json:"gidMaps,omitempty"`
	GlobalArgs            []string          `json:"globalArgs,omitempty"`
	Group                 string            `json:"group,omitempty"`
	HealthCmd             string            `json:"healthCmd,omitempty"`
	HealthInterval        string            `json:"healthInterval,omitempty"`
	HealthLogDestination  string            `json:"healthLogDestination,omitempty"`
	HealthMaxLogCount     int               `json:"healthMaxLogCount,omitempty"`
	HealthMaxLogSize      int               `json:"healthMaxLogSize,omitempty"`
	HealthOnFailure       string            `json:"healthOnFailure,omitempty"`
	HealthRetries         int               `json:"healthRetries,omitempty"`
	HealthStartPeriod     string            `json:"healthStartPeriod,omitempty"`
	HealthStartupCmd      string            `json:"healthStartupCmd,omitempty"`
	HealthStartupInterval string            `json:"healthStartupInterval,omitempty"`
	HealthStartupRetries  int               `json:"healthStartupRetries,omitempty"`
	HealthStartupSuccess  int               `json:"healthStartupSuccess,omitempty"`
	HealthStartupTimeout  string            `json:"healthStartupTimeout,omitempty"`
	HealthTimeout         string            `json:"healthTimeout,omitempty"`
	Hostname              string            `json:"hostname,omitempty"`
	HttpProxy             bool              `json:"httpProxy,omitempty"`
	Image                 string            `json:"image,omitempty"`
	IP                    string            `json:"ip,omitempty"`
	IP6                   string            `json:"ip6,omitempty"`
	Label                 map[string]string `json:"labels,omitempty"`
	LogDriver             string            `json:"logDriver,omitempty"`
	LogOpt                []string          `json:"logOptions,omitempty"`
	Mask                  string            `json:"mask,omitempty"`
	Memory                string            `json:"memory,omitempty"`
	ContainersConfModule  []string          `json:"modules,omitempty"`
	Mount                 []string          `json:"mounts,omitempty"`
	ContainerName         string            `json:"name,omitempty"`
	NetworkAlias          []string          `json:"networkAliases,omitempty"`
	Network               []string          `json:"networks,omitempty"`
	NoNewPrivileges       bool              `json:"noNewPrivileges,omitempty"`
	Notify                interface{}       `json:"notify,omitempty"` // null, bool, or "healthy"
	PidsLimit             int               `json:"pidsLimit,omitempty"`
	Pod                   string            `json:"pod,omitempty"`
	PodmanArgs            []string          `json:"podmanArgs,omitempty"`
	PublishPort           []string          `json:"publishPorts,omitempty"`
	Pull                  string            `json:"pull,omitempty"`
	ReadOnly              bool              `json:"readOnly,omitempty"`
	ReadOnlyTmpfs         bool              `json:"readOnlyTmpfs,omitempty"`
	ReloadCmd             interface{}       `json:"reloadCmd,omitempty"` // string or []string
	ReloadSignal          string            `json:"reloadSignal,omitempty"`
	Retry                 int               `json:"retry,omitempty"`
	RetryDelay            string            `json:"retryDelay,omitempty"`
	Rootfs                string            `json:"rootfs,omitempty"`
	RunInit               bool              `json:"runInit,omitempty"`
	SeccompProfile        string            `json:"seccompProfile,omitempty"`
	Secret                []string          `json:"secrets,omitempty"`
	SecurityLabelDisable  bool              `json:"securityLabelDisable,omitempty"`
	SecurityLabelFileType string            `json:"securityLabelFileType,omitempty"`
	SecurityLabelLevel    string            `json:"securityLabelLevel,omitempty"`
	SecurityLabelNested   bool              `json:"securityLabelNested,omitempty"`
	SecurityLabelType     string            `json:"securityLabelType,omitempty"`
	ShmSize               string            `json:"shmSize,omitempty"`
	StartWithPod          bool              `json:"startWithPod,omitempty"`
	StopSignal            string            `json:"stopSignal,omitempty"`
	StopTimeout           int               `json:"stopTimeout,omitempty"`
	SubGIDMap             string            `json:"subGIDMap,omitempty"`
	SubUIDMap             string            `json:"subUIDMap,omitempty"`
	Sysctl                map[string]string `json:"sysctl,omitempty"`
	Timezone              string            `json:"timezone,omitempty"`
	Tmpfs                 []string          `json:"tmpfses,omitempty"`
	UIDMap                []string          `json:"uidMaps,omitempty"`
	Ulimit                []string          `json:"ulimits,omitempty"`
	Unmask                string            `json:"unmask,omitempty"`
	User                  string            `json:"user,omitempty"`
	UserNS                string            `json:"userns,omitempty"`
	Volume                []string          `json:"volumes,omitempty"`
	Workdir               string            `json:"workdir,omitempty"`
}
