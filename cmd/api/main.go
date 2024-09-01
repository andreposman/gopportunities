package main

import (
	"fmt"

	"github.com/andreposman/gopportunities/cmd/api/router"
)

func main() {
	fmt.Println("-------------------")
	fmt.Println("Gopportunities")
	fmt.Println("-------------------")

	router.InitRouter()

}
