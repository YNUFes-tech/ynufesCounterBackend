package setting

type (
	ThirdParty struct {
		Line Line `yaml:"line"`
	}
	Line struct {
		ClientID       string `yaml:"client_id"`
		ClientSecret   string `yaml:"client_secret"`
		CallbackURI    string `yaml:"callback_uri"`
		CipherKey      string `yaml:"cipher_key"`
		EnableLineAuth bool   `yaml:"enable_line_auth"`
	}
)
