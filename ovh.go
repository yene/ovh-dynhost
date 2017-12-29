package main

import (
	"fmt"
	"log"
	"net"

	"github.com/parnurzeal/gorequest"
	"github.com/urfave/cli"
)

const ovhApiBaseUrl = "https://www.ovh.com/nic/update"

func UpdateRecord(c *cli.Context) error {
	if !c.Args().Present() {
		log.Fatal("The domain to update (e.g. host.example.org) should be provided as argument")
	}

	ip := getIpAddress(c)
	sendUpdateApi(config.Username, config.Password, c.Args().First(), ip)

	return nil
}

// OVH APIs
func sendUpdateApi(username string, password string, domain string, ip string) {
	req := gorequest.
		New().
		SetBasicAuth(username, password).
		Get(ovhApiBaseUrl).
		Param("system", "dyndns").
		Param("hostname", domain)

	if ip != "" {
		log.Printf("Passing the IP address %s for update of domain %s", ip, domain)
		req.Param("myip", ip)
	}

	log.Printf("Request to the OVH API: %s", req)

	resp, _, errs := req.End()
	log.Printf("Response from update record call to the OVH API: %s", resp)
	if errs != nil {
		log.Fatalf("Error while issuing the update record call to the OVH API: %s", errs)
	}
	if resp.StatusCode != 200 {
		log.Fatalf("Error while issuing the update record call to the OVH API: %s", resp)
	}

	fmt.Printf("IP Address of domain %s updated!\n", domain)
}

// IP/Network Interface management
func getIpAddress(c *cli.Context) string {
	if config.IpAddress != "" {
		log.Printf("An IP address was provided (%s), bypassing auto-detection.", config.IpAddress)
		return config.IpAddress
	}

	if config.NetworkInterface != "" {
		log.Printf("A network interface was provided (%s), using it for IP detection", config.NetworkInterface)
		iface, err := net.InterfaceByName(config.NetworkInterface)
		if err != nil {
			log.Fatalf("Unable to find the network interface %s: %s", config.NetworkInterface, err)
		}

		return getIpAddressOfInterface(iface)
	}

	log.Println("No interface or IP address specified: will use OVH autodetection")
	return ""
}

func getIpAddressOfInterface(iface *net.Interface) string {
	addrs, err := iface.Addrs()
	if err != nil {
		log.Fatalf("Unable to get addresses for the network interface %s: %s", iface.Name, err)
	}
	if len(addrs) == 0 {
		log.Fatalf("The network interface %s has no addresses", iface.Name)
	}

	ip, _, err := net.ParseCIDR(addrs[0].String())
	if len(addrs) == 0 {
		log.Fatalf("Unable to parse the IP (%s) from the network interface %s: %s", addrs[0].String(), iface.Name, err)
	}
	log.Printf("Returning IP address %s", ip.String())

	return ip.String()
}
