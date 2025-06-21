package main

import "fmt"
func main() {
	day := "Monday"
	switch day {
	case "Monday":
		fmt.Println("It's Monday")
	case "Tuesday":
		fmt.Println("It's Tuesday")
	case "Wednesday":
		fmt.Println("It's Wednesday")
	case "Thursday":
		fmt.Println("It's Thursday")
	case "Friday":
		fmt.Println("It's Friday")
	default:
		fmt.Println("It's the weekend")
	}
	num := 2
	switch num {
	case 1, 2, 3:
		fmt.Println("Number is between 1 and 3")
	case 4, 5, 6:
		fmt.Println("Number is between 4 and 6")
	default:
		fmt.Println("Number is greater than 6 or less than 1")
	}

	// type switch
	whoAmI := func (i interface{})  {
		switch i := i.(type) {
		case int:
			fmt.Println("I am an int:", i)
		case string:
			fmt.Println("I am a string:", i)
		case bool:
			fmt.Println("I am a bool:", i)	
		default:
			fmt.Println("other:", i)					
		}
	}
	whoAmI(42.09)
}