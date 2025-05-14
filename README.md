# go-firewalld  

[![Go Reference](https://pkg.go.dev/badge/gitee.com/weidongkl/go-firewalld.svg)](https://pkg.go.dev/gitee.com/weidongkl/go-firewalld)  
A Go library for managing **firewalld** dynamically via D-Bus, supporting zones, services, port forwarding, and rich rules.  

---

## Features  

✅ **D-Bus Integration**  
   - Interact with `firewalld` programmatically using Go.  
   - No need for shell commands or manual config edits.  

✅ **Firewall Management**  
   - **Zones**: Configure default/public/trusted zones.  
   - **Services**: Enable/disable predefined services (e.g., HTTP, SSH).  
   - **Ports**: Open/close ports with TCP/UDP support.  
   - **Rich Rules**: Define complex rules (e.g., source IP, logging).  
   - **Port Forwarding**: Set up forwarding between ports/interfaces.  

✅ **Lightweight & Efficient**  
   - Pure Go implementation (no CGO dependencies).  
   - Minimal overhead for cloud/container environments.  

---

## Installation  

```sh
go get gitee.com/weidongkl/go-firewalld
```

**Prerequisites**:  
- Linux system with `firewalld` installed and running.  
- Go 1.16+ (tested on modern Linux distributions).  

---

## Quick Start  

```go
package main

import (
	"gitee.com/weidongkl/go-firewalld"
	"log"
)

func main() {
	client, err := firewalld.NewClient(&firewalld.Options{})
	if err != nil {
		log.Fatalf("NewClient failed: %s", err)
	}
	log.Println("version: ", firewalld.Version())
	zone, _ := client.GetDefaultZone()
	log.Println("default zone: ", zone)
}

```

## Documentation  

Full API reference:  
- [GoDoc](https://pkg.go.dev/gitee.com/weidongkl/go-firewalld)  
- See [`examples/`](https://gitee.com/weidongkl/go-firewalld/tree/master/examples) for more use cases.  

---

## Contributing  

Pull requests and issues are welcome!  
1. Fork the repository.  
2. Test changes with `go test`.  
3. Ensure compatibility with major Linux distros (CentOS, Fedora, RHEL).  

---

## License  

MIT License. See [LICENSE](https://gitee.com/weidongkl/go-firewalld/blob/master/LICENSE).  

---

## Why Use This Library?  
- **Cloud-Native**: Ideal for dynamic firewall management in orchestrated environments.  
- **DevOps-Friendly**: Replace error-prone shell scripts with type-safe Go code.  
- **Performance**: Low-latency D-Bus calls compared to CLI alternatives.  

**Note**: Requires `firewalld` D-Bus API (default on CentOS/RHEL/Fedora).  
