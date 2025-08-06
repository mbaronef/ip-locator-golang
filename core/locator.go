package core

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"

	"github.com/iplocate/go-iplocate"
)

func ValidateIPs(ips []string) error {
	for _, ip := range ips {
		if net.ParseIP(ip) == nil {
			return fmt.Errorf("invalid IP format: %s", ip)
		}
	}
	return nil
}

func LookupSelf(client *iplocate.Client) (*iplocate.LookupResponse, error) {
	return client.LookupSelf()
}

func LookupIPs(client *iplocate.Client, ips []string) ([]*iplocate.LookupResponse, []error) {
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

	return results, errs
}

func ReadIPsFromFile(path string) ([]string, error) {
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
