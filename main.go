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

	fmt.Println("")
	fmt.Println("---------------------------ANÁLISE DA MEGASENA-------------------------------")
	fmt.Println("")
	fmt.Print("Digite um número para saber a frequência entre os sorteios:   ")
	

	numeroescolhido := 1
	fmt.Scan(&numeroescolhido)
	
	for numeroescolhido != 0 {	
	

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



	
	
	fmt.Println("")
	fmt.Println("----------------------------ANÁLISE DO NÚMERO ESCOLHIDO-----------------------------")
	fmt.Println("")
	fmt.Println("Total de jogos analisados:", df.Nrow())	
	fmt.Println("Número escolhido ", numeroescolhido)
	fmt.Println("Apareceu em: ", fil.Nrow(), " sorteios")
	fmt.Println("Em porcentagem isto significa: ",(float64(fil.Nrow()) / float64(df.Nrow()))*100., "%" )
	fmt.Println("")

	fmt.Println("--------------------------------------OUTRA ANÁLISE ?-----------------------------------------")
	fmt.Println("")
	fmt.Print("Digite um número ou 0 para sair do programa: ")
	fmt.Scan(&numeroescolhido)

	}	
	fmt.Println("")
	fmt.Println("------------------------------ESCOLHA UM JOGO PARA VER OS NÚMEROS SORTEADOS -----------------")
	i := 1
	fmt.Print("Digite um concurso entre os ",df.Nrow()  ,". ou 0 para sair do programa:  ")
	fmt.Scan(&i)

	for i != 0 {
		
		fmt.Println(df.Subset(df.Nrow() - i))
		fmt.Println("------------------------------ESCOLHA OUTRO JOGO PARA VER OS NÚMEROS SORTEADOS -----------------")
		fmt.Println("")
		fmt.Print("Digite um concurso entre os ",df.Nrow()  ,". ou 0 para sair do programa:  ")
		fmt.Scan(&i)
	}
	fmt.Println("")
	fmt.Println("---------------------------O SISTEMA AGRADECE SUA VISITA-----------------------------")
	fmt.Println("")

	file.Close()
}

