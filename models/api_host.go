package models

import "net"

// ApiHost - the host struct for API usage
type ApiHost struct {
	ID              string   `json:"id"`
	Verbosity       int      `json:"verbosity"`
	FirewallInUse   string   `json:"firewallinuse"`
	Version         string   `json:"version"`
	Name            string   `json:"name"`
	OS              string   `json:"os"`
	Debug           bool     `json:"debug"`
	IsStatic        bool     `json:"isstatic"`
	ListenPort      int      `json:"listenport"`
	LocalAddress    string   `json:"localaddress"`
	LocalRange      string   `json:"localrange"`
	LocalListenPort int      `json:"locallistenport"`
	ProxyListenPort int      `json:"proxy_listen_port"`
	MTU             int      `json:"mtu" yaml:"mtu"`
	Interfaces      []Iface  `json:"interfaces" yaml:"interfaces"`
	EndpointIP      string   `json:"endpointip" yaml:"endpointip"`
	PublicKey       string   `json:"publickey"`
	MacAddress      string   `json:"macaddress"`
	InternetGateway string   `json:"internetgateway"`
	Nodes           []string `json:"nodes"`
	ProxyEnabled    bool     `json:"proxy_enabled" yaml:"proxy_enabled"`
	IsDefault       bool     `json:"isdefault" yaml:"isdefault"`
}

// Host.ConvertNMHostToAPI - converts a Netmaker host to an API editable host
func (h *Host) ConvertNMHostToAPI() *ApiHost {
	a := ApiHost{}
	a.Debug = h.Debug
	a.EndpointIP = h.EndpointIP.String()
	a.FirewallInUse = h.FirewallInUse
	a.ID = h.ID.String()
	a.Interfaces = h.Interfaces
	a.InternetGateway = h.InternetGateway.String()
	if isEmptyAddr(a.InternetGateway) {
		a.InternetGateway = ""
	}
	a.IsStatic = h.IsStatic
	a.ListenPort = h.ListenPort
	a.LocalAddress = h.LocalAddress.String()
	if isEmptyAddr(a.LocalAddress) {
		a.LocalAddress = ""
	}
	a.LocalListenPort = h.LocalListenPort
	a.LocalRange = h.LocalRange.String()
	if isEmptyAddr(a.LocalRange) {
		a.LocalRange = ""
	}
	a.MTU = h.MTU
	a.MacAddress = h.MacAddress.String()
	a.Name = h.Name
	a.OS = h.OS
	a.Nodes = h.Nodes
	a.ProxyEnabled = h.ProxyEnabled
	a.ProxyListenPort = h.ProxyListenPort
	a.PublicKey = h.PublicKey.String()
	a.Verbosity = h.Verbosity
	a.Version = h.Version
	a.IsDefault = h.IsDefault

	return &a
}

// APIHost.ConvertAPIHostToNMHost - convert's a given apihost struct to
// a Host struct
func (a *ApiHost) ConvertAPIHostToNMHost(currentHost *Host) *Host {
	h := Host{}
	h.ID = currentHost.ID
	h.HostPass = currentHost.HostPass
	h.DaemonInstalled = currentHost.DaemonInstalled
	h.EndpointIP = net.ParseIP(a.EndpointIP)
	h.Debug = a.Debug
	h.FirewallInUse = a.FirewallInUse
	h.IPForwarding = currentHost.IPForwarding
	h.Interface = currentHost.Interface
	h.Interfaces = currentHost.Interfaces
	h.InternetGateway = currentHost.InternetGateway
	h.IsDocker = currentHost.IsDocker
	h.IsK8S = currentHost.IsK8S
	h.IsStatic = a.IsStatic
	h.ListenPort = a.ListenPort
	h.LocalListenPort = currentHost.ListenPort
	h.MTU = a.MTU
	h.MacAddress = currentHost.MacAddress
	h.PublicKey = currentHost.PublicKey
	h.Name = a.Name
	h.Version = currentHost.Version
	h.Verbosity = a.Verbosity
	h.Nodes = currentHost.Nodes
	h.TrafficKeyPublic = currentHost.TrafficKeyPublic
	h.OS = currentHost.OS
	if len(a.LocalAddress) > 0 {
		_, localAddr, err := net.ParseCIDR(a.LocalAddress)
		if err == nil {
			h.LocalAddress = *localAddr
		}
	} else if !isEmptyAddr(currentHost.LocalAddress.String()) {
		h.LocalAddress = currentHost.LocalAddress
	}
	if len(a.LocalRange) > 0 {
		_, localRange, err := net.ParseCIDR(a.LocalRange)
		if err == nil {
			h.LocalRange = *localRange
		}
	} else if !isEmptyAddr(currentHost.LocalRange.String()) {
		h.LocalRange = currentHost.LocalRange
	}
	h.ProxyEnabled = a.ProxyEnabled
	h.IsDefault = a.IsDefault

	return &h
}
