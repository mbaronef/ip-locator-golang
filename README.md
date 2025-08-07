# IP Locator - Desktop GUI

A user-friendly desktop application for IP geolocation lookups built with Go and Fyne. Get detailed information about any IP address with a clean, native GUI.

**Using the geolocation API provided by [IPLocate.io](https://iplocate.io)**

## Project Structure

This project uses a branch-based approach to separate different interfaces:

- 📋 **main** branch: **Pure CLI tool** - Fast, lightweight command-line interface
- 🌐 **ui/web** branch: **Web interface** - User-friendly browser-based UI  
- 🖥️ **ui/desktop** branch: **Desktop application** - Native GUI application

## Features

- 🖥️ **Native GUI**: Cross-platform desktop application with native look and feel
- 🌍 **Accurate Geolocation**: Get country, city, and coordinates for any IP address
- 🏢 **ISP Information**: Retrieve ASN and provider details
- 🕒 **Timezone Data**: Get timezone information for the IP location
- 🔒 **Privacy Detection**: Detect VPN and proxy usage
- 📂 **File Processing**: Load and process multiple IPs from text files
- 💾 **Export Results**: Save results to JSON or readable text formats
- 🚀 **Concurrent Processing**: Fast parallel IP lookups
- 🔍 **Self Lookup**: Check your own public IP information
- ⚠️ **Smart Filtering**: Automatically detects and warns about private IP addresses

## Installation

### Prerequisites

- Go 1.19 or higher
- IPLocate API key (get one at [iplocate.io](https://iplocate.io))

### Platform-specific Requirements

**Windows:**
- No additional requirements

**macOS:**
- Xcode command line tools: `xcode-select --install`

**Linux:**
- Development packages: `sudo apt install gcc libc6-dev libgl1-mesa-dev xorg-dev` (Ubuntu/Debian)

### Build from Source

```bash
git clone https://github.com/mbaronef/ip-locator-golang.git
cd ip-locator-golang
git checkout ui/desktop
go mod tidy
go build -o iplocator-gui .
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

### Running the Application

```bash
# Run the desktop GUI
./iplocator-gui

# Or directly with Go
go run .
```

### App Features

- **Single IP Lookup**: Enter an IP address and click "Lookup"
- **Self Lookup**: Click "Self Lookup" to check your own public IP information
- **File Upload**: Use "Upload File" to process multiple IPs from a text file (one IP address per line)
- **Smart IP Filtering**: Automatically detects private IPs with user-friendly dialog warnings
- **Card-based Results**: Each IP result displays in its own organized card
- **Export Options**: Save results as JSON or readable text files
- **Scrollable Interface**: Browse through multiple results easily

## Dependencies

- [fyne.io/fyne/v2](https://fyne.io/) - Cross-platform GUI toolkit
- [iplocate/go-iplocate](https://github.com/iplocate/go-iplocate) - IPLocate API client



