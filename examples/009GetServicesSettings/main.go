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
	ss, err := client.GetServiceSettings("ssh")
	if err != nil {
		log.Fatalf("get service settings failed: %s\n", err)
	}
	log.Printf("%#v\n", ss)
}
