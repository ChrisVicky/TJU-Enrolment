// Load Configuration
package conf

import (
	"errors"
	"os"

	"github.com/pelletier/go-toml/v2"
)

const Fn = "config.toml"

// Account
type Account struct {
	Courses   map[string]string `toml:"courses"`  // courses: {id:comment, id:comment}
	StudentNo string            `toml:"no"`       // Student Number
	Password  string            `toml:"password"` // Password
	Comment   string            `toml:"comment"`  // comment
}

// Program Specification
type Program struct {
	Threads int `toml:"threads"` // Max Threads for Each Course
}

type Ocr struct {
	Payload string `toml:"api"`  // payload
	Type    int    `toml:"type"` // 0, 1, 2
}

// Configuration Main
type Conf struct {
	fn  string    // config file name
	Ocr Ocr       `toml:"Ocr"`
	Ac  []Account `toml:"Account"`
	Pg  Program   `toml:"Program"`
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
	return err
}
