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

	

	if df.Elem(4,2) == df.Elem(3,2) {
		fmt.Println("São iguais")
	}else{
		fmt.Println("São diferentes")
	}
	

	file.Close()
}

