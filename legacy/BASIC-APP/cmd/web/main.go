package main // package name

import (
	"fmt"
	"net/http"
	"udemy/pkg/handlers"
)

const portNumber = ":8080"

// main is the main application function
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)


	fmt.Printf("Starting application on port %s\n", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}

// Section Archives
// func main() {
// 	// request is pointer
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		// need writer in first arg
// 		n, err := fmt.Fprintf(w, "Hello World!")
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		// Sprintf any data type to string
//    // %d = int, %s = string
// 		fmt.Println(fmt.Sprintf("Number of bytes written: %d", n))
// 	})

// 	// start web sever which listens in go
// 	_ = http.ListenAndServe(":8080", nil)
// }

// =======
// func Home(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "This is the home page")
// }

// // About is the about page handler
// func About(w http.ResponseWriter, r *http.Request) {
// 	sum := addValues(2,2)
// 	_,_ = fmt.Fprintf(w, "This is about page and 2 + 2 is %d\n", sum)
// }

// // addValues adds two ints and return the sum
// func addValues(x, y int) int {
// 	sum := x + y
// 	return sum
// }

// func Divide(w http.ResponseWriter, r *http.Request) {
// 	f, err := divideValues(100.0, 0.0)
// 	if  err != nil {
// 		fmt.Fprint(w,"Cannot divide by 0")
// 		return
// 	}

// 	fmt.Fprintf(w, "%f divided by %f is % f\n", 100.0, 0.0, f)
// }

// func divideValues(x, y float32) (float32, error) {
// 	if y <= 0 {
// 		err := errors.New("cannot divide by zero")
// 		return 0, err
// 	}
// 	result := x / y
// 	return result, nil
// }