package env

type Conf struct {
	App struct {
		Name       string `env:"APP_NAME"`
		Port       string `env:"APP_PORT"`
		Mode       string `env:"APP_MODE"`
		Url        string `env:"APP_URL"`
		Secret_key string `env:"APP_SECRET"`
	}
	DB struct {
		Host string `env:"DB_HOST"`
		Name string `env:"DB_NAME"`
		User string `env:"DB_USER"`
		Pass string `env:"DB_PASSWORD"`
		Port string `env:"DB_PORT"`
	}
	DB_Prod struct {
		Host string `env:"DB_HOST_PROD"`
		Name string `env:"DB_NAME_PROD"`
		User string `env:"DB_USER_PROD"`
		Pass string `env:"DB_PASSWORD_PROD"`
		Port string `env:"DB_PORT_PROD"`
	}
}
