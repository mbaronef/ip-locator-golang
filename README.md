# IP Locator

A fast and efficient command-line tool written in Go for IP geolocation lookups. Get detailed information about any IP address including location, ISP, timezone, and privacy detection (VPN/Proxy).

**Using the geolocation API provided by [IPLocate.io](https://iplocate.io)**


## Project Structure

This project uses a branch-based approach to separate different interfaces:

- 📋 **main** branch: **Pure CLI tool** - Fast, lightweight command-line interface
- 🌐 **ui/web** branch: **Web interface** - User-friendly browser-based UI
- 🖥️ **ui/desktop** branch: **Desktop application** - Native GUI application

## Features

- **Accurate Geolocation**: Get country, city, and coordinates for any IP address
- **ISP Information**: Retrieve ASN and provider details
- **Timezone Data**: Get timezone information for the IP location
- **Privacy Detection**: Detect VPN and proxy usage
- **Multiple Output Formats**: Support for both human-readable and JSON output
- **Batch Processing**: Process multiple IPs from a file
- **Concurrent Processing**: Fast parallel IP lookups
- **Self Lookup**: Check your own public IP information
- **Smart Filtering**: Automatically detects and warns about private IP addresses

## Installation

### Prerequisites

- Go 1.19 or higher
- IPLocate API key (get one at [iplocate.io](https://iplocate.io))

### Build from Source

```bash
git clone https://github.com/mbaronef/ip-locator-golang.git
cd ip-locator-golang
go mod tidy
go build -o iplocator main.go
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

### Basic IP Lookup

```bash
# Lookup a single IP
./iplocator 8.8.8.8

# Lookup your own IP
./iplocator --self
```

### Multiple IPs

```bash
# Lookup multiple IPs
./iplocator 8.8.8.8 1.1.1.1 208.67.222.222

# Mixed public and private IPs (private IPs will be skipped with warning)
./iplocator 8.8.8.8 192.168.1.1 1.1.1.1
```

### File Input

```bash
# Process IPs from a file (one IP per line)
./iplocator --file example.txt
```

### JSON Output

```bash
# Get output in JSON format
./iplocator --json 8.8.8.8
./iplocator --json --file example.txt
```

## Command Line Options

| Flag | Description |
|------|-------------|
| `--json` | Output results in JSON format |
| `--self` | Lookup your own public IP address |
| `--file <path>` | Read IP addresses from a file (one per line) |
| `--help` | Show help information |

## Example Output

### Standard Output
```
IP: 8.8.8.8
Country: United States (US)
City: Mountain View
Coordinates: 37.4056, -122.0775
Time Zone: America/Los_Angeles
ISP: Google LLC (ASN AS15169)
---------------
```

### JSON Output
```json
[
  {
    "ip": "8.8.8.8",
    "country": "United States",
    "country_code": "US",
    "city": "Mountain View",
    "latitude": 37.4056,
    "longitude": -122.0775,
    "time_zone": "America/Los_Angeles",
    "asn": {
      "asn": "AS15169",
      "name": "Google LLC"
    },
    "privacy": {
      "vpn": false,
      "proxy": false
    }
  }
]
```

## Dependencies

- [iplocate/go-iplocate](https://github.com/iplocate/go-iplocate) - IPLocate API client
- [fatih/color](https://github.com/fatih/color) - Colored terminal output
