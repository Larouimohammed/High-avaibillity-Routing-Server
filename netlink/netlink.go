package netlink

import (
	"log"
	"net"

	"github.com/vishvananda/netlink"
)

func Netlinked(IPaddr net.IP, Mask net.IPMask, Iface string) error {

	//netlink.LinkAdd())

	face, _ := netlink.LinkByName(Iface)

	err := netlink.AddrReplace(face, &netlink.Addr{IPNet: &net.IPNet{
		IP:   IPaddr,
		Mask: Mask,
	}})
	if err != nil {
		log.Fatal(err)

	}

	err = netlink.LinkSetUp(face)

	if err != nil {
		log.Fatal(err)

		return err
	}

	return nil

}
