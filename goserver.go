/* -.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.

* File Name : goserver.go

* Purpose :

* Creation Date : 07-24-2014

* Last Modified : Tue 29 Jul 2014 12:19:38 AM UTC

* Created By : Kiyor

_._._._._._._._._._._._._._._._._._._._._.*/

package goserver

import (
	"encoding/json"
	"errors"
	"github.com/kiyor/subnettool"
	"io/ioutil"
	"net"
	// 	"os"
	"strings"
)

func (this *IpInfo) getIP() {
	ips, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}
	this.Ip = Ip{}
	for _, v := range ips {
		ipstr := strings.Split(v.String(), "/")[0]
		ip := net.ParseIP(ipstr)
		if ip.To4() != nil {
			token := subnettool.ParseIPInt(ip.To4())
			if token[0] == 127 {
				continue
			} else if token[0] == 10 || token[0] == 192 || token[0] == 176 {
				this.Ip.PrivIps = append(this.Ip.PrivIps, ip)
			} else {
				this.Ip.PubIps = append(this.Ip.PubIps, ip)
			}
		}
	}
}

func GetHostname() (string, error) {
	b, err := ioutil.ReadFile("/etc/sysconfig/network")
	if err != nil {
		return "", err
	}
	lines := strings.Split(string(b), "\n")
	for _, line := range lines {
		if strings.Contains(line, "HOSTNAME") {
			return strings.Split(line, "=")[1], nil
		}
	}
	return "", errors.New("not found")
}

func (this *IpInfo) Info() string {
	// 	this.getHostname()
	this.getIP()
	j, _ := json.Marshal(this)
	return string(j)
}
