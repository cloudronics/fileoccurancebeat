// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

import "time"

type Config struct {
	Period   time.Duration `config:"period"`
	RootPath string        `config:"rootpath"`
	FileName string        `config:"filename"`
}

var DefaultConfig = Config{
	Period:   1 * time.Minute,
	RootPath: "/bin",
	FileName: "ls",
}
