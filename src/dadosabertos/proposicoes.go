package dadosabertos

import (
	"encoding/json"
	"fmt"
)

type ProposicaoSummary struct {
	Id        int    `json:"id"`
	Uri       string `json:"uri"`
	SiglaTipo string `json:"siglaTipo"`
	CodTipo   int    `json:"codTipo"`
	Numero    int    `json:"numero"`
	Ano       int    `json:"ano"`
	Ementa    string `json:"ementa"`
}

type StatusProposicao struct {
	DataHora            string `json:"dataHora"`
	Sequencia           int    `json:"sequencia"`
	SiglaOrgao          string `json:"siglaOrgao"`
	UriOrgao            string `json:"uriOrgao"`
	Regime              string `json:"regime"`
	DescricaoTramitacao string `json:"descricaoTramitacao"`
	CodTipoTramitacao   string `json:"codTipoTramitacao"`
	DescricaoSituacao   string `json:"descricaoSituacao"`
	CodSituacao         int    `json:"codSituacao"`
	Despacho            string `json:"despacho"`
	Url                 string `json:"url"`
	Ambito              string `json:"ambito"`
	Apreciacao          string `json:"apreciacao"`
}

type Proposicao struct {
	Id                int              `json:"id"`
	Uri               string           `json:"uri"`
	SiglaTipo         string           `json:"siglaTipo"`
	CodTipo           int              `json:"codTipo"`
	Numero            int              `json:"numero"`
	Ano               int              `json:"ano"`
	Ementa            string           `json:"ementa"`
	DataApresentacao  string           `json:"dataApresentacao"`
	UriOrgaoNumerador string           `json:"uriOrgaoNumerador"`
	StatusProposicao  StatusProposicao `json:"statusProposicao"`
	UriAutores        string           `json:"uriAutores"`
	DescricaoTipo     string           `json:"descricaoTipo"`
	EmentaDetalhada   string           `json:"ementaDetalhada"`
	Keywords          string           `json:"keywords"`
	UriPropPrincipal  string           `json:"uriPropPrincipal"`
	UriPropAnterior   string           `json:"uriPropAnterior"`
	UriPropPosterior  string           `json:"uriPropPosterior"`
	UrlInteiroTeor    string           `json:"urlInteiroTeor"`
	UrnFinal          string           `json:"urnFinal"`
	Texto             string           `json:"texto"`
	Justificativa     string           `json:"justificativa"`
}

type ProposicoesResponse struct {
	Dados []ProposicaoSummary `json:"dados"`
	Links []Link              `json:"links"`
}

type ProposicaoResponse struct {
	Dados Proposicao `json:"dados"`
}

func GetProposicoes(page int) ([]ProposicaoSummary, string) {
	var result ProposicoesResponse
	response, err := GetRequest(fmt.Sprintf("/proposicoes?ordem=DESC&ordenarPor=id&pagina=%d&itens=100", page))
	if err != nil {
		logger.Error(err)
		return []ProposicaoSummary{}, ""
	}

	if err = json.Unmarshal(response, &result); err != nil {
		logger.Error(err)
		return []ProposicaoSummary{}, ""
	}

	next := ""
	for _, link := range result.Links {
		if link.Rel == "next" {
			next = link.Href
		}
	}
	return result.Dados, next
}

func GetProposicao(id int) (Proposicao, error) {
	var result ProposicaoResponse
	response, err := GetRequest(fmt.Sprintf("/proposicoes/%d", id))
	if err != nil {
		logger.Errorf("Error unmarshalling proposicao %v", response)
		return Proposicao{}, err
	}

	if err = json.Unmarshal(response, &result); err != nil {
		logger.Error(err)
		logger.Error(string(response))
		return Proposicao{}, err
	}

	return result.Dados, nil
}

func GetProposicaoIngest(id int, proposicoes chan Proposicao) {
	proposicao, err := GetProposicao(id)
	if err != nil {
		return
	}
	proposicoes <- proposicao
}
