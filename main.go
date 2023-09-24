package main

import (
	"acompanhalegis-analytics-api/src/dadosabertos"
	"acompanhalegis-analytics-api/src/elastic"
	"fmt"
	"time"
)

func main() {
	var pageNumber int = 1
	var counter int = 0
	var ids []int
	var proposicoesChan = make(chan dadosabertos.Proposicao, 1000)
	go ingester(proposicoesChan)
	for {
		page, next := dadosabertos.GetProposicoes(pageNumber)
		if next == "" {
			break
		}
		for _, proposicao := range page {
			ids = append(ids, proposicao.Id)
			counter++
		}
		pageNumber++
		time.Sleep(3 * time.Second)
	}
	fmt.Println(fmt.Sprintf("Total de ids: %d", len(ids)))
	time.Sleep(2 * time.Second)
	for i, id := range ids {
		go dadosabertos.GetProposicaoIngest(id, proposicoesChan)
		if i%20 == 0 {
			time.Sleep(10 * time.Second)
		}
	}
	fmt.Println(fmt.Sprintf("Total de proposicoes: %d", counter))
}

func ingester(proposicoes chan dadosabertos.Proposicao) {
	for {
		proposicao := <-proposicoes
		go elastic.IngestProposicao(proposicao)
	}
}
