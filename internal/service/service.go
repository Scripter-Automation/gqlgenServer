package service

import (
	"fmt"
	"log"
)

//Services manage external service connections
type Services struct {
	Servers   []ExternalServer
	Failed    []ExternalServer
	Connected []ExternalServer
}

//PrintConnected prints Connected
func (c *Services) PrintConnected() {
	for _, server := range c.Connected {
		log.Printf(server.Name() + " Connected Succefully")
	}
}

//PrintFailed failed
func (c *Services) PrintFailed() {
	for _, server := range c.Failed {
		println(server.Name() + " Failed")
	}
}

func (c *Services) Connect() error {
	for _, server := range c.Servers {
		err := server.AttemptConnection()

		if err != nil {
			c.Failed = append(c.Failed, server)
		} else {
			//add to Connected list
			c.Connected = append(c.Connected, server)
		}
	}

	if len(c.Failed) > 0 {
		return fmt.Errorf("some services failed to connect")
	}

	println(len(c.Failed))
	//no response
	return nil
}

//ExternalServer interface
type ExternalServer interface {
	AttemptConnection() error
	Name() string
}
