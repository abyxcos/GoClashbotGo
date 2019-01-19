package main

import (
	"io/ioutil"
	"fmt"
	"strings"

	"github.com/ganino/goclash"
)

// Define a global clash API object so we don't have to pass it around
// There will only ever be one connection and key to the API
var clash *goclash.Clash


func main() {
	api_key, err := ioutil.ReadFile("secret.key")
	if err != nil {
		fmt.Printf("Please get a secret API key from https://developer.clashofclans.com/ and place it in the file secret.key")
	}

	clash = goclash.New(strings.TrimSpace(string(api_key)))

	//player, err :=  client.GetPlayer("#LCCRROVU")
	//if err != nil {
	//	fmt.Printf("Error: player not found: %s\n", err)
	//}
	//fmt.Printf("%+v\n", player)

	getClan("#9yu8pqgv")
}


func getClan(tag string) {
	clan, _ := clash.GetClan(tag)

	fmt.Printf("%+v\n", clan)

	return clan, nil
}
