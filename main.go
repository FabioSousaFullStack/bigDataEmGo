package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)


func main() {

	file, err := os.Open("./resultados.csv")

	if err != nil{
		log.Fatal(err)
	}

	df := dataframe.ReadCSV(file)
	fmt.Println("Serão analisados ",df.Nrow(), " jogos")

	
	teste := df.Col("S1")
	fmt.Println(teste)

	
	

	file.Close()
}

