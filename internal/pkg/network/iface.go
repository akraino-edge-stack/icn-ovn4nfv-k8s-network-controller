package network

import (
	"errors"
	"fmt"
	"net"
	"strings"
	"syscall"

	"github.com/vishvananda/netlink"
)

//GetGatewayInterface return gw string for a given dst route
func GetGatewayInterface(dst string) (string, error) {
	afInetVersion := syscall.AF_INET
	if strings.Contains(dst, ":") {
		afInetVersion = syscall.AF_INET6
	}
	routes, err := netlink.RouteList(nil, afInetVersion)
	if err != nil {
		return "", err
	}

	for _, route := range routes {
		if route.Dst != nil {
			if route.Dst.String() == dst {
				if afInetVersion == syscall.AF_INET {
					if route.Gw.To4() == nil {
						return "", errors.New("Found IPv4 dst route but could not determine gateway")
					}
					return route.Gw.To4().String(), nil
				} else {
					if route.Gw.To16() == nil {
						return "", errors.New("Found IPv6 dst route but could not determine gateway")
					}
					return route.Gw.To16().String(), nil
				}
			}
		}
	}

	return "", fmt.Errorf("Unable to find gw for route dst -%s", dst)
}

//IsRouteExist return true for gw string for a given dst route
func IsRouteExist(dst string, gw string) (bool, error) {
	afInetVersion := syscall.AF_INET
	if strings.Contains(dst, ":") {
		afInetVersion = syscall.AF_INET6
	}
	routes, err := netlink.RouteList(nil, afInetVersion)
	if err != nil {
		return false, err
	}

	for _, route := range routes {
		if route.Dst != nil {
			if route.Dst.String() == dst {
				if afInetVersion == syscall.AF_INET {
					if route.Gw.To4() == nil {
						return false, nil
					}
				} else {
					if route.Gw.To16() == nil {
						return false, nil
					}
				}
				if route.Gw.String() == gw {
					return true, nil
				}
			}
		}
	}

	return false, nil
}

//GetDefaultGateway return default gateway of the network namespace
func GetDefaultGateway(afInetVersion int) (string, error) {
	fmt.Println("GetDefaultGateway")
	routes, err := netlink.RouteList(nil, afInetVersion)
	if err != nil {
		return "", err
	}

	zeroAddress := "0.0.0.0/0"

	if afInetVersion == syscall.AF_INET6 {
		zeroAddress = "::"
	}

	for _, route := range routes {
		fmt.Printf("GetDefaultGateway route: %v\n", route)
		if route.Dst == nil || route.Dst.String() == zeroAddress {
			if route.Gw.To4() == nil {
				return "", errors.New("Found default route but could not determine gateway")
			}
			var gateway string
			if afInetVersion == syscall.AF_INET {
				gateway = route.Gw.To4().String()
			} else {
				gateway = route.Gw.To16().String()
			}
			return gateway, nil
		}
	}

	return "", errors.New("Unable to find default route")
}

//CheckRoute return bool isPresent
func CheckRoute(dst, gw string) (bool, error) {
	var isPresent bool
	afInetVersion := syscall.AF_INET
	if strings.Contains(dst, ":") {
		afInetVersion = syscall.AF_INET6
	}
	routes, err := netlink.RouteList(nil, afInetVersion)
	if err != nil {
		return isPresent, err
	}

	for _, route := range routes {
		var routeGW string
		if afInetVersion == syscall.AF_INET {
			routeGW = route.Gw.To4().String() 
		} else {
			routeGW = route.Gw.To16().String() 
		}
		if route.Dst.String() == dst && routeGW == gw {
			isPresent = true
		}
	}

	return isPresent, nil

}

// GetDefaultGatewayInterface return default gateway interface link
func getDefaultGatewayInterface(afInetVersion int) (*net.Interface, error) {
	fmt.Println("getDefaultGatewayInterface")
	routes, err := netlink.RouteList(nil, afInetVersion)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Routes: %v\n", routes)

	zeroAddress :=  "0.0.0.0/0"

	if afInetVersion == syscall.AF_INET6 {
		zeroAddress = "::"
	}

	for _, route := range routes {
		fmt.Printf("getDefaultGatewayInterface route: %v\n", route)
		if route.Dst == nil || route.Dst.String() == zeroAddress {
			if route.LinkIndex <= 0 {
				return nil, errors.New("Found default route but could not determine interface")
			}
			return net.InterfaceByIndex(route.LinkIndex)
		}
	}

	return nil, errors.New("Unable to find default route")
}

func getIfaceAddrs(iface *net.Interface, afInetVersion int) ([]netlink.Addr, error) {

	link := &netlink.Device{
		netlink.LinkAttrs{
			Index: iface.Index,
		},
	}

	return netlink.AddrList(link, afInetVersion)
}

//GetInterfaceIP4Addr return IP4addr of a interface
func GetInterfaceIPAddr(iface *net.Interface, afInetVersion int) (netlink.Addr, error) {
	addrs, err := getIfaceAddrs(iface, afInetVersion)
	if err != nil {
		return netlink.Addr{}, err
	}

	// prefer non link-local addr
	var ll netlink.Addr

	for _, addr := range addrs {

		if afInetVersion == syscall.AF_INET {
			if addr.IP.To4() == nil {
				continue
			}
		} else {
			if addr.IP.To16() == nil {
				continue
			}
		}

		if addr.IP.IsGlobalUnicast() {
			return addr, nil
		}

		if addr.IP.IsLinkLocalUnicast() {
			ll = addr
		}
	}

	// didn't find global but found link-local. it'll do.
	if afInetVersion == syscall.AF_INET {
		if ll.IP.To4() != nil {
			return ll, nil
		}
	} else {
		if ll.IP.To16() != nil {
			return ll, nil
		}
	}

	return netlink.Addr{}, errors.New("No IP address found for given interface")
}

//GetHostNetwork return default gateway interface network
func GetHostNetwork(afInetVersion int) (string, error) {

	iface, err := getDefaultGatewayInterface(afInetVersion)
	if err != nil {
		fmt.Printf("error in gettting default gateway interface: %s\n", err.Error())
		log.Error(err, "error in gettting default gateway interface")
		return "", err
	}

	ipaddr, err := GetInterfaceIPAddr(iface, afInetVersion)
	if err != nil {
		fmt.Printf("error in gettting default gateway interface IP address: %s\n", err.Error())
		log.Error(err, "error in gettting default gateway interface IP address")
		return "", err
	}

	_, ipNet, err := net.ParseCIDR(ipaddr.IPNet.String())
	if err != nil {
		fmt.Printf("error in gettting default gateway interface network: %s\n", err.Error())
		log.Error(err, "error in gettting default gateway interface network")
		return "", err
	}

	return ipNet.String(), nil
}
