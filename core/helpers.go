package core

import (
    "encoding/json"
    "fmt"
    
    "github.com/fatih/color"
    "github.com/iplocate/go-iplocate"
)

func FormatResult(result *iplocate.LookupResponse) string {
    var out string
    
    green := color.New(color.FgGreen).SprintFunc()
    red := color.New(color.FgRed).SprintFunc()

    out += fmt.Sprintf("%s: %s\n", green("IP"), result.IP)
    if result.Country != nil {
        out += fmt.Sprintf("%s: %s (%s)\n", green("Country"), *result.Country, *result.CountryCode)
    }
    if result.City != nil {
        out += fmt.Sprintf("%s: %s\n", green("City"), *result.City)
    }
    if result.Latitude != nil && result.Longitude != nil {
        out += fmt.Sprintf("%s: %.4f, %.4f\n", green("Coordinates"), *result.Latitude, *result.Longitude)
    }
    if result.TimeZone != nil {
        out += fmt.Sprintf("%s: %s\n", green("Time Zone"), *result.TimeZone)
    }
    if result.ASN != nil {
        out += fmt.Sprintf("%s: %s (ASN %s)\n", green("ISP"), result.ASN.Name, result.ASN.ASN)
    }
    if result.Privacy.IsVPN {
        out += fmt.Sprintf("%s\n", red("Uses VPN"))
    }
    if result.Privacy.IsProxy {
        out += fmt.Sprintf("%s\n", red("Uses Proxy"))
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
