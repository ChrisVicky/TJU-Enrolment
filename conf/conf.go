// Load Configuration
package conf

import (
	"enrollment/logger"
	"errors"
	"os"

	"github.com/pelletier/go-toml/v2"
)

const Fn = "config.toml"

// Account
type Account struct {
	StudentNo     string   `toml:"no"`             // Student Number
	Password      string   `toml:"password"`       // Password
	Comment       string   `toml:"comment"`        // comment
	CourseNo      []string `toml:"courses"`        // Course Numbers
	CourseComment []string `toml:"coursesComment"` // courseNames or others
}

// Program Specification
type Program struct {
	Threads int `toml:"threads"` // Max Threads for Each Course
}

// Configuration Main
type Conf struct {
	fn string    // config file name
	Ac []Account `toml:"Account"`
	Pg Program   `toml:"Program"`
}

func NewConfig() *Conf {
	ret := &Conf{fn: Fn}
	return ret
}

func (c *Conf) SetFn(fn string) {
	c.fn = fn
}

func (c *Conf) LoadConfig() error {
	if _, err := os.Stat(c.fn); err != nil {
		c.fn = Fn
	}
	b, err := os.ReadFile(c.fn)
	if err != nil {
		return err
	}
	if err := toml.Unmarshal(b, c); err != nil {
		return err
	}
	return c.clean()
}

var ErrorConfig = errors.New("configuration error")

func (c *Conf) clean() (err error) {
	for _, a := range c.Ac {
		if len(a.CourseNo) != len(a.CourseComment) {
			err = ErrorConfig
			logger.Errorf("%v Must Aligned with %v", a.CourseNo, a.CourseComment)
		}
	}
	return err
}
