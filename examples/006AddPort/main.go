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
	err = client.AddPort("public", "8081", "tcp", 0)
	if err != nil {
		log.Fatalf("NewClient failed: %s", err)
	}
	activeZones, err := client.GetActiveZones()
	if err != nil {
		log.Fatalf("get active zones failed: %s", err)
	}
	log.Printf("%#v\n", activeZones)
}
