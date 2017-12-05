package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	macString := "E0:3F:49:18:9F:BF"
	macParts := strings.Split(macString, ":")
	//var mac []byte
	l := len(macParts)
	for i := 0; i < l; i++ {
		x, err := strconv.ParseInt(macParts[i], 16, 0)
		fmt.Println(macParts[i], byte(x), err)
		// if i, err := strconv.ParseInt(macParts[i], 16, 0); err != nil {
		//  	fmt.Println(i)
		// }

	}

	//macStringWithoutDots := strings.Replace(macString, ":", "", -1)
}
