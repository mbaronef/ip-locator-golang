package core

import (
	"fmt"
	"net"

	"github.com/iplocate/go-iplocate"
)

func FormatResult(result *iplocate.LookupResponse) string {
	var out string

	out += fmt.Sprintf("✅ IP: %s\n\n", result.IP)
	if result.Country != nil {
		out += fmt.Sprintf("🌍 Country: %s (%s)\n", *result.Country, *result.CountryCode)
	}
	if result.City != nil {
		out += fmt.Sprintf("🏙️ City: %s\n", *result.City)
	}
	if result.Latitude != nil && result.Longitude != nil {
		out += fmt.Sprintf("📍 Coordinates: %.4f, %.4f\n", *result.Latitude, *result.Longitude)
	}
	if result.TimeZone != nil {
		out += fmt.Sprintf("🕒 Time Zone: %s\n", *result.TimeZone)
	}
	if result.ASN != nil {
		out += fmt.Sprintf("🏢 ISP: %s (ASN %s)\n", result.ASN.Name, result.ASN.ASN)
	}
	if result.Privacy.IsVPN {
		out += "🔒 Uses VPN: Yes\n"
	}
	if result.Privacy.IsProxy {
		out += "🔒 Uses Proxy: Yes\n"
	}
	if !result.Privacy.IsVPN && !result.Privacy.IsProxy {
		out += "🔓 Privacy: No VPN or Proxy detected\n"
	}
	return out
}

func IsPrivateIP(ipStr string) bool {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return false
	}

	// IPv4 private ranges
	if ip.To4() != nil {
		// 10.0.0.0/8
		if ip[12] == 10 {
			return true
		}
		// 172.16.0.0/12
		if ip[12] == 172 && ip[13] >= 16 && ip[13] <= 31 {
			return true
		}
		// 192.168.0.0/16
		if ip[12] == 192 && ip[13] == 168 {
			return true
		}
		// 127.0.0.0/8 (localhost)
		if ip[12] == 127 {
			return true
		}
	}

	// Check for IPv6 private ranges
	if ip.IsLoopback() || ip.IsLinkLocalUnicast() {
		return true
	}

	return false
}

func SeparatePublicAndPrivateIPs(ips []string) (publicIPs []string, privateIPs []string) {
	for _, ip := range ips {
		if IsPrivateIP(ip) {
			privateIPs = append(privateIPs, ip)
		} else {
			publicIPs = append(publicIPs, ip)
		}
	}
	return publicIPs, privateIPs
}
