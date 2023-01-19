package main

import "serverLocalGo/Routes"

func main() {
	r := Routes.SetupRouter()
	r.Run()
}
