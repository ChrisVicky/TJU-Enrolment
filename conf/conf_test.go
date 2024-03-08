package conf_test

import (
	"enrolment/conf"
	"testing"
)

func TestNewConf(t *testing.T) {
	c := conf.NewConfig()
	c.SetFn("example.toml")
	if err := c.LoadConfig(); err != nil {
		t.Error(err)
	}
}

func TestExampleConf(t *testing.T) {
	c := conf.NewConfig()
	c.SetFn("example.toml")
	if err := c.LoadConfig(); err != nil {
		t.Error(err)
	}
	t.Logf("%+v", c)
}
