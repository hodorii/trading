---
name: "curl-cli-guide"
displayName: "cURL CLI Guide"
description: "Complete guide for using cURL command-line tool for HTTP requests, file transfers, and API testing with practical examples and troubleshooting."
keywords: ["curl", "http", "api", "cli", "requests", "download", "upload"]
author: "Kiro User"
---

# cURL CLI Guide

## Overview

cURL is a powerful command-line tool for transferring data with URLs. It supports numerous protocols including HTTP, HTTPS, FTP, and more. This guide covers installation, common usage patterns, and troubleshooting for everyday HTTP requests and API interactions.

Whether you're testing APIs, downloading files, or debugging web services, cURL provides the flexibility and control you need from the command line.

## Onboarding

### Installation

#### Windows
```bash
# cURL comes pre-installed on Windows 10+ 
curl --version

# If not available, install via chocolatey
choco install curl

# Or download from official site
# https://curl.se/windows/
```

#### macOS
```bash
# Pre-installed on macOS
curl --version

# Or install latest via Homebrew
brew install curl
```

#### Linux (Ubuntu/Debian)
```bash
# Usually pre-installed
curl --version

# If not available
sudo apt update
sudo apt install curl
```

### Prerequisites
- Command line access (Terminal, PowerShell, or Command Prompt)
- Basic understanding of HTTP methods (GET, POST, PUT, DELETE)
- No additional dependencies required

### Verification
```bash
# Verify installation
curl --version

# Expected output:
# curl 7.x.x (platform) libcurl/7.x.x ...
# Protocols: http https ftp ftps ...
```

## Common Workflows

### Workflow: Basic HTTP Requests

**Goal:** Make simple HTTP requests to APIs and websites

**Commands:**
```bash
# GET request (default)
curl https://api.example.com/users

# GET with headers
curl -H "Accept: application/json" https://api.example.com/users

# POST request with data
curl -X POST -H "Content-Type: application/json" \
     -d '{"name":"John","email":"john@example.com"}' \
     https://api.example.com/users

# PUT request
curl -X PUT -H "Content-Type: application/json" \
     -d '{"name":"John Updated"}' \
     https://api.example.com/users/1

# DELETE request
curl -X DELETE https://api.example.com/users/1
```

**Complete Example:**
```bash
# Test a REST API endpoint
curl -X GET \
     -H "Accept: application/json" \
     -H "Authorization: Bearer your-token-here" \
     https://jsonplaceholder.typicode.com/posts/1

# Expected response:
# {
#   "userId": 1,
#   "id": 1,
#   "title": "sunt aut facere...",
#   "body": "quia et suscipit..."
# }
```

### Workflow: File Operations

**Goal:** Download and upload files using cURL

**Commands:**
```bash
# Download file
curl -O https://example.com/file.zip

# Download with custom filename
curl -o myfile.zip https://example.com/file.zip

# Download with progress bar
curl --progress-bar -O https://example.com/largefile.zip

# Upload file via POST
curl -X POST -F "file=@/path/to/file.txt" https://api.example.com/upload

# Upload with additional form data
curl -X POST \
     -F "file=@document.pdf" \
     -F "description=Important document" \
     https://api.example.com/upload
```

### Workflow: Authentication & Headers

**Goal:** Handle various authentication methods and custom headers

**Commands:**
```bash
# Basic authentication
curl -u username:password https://api.example.com/protected

# Bearer token
curl -H "Authorization: Bearer your-jwt-token" https://api.example.com/protected

# API key in header
curl -H "X-API-Key: your-api-key" https://api.example.com/data

# Multiple custom headers
curl -H "Accept: application/json" \
     -H "User-Agent: MyApp/1.0" \
     -H "X-Custom-Header: value" \
     https://api.example.com/endpoint
```

## Command Reference

### curl

**Purpose:** Transfer data from or to servers using various protocols

**Syntax:**
```bash
curl [options] [URL...]
```

**Common Options:**
| Flag | Description | Example |
|------|-------------|---------|
| `-X, --request` | HTTP method | `-X POST` |
| `-H, --header` | Add header | `-H "Content-Type: application/json"` |
| `-d, --data` | Send data in POST | `-d '{"key":"value"}'` |
| `-o, --output` | Write output to file | `-o filename.txt` |
| `-O, --remote-name` | Use remote filename | `-O` |
| `-u, --user` | Authentication | `-u user:pass` |
| `-v, --verbose` | Verbose output | `-v` |
| `-s, --silent` | Silent mode | `-s` |
| `-f, --fail` | Fail silently on errors | `-f` |
| `-L, --location` | Follow redirects | `-L` |

**Examples:**
```bash
# Verbose GET request
curl -v https://httpbin.org/get

# Silent POST with JSON
curl -s -X POST -H "Content-Type: application/json" \
     -d '{"test":"data"}' https://httpbin.org/post

# Follow redirects and save
curl -L -o page.html https://example.com/redirect-url
```

## Troubleshooting

### Error: "curl: command not found"
**Cause:** cURL is not installed or not in PATH
**Solution:**
1. Check if installed: `which curl` (Linux/macOS) or `where curl` (Windows)
2. Install using package manager (see Installation section)
3. Restart terminal after installation

### Error: "SSL certificate problem"
**Cause:** SSL certificate verification failed
**Solution:**
1. **Recommended:** Fix the certificate issue on the server
2. **For testing only:** Skip verification with `-k` flag:
   ```bash
   curl -k https://self-signed-site.com
   ```
3. Specify CA bundle:
   ```bash
   curl --cacert /path/to/ca-bundle.crt https://site.com
   ```

### Error: "Connection refused" or "Connection timeout"
**Cause:** Server is down, firewall blocking, or wrong URL
**Solution:**
1. Verify URL is correct
2. Check if server is running: `ping hostname`
3. Try with verbose flag: `curl -v https://example.com`
4. Check firewall settings
5. Test with different network/VPN

### Error: "HTTP 401 Unauthorized"
**Cause:** Missing or invalid authentication
**Solution:**
1. Verify credentials are correct
2. Check authentication method (Basic, Bearer, API key)
3. Ensure proper header format:
   ```bash
   # Correct Bearer token format
   curl -H "Authorization: Bearer your-token"
   
   # Correct Basic auth format
   curl -u username:password
   ```

### Error: "HTTP 400 Bad Request"
**Cause:** Malformed request data or headers
**Solution:**
1. Validate JSON syntax if sending JSON data
2. Check Content-Type header matches data format
3. Use verbose mode to see full request: `curl -v`
4. Test with simple request first, then add complexity

## Best Practices

- **Use verbose mode (`-v`) for debugging** - Shows full request/response headers
- **Always specify Content-Type** when sending data: `-H "Content-Type: application/json"`
- **Use proper HTTP methods** - GET for reading, POST for creating, PUT for updating, DELETE for removing
- **Handle errors gracefully** with `-f` flag to fail on HTTP errors
- **Follow redirects** with `-L` when needed
- **Store sensitive data in environment variables** instead of command line:
  ```bash
  curl -H "Authorization: Bearer $API_TOKEN" https://api.example.com
  ```
- **Use configuration files** for complex requests with many options
- **Test with simple requests first** before adding authentication and complex data

## Additional Resources

- Official Documentation: https://curl.se/docs/
- Manual Page: `man curl` or https://curl.se/docs/manpage.html
- HTTP Status Codes: https://httpstatuses.com/
- JSON Validator: https://jsonlint.com/

---

**CLI Tool:** `curl`
**Installation:** Pre-installed on most systems, or via package manager