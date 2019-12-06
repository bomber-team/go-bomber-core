package discovery

import (
	"net"
	"time"

	"github.com/go-akka/configuration"
)

/*Speaker - producer packets for discovered by main backend service on multicast listen*/
type Speaker struct {
	Timeout         int
	Packet          []byte
	Address         string
	MaxDatagramSize int32
}

const (
	timeoutField = "timeout"
)

func NewSpeaker(config *configuration.Config) *Speaker {
	res := &Speaker{}
	res.parsingParamsFromConfig(config)
}

func (speak *Speaker) parsingParamsFromConfig(config *configuration.Config) error {
	speak.Timeout = configuration_core.TryParsingNumbers(config, timeoutField)

}

func (speak *Speaker) RunEcho() error {
	addr, err := net.ResolveUDPAddr("udp", speak.Address)
	if err != nil {
		return err
	}
	connection, err2 := net.DialUDP("udp", nil, addr)
	if err2 != nil {
		return err2
	}

	for {
		connection.Write(speak.Packet)
		time.Sleep(time.Second * time.Duration(speak.Timeout))
	}
}
