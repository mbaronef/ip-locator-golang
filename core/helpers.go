package core

import (
	"encoding/json"
	"fmt"

	"github.com/iplocate/go-iplocate"
)

func FormatResult(result *iplocate.LookupResponse) string {
	var out string

	out += fmt.Sprintf("IP: %s\n", result.IP)
	if result.Country != nil {
		out += fmt.Sprintf("Country: %s (%s)\n", *result.Country, *result.CountryCode)
	}
	if result.City != nil {
		out += fmt.Sprintf("City: %s\n", *result.City)
	}
	if result.Latitude != nil && result.Longitude != nil {
		out += fmt.Sprintf("Coordinates: %.4f, %.4f\n", *result.Latitude, *result.Longitude)
	}
	if result.TimeZone != nil {
		out += fmt.Sprintf("Time Zone: %s\n", *result.TimeZone)
	}
	if result.ASN != nil {
		out += fmt.Sprintf("ISP: %s (ASN %s)\n", result.ASN.Name, result.ASN.ASN)
	}
	if result.Privacy.IsVPN {
		out += "Uses VPN: Yes\n"
	}
	if result.Privacy.IsProxy {
		out += "Uses Proxy: Yes\n"
	}
	if !result.Privacy.IsVPN && !result.Privacy.IsProxy {
		out += "Privacy: No VPN or Proxy detected\n"
	}
	return out
}

func FormatJSON(results []*iplocate.LookupResponse) (string, error) {
	jsonData, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return "", fmt.Errorf("error formatting results as JSON: %v", err)
	}
	return string(jsonData), nil
}
