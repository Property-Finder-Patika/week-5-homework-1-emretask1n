package main

import "fmt"

func (p *Package) Request() {
	fmt.Println("Package is started")
}

func (o *OnlineNumber) Request() {
	if o.users.UserCount <= 3 {
		o.server.Request()
	} else {
		fmt.Println("all of the licenses have all been taken up and retry later.")
	}
}
