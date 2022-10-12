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
	fmt.Println("Foram analisados ",df.Nrow(), " jogos")

	numeroescolhido := 15
	

	fil := df.Filter(
		dataframe.F{
			Colname:    "S1",
			Comparator: series.Eq,
			Comparando: numeroescolhido,
		},
		dataframe.F{
			Colname:    "S2",
			Comparator: series.Eq,
			Comparando: numeroescolhido,
		},
		dataframe.F{
			Colname:    "S3",
			Comparator: series.Eq,
			Comparando: numeroescolhido,
		},
		dataframe.F{
			Colname:    "S4",
			Comparator: series.Eq,
			Comparando: numeroescolhido,
		},
		dataframe.F{
			Colname:    "S5",
			Comparator: series.Eq,
			Comparando: numeroescolhido,
		},
		dataframe.F{
			Colname:    "S6",
			Comparator: series.Eq,
			Comparando: numeroescolhido,
		},
		
)
	
	fmt.Println("O número 23 apareceu em ",fil.Nrow(), " jogos analisados")

	file.Close()
}

