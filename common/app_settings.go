package common

import (
	"github.com/BurntSushi/toml"
)

type NavbarSettings struct {
	Links     []Link            `toml:"links"`
	Dropdowns map[string][]Link `toml:"dropdowns"`
}

type AppSettings struct {
	DatabaseAddress  string             `toml:"database_address"`
	DatabasePort     int                `toml:"database_port"`
	DatabaseUser     string             `toml:"database_user"`
	DatabasePassword string             `toml:"database_password"`
	DatabaseName     string             `toml:"database_name"`
	WebserverPort    int                `toml:"webserver_port"`
	AdminPort        int                `toml:"admin_port"`
	ImageDirectory   string             `toml:"image_dir"`
	CacheEnabled     bool               `toml:"cache_enabled"`
	RecaptchaSiteKey string             `toml:"recaptcha_sitekey,omitempty"`
	RecaptchaSecret  string             `toml:"recaptcha_secret,omitempty"`
	AppNavbar        NavbarSettings     `toml:"navbar"`
	Galleries        map[string]Gallery `toml:"gallery"`
	StickyPosts      []int              `toml:"sticky_posts"`
}

func ReadConfigToml(filepath string) (AppSettings, error) {
	var config AppSettings
	_, err := toml.DecodeFile(filepath, &config)
	if err != nil {
		return AppSettings{}, err
	}

	return config, nil
}
