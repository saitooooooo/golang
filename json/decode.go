package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type ip struct {
	Origin string `json:"origin"`
	URL    string `json:"url"`
}

// 継続的にデータが流れてくるようなストリーミングデータの操作をサポートしているのが、Decode
func Decode() {
	f, err := os.Open("ip.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var resp ip
	if err := json.NewDecoder(f).Decode(&resp); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
