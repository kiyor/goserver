/* -.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.

* File Name : struct.go

* Purpose :

* Creation Date : 07-24-2014

* Last Modified : Mon 28 Jul 2014 06:28:47 PM UTC

* Created By : Kiyor

_._._._._._._._._._._._._._._._._._._._._.*/

package goserver

import (
	"net"
)

type Ip struct {
	PubIps  []net.IP `json:"pub_ips"`
	PrivIps []net.IP `json:"priv_ips"`
}
type IpInfo struct {
	Ip Ip `json:"ip"`
}
type Group struct {
	Id      int               `json:"id"`
	Name    string            `json:"name"`
	Servers map[string]IpInfo `json:"servers"`
}

type ServerCert struct {
	Certs []Cert `json:"certs"`
}

type Cert struct {
	Ip   net.IP `json:"ip"`
	Cert string `json:"cert"`
}
