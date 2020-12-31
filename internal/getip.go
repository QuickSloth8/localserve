package internal

import (
  "net"
  "log"
)

func GetIp() string {
  ifaces, err := net.Interfaces()
  if err != nil {
    log.Fatal(err)
  }
  // handle err
  for _, i := range ifaces {
      addrs, err := i.Addrs()
      if err != nil {
        log.Fatal(err)
      }
      // handle err
      for _, addr := range addrs {
          var ip net.IP
          switch v := addr.(type) {
          case *net.IPNet:
                  ip = v.IP
          case *net.IPAddr:
                  ip = v.IP
          }
          // process IP address
          ipv4 := ip.To4()
          if ipv4 != nil && (ipv4.IsLoopback() == false) {
            return ipv4.String()
          }
      }
  }
  return ""
}
