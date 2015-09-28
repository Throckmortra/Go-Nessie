package main

import (
	"Go-Nessie/nessie-go-wrapper"
	"fmt"
)

func main() {
	client := gonessie.NesClient()
	client.SetKey("d566c0e9c969eb4c02760ef8ecbcabf0")
	//account = client.Account()
	//x := gomodels.Account{}

	//var x gomodels.Account
	//fmt.Println("ayy" + x)
	x := client.Account()
	fmt.Println("aaaaay " + x[0].GetID())
}
