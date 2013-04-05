package main

import (
	"fmt"
	"log"

	"github.com/AlekSi/airbrake-go"
)

func main() {
	err := airbrake.Notify(fmt.Errorf("%s", "PANIC!"))
	log.Print(err)
}
