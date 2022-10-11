package main

import (
	"fmt"
	"log"
	"os"
	"github.com/go-gota/gota/dataframe"
	
)


func main() {

	file, err := os.Open("./resultados.csv")

	if err != nil{
		log.Fatal(err)
	}

	df := dataframe.ReadCSV(file)
	fmt.Println("Serão analisados ",df.Nrow(), " jogos")

	numero := 3

	if df.Elem(4,2) == numero.Elem() {
		fmt.Println("São iguais")
	}else{
		fmt.Println("São diferentes")
	}
	

	file.Close()
}

