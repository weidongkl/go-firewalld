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
	testZoneSetting := firewalld.ZoneSetting{
		Version:     "1.0",
		Name:        "testzone",
		Description: "temporary zone for testing",
		Target:      "default",
	}
	// firewall-cmd  --delete-zone testzone --permanent
	err = client.AddZone(testZoneSetting)
	if err != nil {
		log.Fatalf("set zone failed: %s\n", err)
	}
	names, err := client.GetZoneNames()
	if err != nil {
		log.Fatalf("set zone failed: %s\n", err)
	}
	log.Printf("zone list: %#v\n", names)
	zoneSetting, err := client.GetZoneSettings("testzone")
	if err != nil {
		log.Fatalf("get zone settings failed: %s\n", err)
	}
	log.Printf("zone setting: %#v\n", zoneSetting)
	path, err := client.GetZoneByName("testzone")
	if err != nil {
		log.Fatalf("get zone  path failed: %s\n", err)
	}
	log.Printf("zone path: %#v\n", path)
}
