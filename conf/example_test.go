package conf_test

import (
	"enrolment/conf"
	"fmt"
)

func ExampleConf() {
	c := conf.NewConfig()
	c.SetFn("example.toml")
	if err := c.LoadConfig(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Loaded")
	// Output:
	// Loaded
}
