<h1 align="center">
  <img src="images/cachex-logo.png" alt="cachex" width="100px">
  <br>
</h1>

<h3 align="center">A high-accuracy, behavioral cache poisoning scanner for modern Web APIs</h3>

<p align="center">
  <img src="https://img.shields.io/badge/cacheX-blueviolet?style=flat-square">
  <img src="https://img.shields.io/github/go-mod/go-version/ayuxsec/cachex?style=flat-square">
  <img src="https://img.shields.io/github/license/ayuxsec/cachex?style=flat-square">
</p>


## 🌟 Why CacheX?

Most cache poisoning scanners only check:

* whether a response changes with certain headers
* or whether cache-related headers exist

This produces **tons of false positives** and rarely confirms a real exploit.

**CacheX is different.**

It performs **behavioral diffing**, **multi-threaded poisoning**, and **persistence verification**, confirming only real, weaponizable cache poisoning.

[![demo](https://asciinema.org/a/0t6ga94iTdGmMuCP99KQsAZDs.svg)](https://asciinema.org/a/0t6ga94iTdGmMuCP99KQsAZDs)

## 🔥 Features

* ⚡ **High-speed multi-threaded scanning**
* 🎯 **Zero-FP design with behavioral diffing**
* 🔁 **Real-time cache poisoning attempts**
* 🧪 **Persistence confirmation for true vulnerabilities**
* 🔍 **Single and multi-header scan modes**
* 🧩 **YAML-based payload configuration**
* 📤 **JSON or pretty output formats**
* 📁 **Optional file-based export**
* 🏷 **Tentative vs confirmed vuln tagging**

## 🔧 Installation

```bash
go install github.com/ayuxsec/cachex/cmd/cachex@latest
```

Or build manually:

```bash
git clone https://github.com/ayuxsec/cachex
cd cachex
go build -o cachex "cmd/cachex/main.go"
./cachex -h
```

## 🚀 Usage

### ▶️ Scan a single URL

```bash
cachex -u https://example.com
```

### ▶️ Scan multiple targets

```bash
cachex -l urls.txt
```

### ▶️ Scan URLs via pipeline

```bash
echo "https://example.com" | cachex
```

or:

```bash
cat urls.txt | cachex
```

---

## 📌 All CLI Flags

| Category          | Flag              | Description                 |
| ----------------- | ----------------- | --------------------------- |
| Input             | `-u, --url`       | URL to scan                 |
|                   | `-l, --list`      | File with list of URLs      |
| Concurrency       | `-t, --threads`   | Number of scanning threads  |
|                   | `-m, --scan-mode` | `single` or `multi`         |
| HTTP Client       | `--timeout`       | Total request timeout       |
|                   | `--proxy`         | Proxy URL                   |
| Persistence Check | `--no-chk-prst`   | Disable persistence checker |
|                   | `--prst-requests` | Poisoning requests          |
|                   | `--prst-threads`  | Threads for poisoning       |
| Output            | `-o, --output`    | Output file                 |
|                   | `-j, --json`      | JSON output                 |
| Payloads          | `--pcf`           | Custom payload config file  |


## 💡 Example

```bash
cachex -l targets.txt -t 50 --pcf payloads.yaml --json -o results.json
```

## ⚙️ Configuration

CacheX automatically loads:

```
~/.config/cachex/config.yaml
~/.config/cachex/payloads.yaml
```

You can configure:

* Payload headers
* Default request headers
* Timeouts & concurrency
* Logging mode
* Proxy settings
* Persistence checker behavior

## 📝 Output Formats

### Pretty Output

```
[vuln] [https://target.com] [Location Poisoning] [header: X-Forwarded-Host: evil.com] [poc: https://target.com?cache=XYZ]
```

### JSON Output

```json
{
  "URL": "https://target.com/",
  "IsVulnerable": true,
  "IsResponseManipulable": true,
  "ManipulationType": "ChangedBody",
  "RequestHeaders": {
    "Accept": "*/*",
    "User-Agent": "Mozilla/5.0"
  },
  "PayloadHeaders": {
    "X-Forwarded-Host": "evil.com"
  },
  "OriginalResponse": {
    "StatusCode": 200,
    "Headers": {
      "...": "..."
    },
    "Body": "...",
    "Location": ""
  },
  "ModifiedResponse": {
    "StatusCode": 200,
    "Headers": {
      "...": "..."
    },
    "Body": "...",
    "Location": ""
  },
  "PersistenceCheckResult": {
    "IsPersistent": true,
    "PoCLink": "https://target.example.com/?cache=XYZ",
    "FinalResponse": {
      "StatusCode": 200,
      "Headers": {
        "...": "..."
      },
      "Body": "...",
      "Location": ""
    }
  }
}
```

## 🎛 Scan Modes

* `single`: precise, tests each header independently
* `multi`: fast, tests all payload headers together

## 🧩 Payload Headers

Defined in:

```
~/.config/cachex/payloads.yaml
```

Example:

```yaml
payload_headers:
    X-Forwarded-Host: evil.com
    X-Forwarded-For: 127.0.0.1
    X-Original-URL: /evilpath
    X-Client-IP: 127.0.0.1
```

## 📁 Configuration File Example (`config.yaml`)

```yaml
scan_mode: single
threads: 25

request_headers:
  Accept: '*/*'
  User-Agent: Mozilla/5.0 (...)

client:
  dial_timeout: 5
  handshake_timeout: 5
  response_timeout: 10
  proxy_url: ""

persistence_checker:
  enabled: true
  num_requests_to_send: 10
  threads: 5

logger:
  log_error: false
  log_mode: pretty
  debug: false
  output_file: ""
  skip_tentative: true
```

## 🧠 How CacheX Works

1. Fetches baseline response
2. Injects payload headers
3. Detects response manipulation (body, code, redirect)
4. If changed → launches concurrent poisoning attempts
5. Fetches clean requests
6. If poisoned response persists → confirmed vulnerability
7. Outputs PoC link

## 📁 Project Structure

```console
cachex/
├── cmd/
│   └── cachex/
│       └── main.go                # CLI entrypoint
│
├── internal/
│   ├── app/
│   │   └── cachex/
│   │       └── cmd/
│   │           ├── banner.go      # ASCII banner
│   │           ├── flags.go       # CLI flags + config binding
│   │           ├── helper.go      # Help message builder
│           ├── root.go        # Main CLI logic & runner
│           └── utils.go           # File helpers
│
│   ├── pkg/
│   │   ├── client/
│   │   │   ├── client.go          # Custom HTTP client & transport
│   │   │   └── request.go         # Fetch + send raw requests
│   │   ├── config/
│   │   │   └── config.go          # Legacy internal config
│   │   └── logger/
│   │       ├── colors.go          # Color themes
│   │       └── logger.go          # Pretty logger (info/warn/debug/vuln)
│
│   └── scanner/
│       ├── core.go                # Core poisoning test logic
│       ├── detector.go            # Behavioral response diffing
│       ├── logger.go              # Pretty + JSON output formatter
│       ├── output.go              # JSON serialization helpers
│       ├── persistchk.go          # Persistence checker (real-time poisoning)
│       ├── scanner.go             # Scan controller (single/multi mode)
│       ├── types.go               # All scanner structs & enums
│       └── utils.go               # Cache buster, merging maps, helpers
│
├── pkg/
│   └── cachex/
│       ├── scanner.go             # Public API wrapper for internal scanner
│       ├── utils.go               # Config mappers (log mode, scan mode)
│       └── validate.go            # Config validation
│
│   └── config/
│       ├── config.go              # YAML config schema
│       ├── default.go             # Default paths + default config
│       └── payloads.go            # Default payload headers
│
├── .github/workflows/
│   └── release.yml                # Automated builds via GoReleaser
│
├── images/
│   ├── cachex-logo.png            # Logo
│   └── cachex-demo.gif            # Showcase GIF
│
├── .goreleaser.yaml               # Multi-platform binary releases
├── .gitignore
├── go.mod
├── go.sum
├── LICENSE
└── Makefile                       # Build / install helpers
```

## 🤝 Contribute

Sure, PRs are welcome!

## 📜 License

MIT © [@ayuxsec](https://github.com/ayuxsec)
