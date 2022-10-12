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

	

	fil := df.Filter(
		dataframe.F{
			Colname:    "S1",
			Comparator: series.Eq,
			Comparando: 23,
		},

)
	
	fmt.Println("O número 23 aparece em ",fil.Nrow(), "jogos")
	


	
	

	file.Close()
}

