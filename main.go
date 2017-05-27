package main

import "github.com/koodinikkarit/wompatti/wompatti"

func main() {
	context := wompatti.CreateContext(
		"jaska",
		"asdf321",
		"localhost",
		"3306",
		"wompatti",
		"5052",
	)

	context.Start()
}
