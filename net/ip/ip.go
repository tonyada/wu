package ip

import (
	"io"
	"net"
	"net/http"
	"strings"
	. "wu"
)

// return LAN IP address
func LanIP() string {
	inter, err := net.Interfaces()
	if Err(err) {
		return ""
	}
	for _, t := range inter {
		aa, err := t.Addrs()
		if err != nil {
			return ""
		}
		for _, a := range aa {
			ipnet, ok := a.(*net.IPNet)
			if !ok {
				continue
			}
			v4 := ipnet.IP.To4()
			if v4 == nil || v4[0] == 127 { // loopback address
				continue
			}
			// Println(v4)
			return v4.String()
		}
	}
	// if ip == "" {
	// 	conn, err := net.Dial("udp", "google.com:80")
	// 	// defer conn.Close()
	// 	if err != nil {
	// 		return ""
	// 	}
	// 	// conn.LocalAddr().String() returns ip_address:port
	// 	ip = strings.Split(conn.LocalAddr().String(), ":")[0]
	// }
	return ""
}

func GetIPFromURL(url string) string {
	resp, err := http.Get(url)
	if Err(err) {
		return ""
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if Err(err) {
		return ""
	}

	ip := string(body)

	if net.ParseIP(ip) == nil {
		Println("Parse IP error from ", url)
		return ""
	}
	return ip
}

// return WAN IP address
func WanIP() string {

	ip := GetIPFromURL("http://whatismyip.akamai.com")

	if ip == "" {
		ip = GetIPFromURL("http://ipecho.net/plain")
	}
	// if ip == "" {
	// 	ip = GetIPFromURL("http://httpbin.org/ip")
	// }
	return ip
}

// return All IPs binded with interfaces like en0, lo0...
func AllIP() string {
	info, _ := net.InterfaceAddrs()
	IPs := ""

	for _, addr := range info {
		ip := strings.Split(addr.String(), "/")[0]
		if len(ip) > 0 {
			IPs += addr.Network() + " | " + addr.String() + " | " + string(ip) + "\n"
		}
	}
	return IPs
}
