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
	
	
	for i := 0; i <= df.Nrow(); i++{
		row := df.Subset(i)
		fmt.Println(row)
	}
	
	

	file.Close()
}