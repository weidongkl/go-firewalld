# go-firewalld

firewalld go client

## Installation

```bash
go get gitee.com/weidongkl/go-firewalld
```

## Quickstart

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

For advanced usage, please refer to [example](examples)

## Contributing

1. Fork the repository.
2. Create a new branch for your feature.
3. Commit your changes.
4. Push the branch.
5. Create a pull request.

## License

This project is licensed under the MIT License.

## Contact

For any questions or issues, please contact weidongkx@gmail.com.

