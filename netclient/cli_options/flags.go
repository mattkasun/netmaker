package cli_options

import "github.com/urfave/cli/v2"

// GetFlags - Returns the flags used by cli
func GetFlags(hostname string) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    "network",
			Aliases: []string{"n"},
			EnvVars: []string{"NETCLIENT_NETWORK"},
			Value:   "all",
			Usage:   "Network to perform specified action against.",
		},
		&cli.StringFlag{
			Name:    "password",
			Aliases: []string{"p"},
			EnvVars: []string{"NETCLIENT_PASSWORD"},
			Value:   "",
			Usage:   "Password for authenticating with netmaker.",
		},
		&cli.StringFlag{
			Name:    "endpoint",
			Aliases: []string{"e"},
			EnvVars: []string{"NETCLIENT_ENDPOINT"},
			Value:   "",
			Usage:   "Reachable (usually public) address for WireGuard (not the private WG address).",
		},
		&cli.StringFlag{
			Name:    "macaddress",
			Aliases: []string{"m"},
			EnvVars: []string{"NETCLIENT_MACADDRESS"},
			Value:   "",
			Usage:   "Mac Address for this machine. Used as a unique identifier within Netmaker network.",
		},
		&cli.StringFlag{
			Name:    "publickey",
			Aliases: []string{"pubkey"},
			EnvVars: []string{"NETCLIENT_PUBLICKEY"},
			Value:   "",
			Usage:   "Public Key for WireGuard Interface.",
		},
		&cli.StringFlag{
			Name:    "privatekey",
			Aliases: []string{"privkey"},
			EnvVars: []string{"NETCLIENT_PRIVATEKEY"},
			Value:   "",
			Usage:   "Private Key for WireGuard Interface.",
		},
		&cli.StringFlag{
			Name:    "port",
			EnvVars: []string{"NETCLIENT_PORT"},
			Value:   "",
			Usage:   "Port for WireGuard Interface.",
		},
		&cli.IntFlag{
			Name:    "keepalive",
			EnvVars: []string{"NETCLIENT_KEEPALIVE"},
			Value:   0,
			Usage:   "Default PersistentKeepAlive for Peers in WireGuard Interface.",
		},
		&cli.StringFlag{
			Name:    "operatingsystem",
			Aliases: []string{"os"},
			EnvVars: []string{"NETCLIENT_OS"},
			Value:   "",
			Usage:   "Identifiable name for machine within Netmaker network.",
		},
		&cli.StringFlag{
			Name:    "name",
			EnvVars: []string{"NETCLIENT_NAME"},
			Value:   hostname,
			Usage:   "Identifiable name for machine within Netmaker network.",
		},
		&cli.StringFlag{
			Name:    "localaddress",
			EnvVars: []string{"NETCLIENT_LOCALADDRESS"},
			Value:   "",
			Usage:   "Local address for machine. Can be used in place of Endpoint for machines on the same LAN.",
		},
		&cli.StringFlag{
			Name:    "isstatic",
			Aliases: []string{"st"},
			EnvVars: []string{"NETCLIENT_IS_STATIC"},
			Value:   "",
			Usage:   "Indicates if client is static by default (will not change addresses automatically).",
		},
		&cli.StringFlag{
			Name:    "address",
			Aliases: []string{"a"},
			EnvVars: []string{"NETCLIENT_ADDRESS"},
			Value:   "",
			Usage:   "WireGuard address for machine within Netmaker network.",
		},
		&cli.StringFlag{
			Name:    "addressIPv6",
			Aliases: []string{"a6"},
			EnvVars: []string{"NETCLIENT_ADDRESSIPV6"},
			Value:   "",
			Usage:   "WireGuard address for machine within Netmaker network.",
		},
		&cli.StringFlag{
			Name:    "interface",
			Aliases: []string{"i"},
			EnvVars: []string{"NETCLIENT_INTERFACE"},
			Value:   "",
			Usage:   "WireGuard local network interface name.",
		},
		&cli.StringFlag{
			Name:    "apiserver",
			EnvVars: []string{"NETCLIENT_API_SERVER"},
			Value:   "",
			Usage:   "Address + API Port (e.g. 1.2.3.4:8081) of Netmaker server.",
		},
		&cli.StringFlag{
			Name:    "grpcserver",
			EnvVars: []string{"NETCLIENT_GRPC_SERVER"},
			Value:   "",
			Usage:   "Address + GRPC Port (e.g. 1.2.3.4:50051) of Netmaker server.",
		},
		&cli.StringFlag{
			Name:    "grpcssl",
			EnvVars: []string{"NETCLIENT_GRPCSSL"},
			Value:   "",
			Usage:   "Tells clients to use SSL to connect to GRPC if 'on'. Disable if 'off'. Off by default.",
		},
		&cli.StringFlag{
			Name:    "key",
			Aliases: []string{"k"},
			EnvVars: []string{"NETCLIENT_ACCESSKEY"},
			Value:   "",
			Usage:   "Access Key for signing up machine with Netmaker server during initial 'add'.",
		},
		&cli.StringFlag{
			Name:    "token",
			Aliases: []string{"t"},
			EnvVars: []string{"NETCLIENT_ACCESSTOKEN"},
			Value:   "",
			Usage:   "Access Token for signing up machine with Netmaker server during initial 'add'.",
		},
		&cli.StringFlag{
			Name:    "localrange",
			EnvVars: []string{"NETCLIENT_LOCALRANGE"},
			Value:   "",
			Usage:   "Local Range if network is local, for instance 192.168.1.0/24.",
		},
		&cli.StringFlag{
			Name:    "dnson",
			EnvVars: []string{"NETCLIENT_DNS"},
			Value:   "yes",
			Usage:   "Sets private dns if 'yes'. Ignores if 'no'. Will retrieve from network if unset.",
		},
		&cli.StringFlag{
			Name:    "islocal",
			EnvVars: []string{"NETCLIENT_IS_LOCAL"},
			Value:   "",
			Usage:   "Sets endpoint to local address if 'yes'. Ignores if 'no'. Will retrieve from network if unset.",
		},
		&cli.StringFlag{
			Name:    "isdualstack",
			EnvVars: []string{"NETCLIENT_IS_DUALSTACK"},
			Value:   "",
			Usage:   "Sets ipv6 address if 'yes'. Ignores if 'no'. Will retrieve from network if unset.",
		},
		&cli.StringFlag{
			Name:    "udpholepunch",
			EnvVars: []string{"NETCLIENT_UDP_HOLEPUNCH"},
			Value:   "",
			Usage:   "Turns on udp holepunching if 'yes'. Ignores if 'no'. Will retrieve from network if unset.",
		},
		&cli.StringFlag{
			Name:    "ipforwarding",
			EnvVars: []string{"NETCLIENT_IPFORWARDING"},
			Value:   "yes",
			Usage:   "Sets ip forwarding on if 'yes'. Ignores if 'no'. On by default.",
		},
		&cli.StringFlag{
			Name:    "postup",
			EnvVars: []string{"NETCLIENT_POSTUP"},
			Value:   "",
			Usage:   "Sets PostUp command for WireGuard.",
		},
		&cli.StringFlag{
			Name:    "postdown",
			EnvVars: []string{"NETCLIENT_POSTDOWN"},
			Value:   "",
			Usage:   "Sets PostDown command for WireGuard.",
		},
		&cli.StringFlag{
			Name:    "daemon",
			EnvVars: []string{"NETCLIENT_DAEMON"},
			Value:   "on",
			Usage:   "Installs daemon if 'on'. Ignores if 'off'. On by default.",
		},
		&cli.StringFlag{
			Name:    "roaming",
			EnvVars: []string{"NETCLIENT_ROAMING"},
			Value:   "yes",
			Usage:   "Checks for IP changes if 'yes'. Ignores if 'no'. Yes by default.",
		},
		&cli.StringFlag{
			Name:    "force",
			EnvVars: []string{"NETCLIENT_FORCE"},
			Value:   "no",
			Usage:   "Allows to run the command with force, if otherwise prevented.",
		},
		&cli.BoolFlag{
			Name:    "verbosity-level-1",
			Aliases: []string{"v"},
			Value:   false,
			Usage:   "Netclient Verbosity level 1.",
		},
		&cli.BoolFlag{
			Name:    "verbosity-level-2",
			Aliases: []string{"vv"},
			Value:   false,
			Usage:   "Netclient Verbosity level 2.",
		},
		&cli.BoolFlag{
			Name:    "verbosity-level-3",
			Aliases: []string{"vvv"},
			Value:   false,
			Usage:   "Netclient Verbosity level 3.",
		},
	}
}
