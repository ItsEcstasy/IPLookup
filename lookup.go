package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
)

func main() {
	Lookup()
	os.Exit(0)
}

type IPInfo struct {
	IP           string `json:"query"`
	Hostname     string `json:"reverse"`
	Organization string `json:"org"`
	City         string `json:"city"`
	Region       string `json:"regionName"`
	Country      string `json:"country"`
	Zip          string `json:"zip"`
}

func Lookup() {
	if len(os.Args) < 2 {
		fmt.Println("lookup <ip>")
		return
	}

	ip := os.Args[1]
	hostnames, err := net.LookupAddr(ip)
	if err != nil {
		fmt.Println("\x1b[93mError\x1b[97m: \x1b[93m", err)
		return
	}

	fmt.Println("\x1b[93mIP\x1b[97m: ", ip)
	for _, hostname := range hostnames {
		fmt.Println("\x1b[93mHostname\x1b[97m: ", hostname)
	}

	info, err := getIPInfo(ip)
	if err != nil {
		fmt.Println("\x1b[93mError\x1b[91m:", err)
		return
	}

	fmt.Println("\x1b[93mOrganization\x1b[97m: ", info.Organization)
	fmt.Println("\x1b[93mCity\x1b[97m: ", info.City)
	fmt.Println("\x1b[93mRegion\x1b[97m: ", info.Region)
	fmt.Println("\x1b[93mCountry\x1b[97m: ", info.Country)
	fmt.Println("\x1b[93mZip\x1b[97m: ", info.Zip)
}

func getIPInfo(ip string) (*IPInfo, error) {
	url := fmt.Sprintf("http://ip-api.com/json/%s", ip)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var info IPInfo
	err = json.Unmarshal(body, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}
