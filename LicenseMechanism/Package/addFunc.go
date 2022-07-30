package addFunc

import (
	"fmt"
	"os"
)

type NumberOfUsers struct {
	users int
}

var base NumberOfUsers

func (n NumberOfUsers) getUser() int {
	return n.users
}

func Addition() int {
	base.users = base.getUser() + 1 // increment user number

	if base.getUser() > 2 {
		fmt.Print("System is busy right now, TRY AGAIN LATER!")
		os.Exit(0)
	}
	var askContinue string
	var int1 int
	var int2 int
	var flag bool = true
	var result int

	for flag {
		fmt.Println("Please enter your first integer")
		_, err := fmt.Scan(&int1)
		if err != nil {
			os.Exit(0)
		}
		fmt.Println("Please enter your second integer")
		_, err2 := fmt.Scan(&int2)
		if err2 != nil {
			os.Exit(0)
		}

		if base.getUser() > 2 {
			return 0
		} else {
			result = Add(int1, int2)
			// decrement user number
			fmt.Println("Result is", result)
			fmt.Println("Do you want to continue? ( y for yes, n for no)")
			fmt.Scan(&askContinue)
			if askContinue == "n" {
				flag = false
			}
		}
	}
	base.users = base.getUser() - 1
	return result
}

func Add(int1, int2 int) int {
	return int1 + int2
}
