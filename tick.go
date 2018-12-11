package main

import (
	"log"
    "go-tick/util"
)

func main() {
    home := common.GetEnv("HOME", "unknown")
    log.Println(home)
}
