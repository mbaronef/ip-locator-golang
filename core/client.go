package core

import (
    "fmt"
    "os"
    
    "github.com/iplocate/go-iplocate"
)

func NewClient() (*iplocate.Client, error) {
    apiKey := os.Getenv("IPLOCATE_API_KEY")
    if apiKey == "" {
        return nil, fmt.Errorf("missing IPLOCATE_API_KEY environment variable")
    }
    return iplocate.NewClient(nil).WithAPIKey(apiKey), nil
}
