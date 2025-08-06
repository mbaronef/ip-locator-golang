package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"sync"

	"github.com/fatih/color"
	"github.com/iplocate/go-iplocate"
)

type Config struct {
	JSONOutput bool
	SelfLookup bool
	FilePath   string
	IPs        []string
}

func main() {
	config := parseFlags()

	client, err := newClient()
	if err != nil {
		log.Fatal(err)
	}

	results, err := performLookups(config, client)
	if err != nil {
		log.Fatal(err)
	}

	if config.JSONOutput {
		if err := printJSON(results); err != nil {
			log.Fatal(err)
		}
	} else {
		printResultsList(results)
	}
}

func parseFlags() Config {
	flag.Usage = func() {
		fmt.Println("Usage: iplocator [options] [IP]")
		fmt.Println("\nOptions:")
		flag.PrintDefaults()
	}

	jsonOutput := flag.Bool("json", false, "Show output in JSON format")
	selfLookup := flag.Bool("self", false, "Lookup your own IP")
	filePath := flag.String("file", "", "Path to a file with IPs (one per line)")
	flag.Parse()

	return Config{
		JSONOutput: *jsonOutput,
		SelfLookup: *selfLookup,
		FilePath:   *filePath,
		IPs:        flag.Args(),
	}
}

func newClient() (*iplocate.Client, error) {
	apiKey := os.Getenv("IPLOCATE_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("environment variable IPLOCATE_API_KEY is not set")
	}
	return iplocate.NewClient(nil).WithAPIKey(apiKey), nil
}

func performLookups(config Config, client *iplocate.Client) ([]*iplocate.LookupResponse, error) {
	if config.SelfLookup {
		if len(config.IPs) > 0 || config.FilePath != "" {
			return nil, fmt.Errorf("--self cannot be used with other IPs or --file")
		}
		result, err := client.LookupSelf()
		if err != nil {
			return nil, err
		}
		return []*iplocate.LookupResponse{result}, nil
	}

	ips, err := collectIPs(config)
	if err != nil {
		return nil, err
	}

	if err := validateIPs(ips); err != nil {
		return nil, err
	}

	return lookupIPs(client, ips), nil
}

func collectIPs(config Config) ([]string, error) {
	var ips []string

	if config.FilePath != "" {
		fileIPs, err := readIPsFromFile(config.FilePath)
		if err != nil {
			return nil, err
		}
		ips = append(ips, fileIPs...)
	}

	if len(config.IPs) > 0 {
		ips = append(ips, config.IPs...)
	}

	if len(ips) == 0 {
		ip := promptForIP()
		ips = append(ips, ip)
	}

	return ips, nil
}

func readIPsFromFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}
	defer file.Close()

	var ips []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			ips = append(ips, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}
	return ips, nil
}

func promptForIP() string {
	var ip string
	fmt.Print("Enter an IP: ")
	fmt.Scanln(&ip)
	return ip
}

func validateIPs(ips []string) error {
	for _, ip := range ips {
		if net.ParseIP(ip) == nil {
			return fmt.Errorf("invalid IP format: %s", ip)
		}
	}
	return nil
}

func lookupIPs(client *iplocate.Client, ips []string) []*iplocate.LookupResponse {
	var wg sync.WaitGroup
	results := make([]*iplocate.LookupResponse, len(ips))
	errs := make([]error, len(ips))

	for i, ip := range ips {
		wg.Add(1)
		go func(i int, ip string) {
			defer wg.Done()
			res, err := client.Lookup(ip)
			if err != nil {
				errs[i] = err
				return
			}
			results[i] = res
		}(i, ip)
	}
	wg.Wait()

	for i, err := range errs {
		if err != nil {
			log.Printf("Error with %s: %v", ips[i], err)
		}
	}

	return results
}

func printResultsList(results []*iplocate.LookupResponse) {
	for _, r := range results {
		printSingleResult(r)
		fmt.Println("---------------")
	}
}

func printSingleResult(result *iplocate.LookupResponse) {
	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()

	fmt.Printf("%s: %s\n", green("IP"), result.IP)
	if result.Country != nil {
		fmt.Printf("%s: %s (%s)\n", green("Country"), *result.Country, *result.CountryCode)
	}
	if result.City != nil {
		fmt.Printf("%s: %s\n", green("City"), *result.City)
	}
	if result.Latitude != nil && result.Longitude != nil {
		fmt.Printf("%s: %.4f, %.4f\n", green("Coordinates"), *result.Latitude, *result.Longitude)
	}
	if result.TimeZone != nil {
		fmt.Printf("%s: %s\n", green("Time Zone"), *result.TimeZone)
	}
	if result.ASN != nil {
		fmt.Printf("%s: %s (ASN %s)\n", green("ISP"), result.ASN.Name, result.ASN.ASN)
	}
	if result.Privacy.IsVPN {
		fmt.Printf("%s\n", red("Uses VPN"))
	}
	if result.Privacy.IsProxy {
		fmt.Printf("%s\n", red("Uses Proxy"))
	}
}

func printJSON(results []*iplocate.LookupResponse) error {
	jsonData, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(jsonData))
	return nil
}
