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

	sel1 := df.Select([]int{0, 2})

	fmt.Println(sel1)
	

	
	

	file.Close()
}

