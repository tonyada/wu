package net

import (
	"net"
)

func LocalMACAddr() (macAddr string, err error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, inter := range interfaces {
		mac := inter.HardwareAddr
		macAddr = mac.String()
		if macAddr != "" {
			break
		}
	}
	return macAddr, nil
}

func LocalMACAddrs() (macAddrs []string, err error) {
	interfaces, err := net.Interfaces()
	for _, inter := range interfaces {
		addr := inter.HardwareAddr.String()
		if addr != "" {
			macAddrs = append(macAddrs, addr)
		}
	}
	return macAddrs, err
}

func LocalInterfacesWithMACAddr() (inters map[string]string, err error) {
	inters = make(map[string]string)
	interfaces, err := net.Interfaces()
	for _, inter := range interfaces {
		addr := inter.HardwareAddr.String()
		if addr != "" {
			inters[inter.Name] = addr
		}
	}
	return inters, err
}

// func main() {
// 	info, _ := net.InterfaceAddrs()
// 	for _, addr := range info {
// 		fmt.Println(strings.Split(addr.String(), "/")[0])
// 	}

// 	conn, err := net.Dial("udp", "google.com:80")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}
// 	defer conn.Close()
// 	fmt.Println(strings.Split(conn.LocalAddr().String(), ":")[0])
// }
// func Download(durl string, filename string) (err error) {
// 	_, err = url.Parse(durl)
// 	if err != nil {
// 		return
// 	}
// 	if filename == "" {
// 		// auto get file name
// 		parts := strings.Split(durl, "/")
// 		filename = parts[len(parts)-1]
// 		if filename == "" {
// 			filename = "downloadFile"
// 		}
// 	}
// 	Printfln("downloading...%s", filename)

// 	resp, err := http.Get(durl)
// 	if err != nil {
// 		return
// 	}
// 	defer resp.Body.Close()

// 	f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0777)
// 	if err != nil {
// 		return
// 	}
// 	defer f.Close()

// 	_, err = io.Copy(f, resp.Body)
// 	if err != nil {
// 		return
// 	}
// 	return nil
// }
