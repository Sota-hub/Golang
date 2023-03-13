package main

import "log"

// func main() {
// 	animals := []string{"dog","fish","horse","cow","cat"}
// 	// for index item := range slice {}
// 	for _, animal := range animals {
// 		log.Println(animal)
// 	}
// }

// func main() {
// 	animals := make(map[string]string)
// 	animals["dog"] = "Fido"
// 	animals["cat"] = "Fluffy"
//	// first one is key and second one is value if using with map
// 	for key, value := range animals {
// 		log.Println(key, value)
// 	}
// }

// func main() {
// 	var line = "Once upon a time"
// 	line = "x"
// 	for i, l := range line {
// 		log.Println(i, l)
// 		// 0, 79 <- string is bytes
// 	}
// }

func main() {
	type User struct {
		first string
		last string
		email string
		age int
	}
	var users []User
	users = append(users, User{"John", "Smith", "john@email.com",30 })
	users = append(users, User{"Mary", "Jones", "mary@email.com",20 })
	users = append(users, User{"Sally", "Brown", "sarry@email.com",45 })

	for _, l := range users {
		log.Println(l.first, l.last, l.email, l.age)
	}
}