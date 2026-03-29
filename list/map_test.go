package main

import (
	"fmt"
	"testing"
)

type TestProduct struct {
	id    string
	title string
	price float64
}

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
