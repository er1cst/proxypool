package proxy

import (
	"errors"
	"net"
	"sync"
)

type ProxyPool struct {
	Anonymous bool

	count int
	curr  int

	mutex          *sync.Mutex
	proxies        []*Proxy
	anonyCheckFunc AnonymityCheckerFunc
	availCheckFunc AvailablilityCheckerFunc
}

type Proxy struct {
	ip net.IPAddr
}

type AnonymityCheckerFunc func(p Proxy) bool

type AvailabilityCheckerFunc func(p Proxy) bool

func New(volume int) (pool ProxyPool) {
	pool = &ProxyPool{
		Anonymous: false,
		proxies:   make([]*Proxy, volume),
		mutex:     *sync.Mutex{},
	}
	return
}

func (*ProxyPool) Add(ip string) error {
	parsed := net.ParseIP(ip)
	if parsed != nil {
		p := new(Proxy)
		p.ip = parsed
		go checkAndAddProxy(p)
		return nil
	}
	return errors.New(fmt.Sprintf("the input \"%s\" is not a valid IP address", ip))
}

func (*ProxyPool) checkAndAddProxy(p Proxy) error {
	err := checkProxy(p)
	if err != nil {
		return err
	}

}

func (*ProxyPool) checkProxy(p Proxy) error {

}

func (pool *ProxyPool) addProxy(p Proxy) error {
	pool.mutex.Lock()
	defer pool.mutex.Unlock()

	vol := len(pool.proxies)
	if pool.count == vol {
		return errors.New("the proxy pool is full")
	} else {
		pool.curr = (curr + 1) % vol
		count++
	}

}
