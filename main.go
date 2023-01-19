package main

import "Golang/Routes"

func main() {
	r := Routes.SetupRouter()
	r.Run()
}
