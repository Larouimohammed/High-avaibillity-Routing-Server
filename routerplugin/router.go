package routerplugin

import (
	db "HA/DB"
	"HA/netlink"
	"sync"
	"time"

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
	log.WithFields(log.Fields{
		"time":   time.Now().Format(time.RFC3339),
		"Server": r.Key,
	}).Info("Server running")

	defer wg.Done()
	r.Status = true
	r.Role = "Premary"
	netlink.Netlinked(r.IPaddr, r.Mask, r.Iface)

	db.Publish(r.Key)

}

func (r *Router) Watcher(r2 *Router, wg *sync.WaitGroup) {
	log.WithFields(log.Fields{
		"Server": r.Key,
		"Watch To": r2.Key,
		"time":   time.Now().Format(time.RFC3339),
		
	}).Info("Start watching")

	defer wg.Done()
	//fmt.Printf("Start %v watching To %v \n", r.Key, r2.Key)

	db.Watch(r2.Key)

}
