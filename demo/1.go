package main

import (
	"fmt"
	"utils/utils"
)

const salt = "g"

func main() {

	err := utils.InitHashID(salt, 1)
	if err != nil {
		panic(err)
	}
	id, err := utils.EncodeID(1)
	if err != nil {
		return
	}
	fmt.Println(id)
	decodeID, err := utils.DecodeID(id)
	if err != nil {
		return
	}
	fmt.Println(decodeID)

}
