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
	defer client.Close()
	svcSet, err := client.GetServiceSettings("ssh")
	if err != nil {
		log.Fatalf("NewClient failed: %s", err)
	}
	log.Printf("%#v\n", svcSet)
}
