package main

import "log"

// func main() {
// 	cat := "cat"
// 	if cat == "cat" {
// 		log.Println("Cat")
// 	} else {
// 		log.Println("Not Cat")
// 	}
// }

// func main() {
// 	num := 100
// 	isTrue := false
	
// 	if num > 99 && !isTrue {
// 		log.Println("OX")
// 	} else if num < 100 && isTrue {
// 		log.Println("XO")
// 	} else if  num >  100 && isTrue{
// 		log.Println("OO")
// 	} else {
// 		log.Println("XX")
// 	}
// }

func main() {
	myVar := "horse"
	switch myVar {
	case "cat":
		log.Println("Cat")
	case "dog":
		log.Println("Dog")
	case "fish":
		log.Println("Fish")
	default:
		log.Println("--")
	}
}