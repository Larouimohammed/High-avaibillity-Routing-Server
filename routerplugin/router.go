package routerplugin

import (
	db "HA/DB"
	"HA/netlink"
	"fmt"
	"sync"
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
	defer wg.Done()
	fmt.Printf("Server %v running\n", r.Key)
	r.Status = true
	r.Role = "Premary"
	netlink.Netlinked(r.IPaddr, r.Mask, r.Iface)

	db.Publish(r.Key)

}

func (r *Router) Watcher(r2 *Router, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Start %v watching To %v \n", r.Key, r2.Key)

	db.Watch(r2.Key)

}
