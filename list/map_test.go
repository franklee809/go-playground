package main

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	websites := map[string]string{
		"Google": "https://www.google.com",
		"AWS":    "https://www.aws.amazon.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["AWS"])

	delete(websites, websites["Google"])
	fmt.Println(websites)
}
