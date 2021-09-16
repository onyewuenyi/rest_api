package main

import (
	"github.com/onyewuenyi/rest_api/pkg/api"
)

func main() {

	checkErr(api.Start())
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
