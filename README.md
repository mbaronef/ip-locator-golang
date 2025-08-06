# IP Locator - Desktop GUI

A user-friendly desktop application for IP geolocation lookups built with Go and Fyne. Get detailed information about any IP address with a clean, native GUI interface.

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
-  **Batch Processing**: Load and process multiple IPs from files
- � **Export Results**: Save results to JSON or CSV files
- ⚡ **Fast Processing**: Concurrent IP lookups with progress indicators
- 🔍 **Self Lookup**: Check your own public IP information
- 🎨 **Modern UI**: Clean, intuitive interface built with Fyne

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
- **Batch Processing**: Use "Load File" to process multiple IPs from a text file
- **Self Lookup**: Click "My IP" to check your own public IP
- **Export Results**: Save results as JSON or CSV files
- **Copy to Clipboard**: Click any result field to copy to clipboard

## Learning Fyne

### Quick Start Resources:
- [Official Fyne Documentation](https://docs.fyne.io/)
- [Fyne Tutorial](https://docs.fyne.io/started/)
- [Widget Tour](https://docs.fyne.io/widget/)

### Useful Fyne Concepts:
- **Widgets**: UI components (buttons, entries, labels)
- **Containers**: Layout managers (border, grid, vbox, hbox)
- **Canvas**: Drawing surface and coordinate system
- **Themes**: Customizable appearance

## Dependencies

- [fyne.io/fyne/v2](https://fyne.io/) - Cross-platform GUI toolkit
- [iplocate/go-iplocate](https://github.com/iplocate/go-iplocate) - IPLocate API client



