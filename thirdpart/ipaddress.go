package thirdpart

import (
	"sync"

	externalip "github.com/glendc/go-external-ip"
)

var ipOnce sync.Once
var localIP string

func GetLocalIP() string {
	ipOnce.Do(func() {
		consensus := externalip.DefaultConsensus(nil, nil)

		ip, err := consensus.ExternalIP()
		if err == nil {
			localIP = ip.String()
		}
	})

	return localIP

}
