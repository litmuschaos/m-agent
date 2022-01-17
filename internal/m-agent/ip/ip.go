// Copyright 2022 LitmusChaos Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package ip

import (
	"io/ioutil"
	"net"
	"net/http"
)

// GetPublicIP returns the public IP address of the machine
func GetPublicIP() string {

	// Using Ipify API to get the Public IP Address
	url := "https://api.ipify.org?format=text"

	resp, err := http.Get(url)
	if err != nil {
		return GetOutboundIP()
	}

	defer resp.Body.Close()
	ip, _ := ioutil.ReadAll(resp.Body)

	return string(ip)
}

// GetOutboundIP returns the outbound IP address of the machine
func GetOutboundIP() string {

	conn, err := net.Dial("udp", "8.8.8.8:80")

	if err != nil {
		return "<your-external-ip-address>"
	}

	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String()
}
