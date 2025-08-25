package entities

import "fmt"

type Settings struct {
	Port int
}

func (s Settings) GetPort() string {
	return fmt.Sprintf(":%d", s.Port)
}
