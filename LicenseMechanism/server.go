package main

type Package struct{}

type Server interface {
	Request()
}

type UserCount struct {
	UserCount int
}

type OnlineNumber struct {
	server Server
	users  *UserCount
}

func NewUser(user *UserCount) *OnlineNumber {
	return &OnlineNumber{&Package{}, user}
}
