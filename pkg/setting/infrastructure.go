package setting

type Infrastructure struct {
	Firebase Firebase `yaml:"firebase"`
}

type Firebase struct {
	DatabaseURL        string `yaml:"database_url"`
	JsonCredentialFile string `yaml:"json_credential_file"`
}
