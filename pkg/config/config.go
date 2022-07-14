package config

import (
	"flag"
	"os"
)

type Config struct {
	apiKey     string
	domainName string
}

func Get() *Config {
	conf := &Config{}

	flag.StringVar(&conf.apiKey, "apiKey", os.Getenv("API_KEY"), "Freshdesk API Key")
	flag.StringVar(&conf.domainName, "freshDeskDomainName", os.Getenv("DOMAIN_NAME"), "Fresh Desk Domain Name")

	flag.Parse()

	return conf
}

func (c *Config) GetApiKey() string {
	return c.apiKey
}

func (c *Config) GetDomainName() string {
	return c.domainName
}

func (c *Config) GetFreshDeskURL() string {
	return "https://" + c.domainName + ".freshdesk.com"
}
