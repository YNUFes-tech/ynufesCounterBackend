package setting

type Service struct {
	Authentication Authentication `yaml:"authentication"`
}

type Authentication struct {
	JWTSecret string `yaml:"jwt_secret"`
}
