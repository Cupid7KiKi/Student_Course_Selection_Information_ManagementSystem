package main

import (
	"fmt"
	"net"
	"strings"
)

// getInterfaceByName 获取指定名称的网卡
func getInterfaceByName(name string) *net.Interface {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil
	}
	for _, iface := range interfaces {
		if iface.Name == name {
			return &iface
		}
	}
	return nil
	//, errors.New(fmt.Sprintf("未找到名为 %s 的网卡", name))
}

// getIPv4Addresses 获取网卡的所有IPv4地址
func getIPv4Addresses(iface *net.Interface) string {
	addrs, err := iface.Addrs()
	if err != nil {
		return ""
	}
	var ips []string
	for _, addr := range addrs {
		ipNet, ok := addr.(*net.IPNet)
		if ok && ipNet.IP.To4() != nil {
			ips = append(ips, ipNet.IP.String())
		}
	}
	if len(ips) == 0 {
		return ""
		//, errors.New(fmt.Sprintf("网卡 %s 没有IPv4地址", iface.Name))
	}
	ip := strings.Join(ips, ".")
	return ip
}

func main() {
	ifaceName := "WLAN" // 替换为你的网卡名称
	iface := getInterfaceByName(ifaceName)
	//if err != nil {
	//	fmt.Printf("获取网卡失败: %v\n", err)
	//	return
	//}
	ips := getIPv4Addresses(iface)
	//if err != nil {
	//	fmt.Printf("获取IP地址失败: %v\n", err)
	//	return
	//}
	fmt.Printf("网卡 %s 的IPv4地址: %v\n", ifaceName, ips)
}
