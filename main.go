package main

import (
	"fmt"
	"log"
	"os"
	"github.com/go-gota/gota/dataframe"
	
)


func main() {

	file, err := os.Open("./resultadosmega.csv")

	if err != nil{
		log.Fatal(err)
	}

	df := dataframe.ReadCSV(file)

	fmt.Println(df)

	file.Close()
}