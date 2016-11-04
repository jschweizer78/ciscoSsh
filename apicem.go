package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/jbogarin/go-apic-em/apic-em"
)

func main() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	apicEMClient := apicem.NewClient(client)
	user := &apicem.User{
		Username: "",
		Password: "",
	}
	apicEMClient.BaseURL, _ = url.Parse("<URL>")

	fmt.Println("==========================================")

	newTicket, resp, err := apicEMClient.Ticket.AddTicket(user)
	if err != nil {
		fmt.Println(resp.Response.Status)
		log.Fatal(err)
	}

	ticket := newTicket.Response.ServiceTicket
	fmt.Println("Ticket", ticket)
	apicEMClient.Authorization = ticket

	fmt.Println("==========================================")

	networkDevicesCount, resp, err := apicEMClient.NetworkDevice.GetNetworkDeviceCount("all")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(networkDevicesCount.Response)

	fmt.Println("==========================================")

	networkDevices, resp, err := apicEMClient.NetworkDevice.GetAllNetworkDevice("all")
	if err != nil {
		log.Fatal(err)
	}

	for _, networkDevice := range networkDevices.Response {
		fmt.Println(networkDevice.Hostname)
	}
	// for _, networkDevice := range networkDevices.Response {
	// 	fmt.Println(networkDevice.Hostname, networkDevice.SerialNumber)
	// }
}
