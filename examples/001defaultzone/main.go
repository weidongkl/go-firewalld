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
	log.Println("version: ", firewalld.Version())
	zone, _ := client.GetDefaultZone()
	log.Println("default zone: ", zone)
}
