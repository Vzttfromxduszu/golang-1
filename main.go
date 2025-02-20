package main

import "github.com/Vzttfromxduszu/golang-1.git/common/initialize"

func main() {
	print("hello")
	initialize.LoadConfig()
	initialize.MySQL()
	initialize.Router()
}
