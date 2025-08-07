package main

import (
	"errors"
	"fmt"
	"iplocator/core"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/iplocate/go-iplocate"
)

var apiClient *iplocate.Client

func main() {
	client, err := core.NewClient()
	if err != nil {
		log.Fatal("Failed to create API client:", err)
	}
	apiClient = client

	// Create Gin router
	r := gin.Default()

	// Load HTML templates
	r.LoadHTMLGlob("templates/*")

	// Serve static files (CSS, JS)
	r.Static("/static", "./static")

	// Routes
	r.GET("/", homePage)
	r.POST("/lookup", lookupHandler)
	r.POST("/self-lookup", selfLookupHandler)

	// Start server
	log.Println("Server starting on http://localhost:8080")
	r.Run(":8080")
}

func homePage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "IP Locator",
	})
}

func lookupHandler(c *gin.Context) {
	ipList, err := validateAndParseInput(c)
	if err != nil {
		c.HTML(http.StatusOK, "error.html", gin.H{
			"error": err.Error(),
		})
		return
	}

	publicIPs, privateIPs, err := checkPrivateIPs(ipList)
	if err != nil {
		c.HTML(http.StatusOK, "error.html", gin.H{
			"error": err.Error(),
		})
		return
	}

	validResults, errorMessages, err := processLookupResults(publicIPs)
	if err != nil {
		c.HTML(http.StatusOK, "error.html", gin.H{
			"error": err.Error(),
		})
		return
	}

	showResults(c, validResults, errorMessages, privateIPs)
}

func validateAndParseInput(c *gin.Context) ([]string, error) {
	ipInput := strings.TrimSpace(c.PostForm("ip"))
	if ipInput == "" {
		return nil, errors.New("please enter one or more IP addresses")
	}

	ipList := strings.Fields(ipInput)
	if err := core.ValidateIPs(ipList); err != nil {
		return nil, err
	}

	return ipList, nil
}

func checkPrivateIPs(ipList []string) (publicIPs []string, privateIPs []string, err error) {
	publicIPs, privateIPs = core.SeparatePublicAndPrivateIPs(ipList)

	if len(publicIPs) == 0 {
		errorMsg := "all provided IP addresses are private/local. Please enter public IP addresses for geolocation lookup"
		return nil, nil, errors.New(errorMsg)
	}

	return publicIPs, privateIPs, nil
}

func processLookupResults(publicIPs []string) (validResults []*iplocate.LookupResponse, errorMessages []string, err error) {
	results, errs := core.LookupIPs(apiClient, publicIPs)

	for i, lookupErr := range errs {
		if lookupErr != nil {
			errorMessages = append(errorMessages, fmt.Sprintf("Error looking up %s: %s", publicIPs[i], lookupErr.Error()))
		}
	}

	for _, result := range results {
		if result != nil {
			validResults = append(validResults, result)
		}
	}

	if len(validResults) == 0 {
		errorMsg := "No results found for any IP addresses"
		if len(errorMessages) > 0 {
			errorMsg += ":\n" + strings.Join(errorMessages, "\n")
		}
		return nil, nil, errors.New(errorMsg)
	}

	return validResults, errorMessages, nil
}

func showResults(c *gin.Context, validResults []*iplocate.LookupResponse, errorMessages []string, privateIPs []string) {
	var formattedResults []string
	for _, result := range validResults {
		formattedResults = append(formattedResults, core.FormatResult(result))
	}

	c.HTML(http.StatusOK, "results.html", gin.H{
		"formatted_results": formattedResults,
		"count":             len(validResults),
		"errors":            errorMessages,
		"private_ips":       privateIPs,
	})
}

func selfLookupHandler(c *gin.Context) {
	result, err := core.LookupSelf(apiClient)
	if err != nil {
		c.HTML(http.StatusOK, "error.html", gin.H{
			"error": "Self lookup failed: " + err.Error(),
		})
		return
	}

	if result == nil {
		c.HTML(http.StatusOK, "error.html", gin.H{
			"error": "No result returned from self lookup",
		})
		return
	}

	showResults(c, []*iplocate.LookupResponse{result}, []string{}, []string{})
}
