# IP Locator - Web Interface

A user-friendly web interface for IP geolocation lookups built with Go, Gin, and HTMX. Get detailed information about any IP address through a responsive browser interface with real-time interactions.

**Using the geolocation API provided by [IPLocate.io](https://iplocate.io)**


## Project Structure

This project uses a branch-based approach to separate different interfaces:

- 📋 **main** branch: **Pure CLI tool** - Fast, lightweight command-line interface
- 🌐 **ui/web** branch: **Web interface** - User-friendly browser-based UI
- 🖥️ **ui/desktop** branch: **Desktop application** - Native GUI application

## Features

-  **Modern Web Interface**: Responsive design that works on desktop and mobile
-  **Accurate Geolocation**: Get country, city, and coordinates for any IP address
-  **ISP Information**: Retrieve ASN and provider details
-  **Timezone Data**: Get timezone information for the IP location
-  **Privacy Detection**: Detect VPN and proxy usage
-  **Export Results**: Save results to JSON or TXT formats
-  **Multiple IP Support**: Process multiple IP addresses at once (space-separated)
-  **Concurrent Processing**: Fast parallel IP lookups
-  **Self Lookup**: Check your own public IP information
-  **Smart Filtering**: Automatically detects and warns about private IP addresses

## Installation

### Prerequisites

- Go 1.19 or higher
- IPLocate API key (get one at [iplocate.io](https://iplocate.io))

### Build from Source

```bash
git clone https://github.com/mbaronef/ip-locator-golang.git
cd ip-locator-golang
git checkout ui/web
go mod tidy
```

### Set API Key

Set your IPLocate API key as an environment variable:

**Windows (PowerShell):**
```powershell
$env:IPLOCATE_API_KEY="your_api_key_here"
```

**Windows (Command Prompt):**
```cmd
set IPLOCATE_API_KEY=your_api_key_here
```

**Linux/macOS:**
```bash
export IPLOCATE_API_KEY="your_api_key_here"
```

## Usage

### Starting the Web Server

```bash
# Start the web application
go run main.go

# The application will be available at:
# http://localhost:8080
```

### Live Demo

🌐 **Try it online**: [https://ip-locator-golang.onrender.com](https://ip-locator-golang.onrender.com)
> *Note: Free deployment with cold starts (~30s). Demo uses a shared API key with rate limits. "Self Lookup" shows server location (Oregon, US), not your location. For unlimited usage, run locally with your own API key.*

### Web Interface Features

- **Single IP Lookup**: Enter an IP address in the main form and click "Lookup"
- **Multiple IPs**: Enter multiple IP addresses separated by spaces
- **Self Lookup**: Click "Self Lookup" to check your own public IP information
- **Download Results**: Export lookup results as JSON or TXT files with one click
- **Real-time Results**: Results appear instantly without page reloads
- **Private IP Detection**: Automatically warns when private/local IPs are entered
- **Error Handling**: Clear error messages for invalid inputs and failed lookups
- **Responsive Design**: Works seamlessly on desktop, tablet, and mobile devices

## Dependencies
- [gin-gonic/gin](https://github.com/gin-gonic/gin) - Fast HTTP web framework
- [iplocate/go-iplocate](https://github.com/iplocate/go-iplocate) - IPLocate API client
- [HTMX](https://htmx.org/) - Modern web interactions without complex JavaScript
