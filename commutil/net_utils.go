package commutil

import (
	"github.com/xkniu/go-commons-lang/core"
	"net"
	"sync"
)

var NetUtils *netUtils

func init() {
	instance := new(netUtils)
	NetUtils = core.InitializeIfNeeded(instance).(*netUtils)
}

type netUtils struct {
	once      sync.Once
	localIpv4 string // local IPv4 cached result
}

// Get local IPv4 address. The value would be cached after first get.
func (t *netUtils) GetLocalIpv4() string {
	t.once.Do(func() {
		if ias, err := net.InterfaceAddrs(); err == nil {
			for _, address := range ias {
				if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
					if ipNet.IP.To4() != nil {
						t.localIpv4 = ipNet.IP.String()
						return
					}
				}
			}
		}
	})
	return t.localIpv4
}
