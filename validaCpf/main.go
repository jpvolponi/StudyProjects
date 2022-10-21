package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, " POST request successful\n")
	cpf := ValidaCpf(r.FormValue("CPF"))
	//cpfResponse := ValidaCpf(cpf)
	fmt.Fprintf(w, "\nCPF: %s is %v.", r.FormValue("CPF"), cpf)
}
func main() {
	//fmt.Println(ValidaCpf())
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/consulta", formHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
