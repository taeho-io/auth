package server

type Config interface {
	Settings() Settings
}

type DefaultConfig struct {
	Config

	settings Settings
}

func NewConfig(settings Settings) (cfg Config) {
	return &DefaultConfig{
		settings: settings,
	}
}

func MockConfig() (cfg Config) {
	return NewConfig(MockSettings())
}

func (c *DefaultConfig) Settings() Settings {
	return c.settings
}
