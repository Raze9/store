package main

import (
	"GOproject/GIT/mail/conf"

	"GOproject/GIT/mail/routes"
)

func main() {
	conf.Init()
	r := routes.NewRouter()
	_ = r.Run(conf.HttpPort)

}
