package main

import addFunc "License/Package"

func main() {
	/*
		fmt.Println("Please enter your first integer")
		_, err := fmt.Scan(&int1)
		if err != nil {
			return
		}
		fmt.Println("Please enter your second integer")
		_, err2 := fmt.Scan(&int2)
		if err2 != nil {
			return
		}
		if addFunc.Addition(int1, int2) == 0 {
			fmt.Println("System is busy right now, please try again later!")
		} else {
			fmt.Println("the result is", addFunc.Addition(int1, int2))
		}

	*/
	addFunc.Addition()
}
