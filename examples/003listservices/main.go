package main

import (
	"gitee.com/weidongkl/go-firewalld"
	"log"
)

func main() {
	client, err := firewalld.NewClient(&firewalld.Options{Permanent: true})
	if err != nil {
		log.Fatalf("NewClient failed: %s", err)
	}
	defer client.Close()
	services, err := client.ListServices()
	if err != nil {
		log.Fatalf("NewClient failed: %s", err)
	}
	log.Printf("%#v\n", services)
}
