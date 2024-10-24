package models

type Clouds struct {
	Clouds map[string]Cloud `yaml:"clouds"`
}

type Cloud struct {
	Auth               Auth   `yaml:"auth"`
	RegionName         string `yaml:"region_name"`
	Interface          string `yaml:"interface"`
	IdentityAPIVersion int    `yaml:"identity_api_version"`
}

type Auth struct {
	AuthURL         string `yaml:"auth_url"`
	Username        string `yaml:"username"`
	ProjectID       string `yaml:"project_id"`
	ProjectName     string `yaml:"project_name"`
	UserDomainName  string `yaml:"user_domain_name"`
	Password        string `yaml:"password"`
	ProjectDomainID string `yaml:"project_domain_id"`
}
