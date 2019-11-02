package main

import (
	mw "github.com/olegabu/go-mimblewimble"
	"io/ioutil"
)

func main() {
	bytes, _ := ioutil.ReadFile("1g_repost_fix_kernel.json")

	err := mw.ValidateSignature(bytes)
	if err != nil {
		println(err.Error())
	}

	err = mw.ValidateCommitmentsSum(bytes)
	if err != nil {
		println(err.Error())
	}
}
