package main

func main() {
	user := NewUser(&UserCount{4})
	user.Request()
}
