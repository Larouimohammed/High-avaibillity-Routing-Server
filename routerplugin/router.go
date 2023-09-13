package routerplugin

import (
	db "HA/DB"
	"HA/netlink"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

type Router struct {
	Key    string
	Role   string
	Status bool
	IPaddr []byte
	Mask   []byte
	Iface  string
}

func NewRouter(key string, role string, status bool, ipaddr []byte, mask []byte, iface string) *Router {
	return &Router{
		Key:    key,
		Role:   role,
		Status: status,
		IPaddr: ipaddr,
		Mask:   mask,
		Iface:  iface,
	}
}

func (r *Router) Run(wg *sync.WaitGroup) {
	go func(wg *sync.WaitGroup) {
		wg.Add(1)
		logrus.WithFields(log.Fields{
			"time":           time.Now().Format(time.RFC3339),
			"Server Running": r.Key,
		}).Info("Server running")

		defer wg.Done()
		r.Status = true
		r.Role = "Primary"
		netlink.Netlinked(r.IPaddr, r.Mask, r.Iface)

		db.Publish(r.Key)
	}(wg)
}

func (r *Router) Watcher(r2 *Router, wg *sync.WaitGroup) {
	go func(wg *sync.WaitGroup) {
		wg.Add(1)
		logrus.WithFields(log.Fields{
			"Server":   r.Key,
			"Watch To": r2.Key,
			"time":     time.Now().Format(time.RFC3339),
		}).Info("Start watching")

		defer wg.Done()

		db.Watch(r2.Key)

	}(wg)
}
