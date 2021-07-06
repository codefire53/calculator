package main

import (
	"fmt"
	"net/http"
)
func add(num_a, num_b int) (output int) {
	output = num_a + num_b
	return 
} 

func sub(num_a, num_b int) (output int) {
	output = num_a - num_b
	return 
} 

func mul(num_a, num_b int) (output int) {
	output = num_a * num_b
	return 
} 

func div(num_a, num_b int) (output int) {
	output = num_a / num_b
	return 
} 

func pow(base, exp int) (output int) {
	output = base ** exp
	return 
} 

func calculatorHandler(w http.ResponseWriter, req *http.Request) {

	numA := req.URL.Query().Get("num_a")
	numB := req.URL.Query().Get("num_b")
	operator := req.URL.Query().Get("operator")
	res := 0
	switch operator {
		case "+":
			res = add(numA, numB)
			break
		case "-":
			res = sub(numA, numB)
			break
		case "*":
			res = mul(numA, numB)
			break
		case "/":
			res = div(numA, numB)
			break
		case "^":
			res = pow(numA, numB)
			break
		case "**":
			res = pow(numA, numB)
			break
		
	}
	fmt.Fprintf(w, fmt.Sprint(res))
}

func main() {

	http.HandleFunc("/calculator", calculatorHandler)

	fmt.Println("serving HTTP...")
	http.ListenAndServe(":8989", nil)
}
